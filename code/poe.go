package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// =======================
// Config
// =======================

const (
	Version      = "0.1"
	LedgerDir    = "ledger"
	LedgerFull   = "full.txt"
	MaxJSONBytes = 8 * 1024 // 8KiB (alinha com teu doc)
)

// =======================
// Tipos
// =======================

type PoEEvent struct {
	Version          string `json:"version"`
	EventID          string `json:"event_id"`
	PreviousEventHash string `json:"previous_event_hash"`
	PayloadHash      string `json:"payload_hash"`
	SourceID         string `json:"source_id"`
	LocalTimestamp   string `json:"local_timestamp"`
}

type LedgerEntry struct {
	Sequence       int    `json:"sequence"`
	EventHash      string `json:"event_hash"`
	ObservedPrev   string `json:"observed_prev_hash"`
	SourceID       string `json:"source_id"`
	PayloadHash    string `json:"payload_hash"`
	AcceptedAtUTC  string `json:"accepted_at_utc"`
	EventID        string `json:"event_id"`
}

type Receipt struct {
	Accepted       bool   `json:"accepted"`
	Sequence       int    `json:"sequence,omitempty"`
	EventHash      string `json:"event_hash,omitempty"`
	ObservedPrev   string `json:"observed_prev_hash,omitempty"`
	AcceptedAtUTC  string `json:"accepted_at_utc,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	Error          string `json:"error,omitempty"`
}

type BatchRequest struct {
	Events []PoEEvent `json:"events"`
}

type BatchResponse struct {
	AcceptedCount int       `json:"accepted_count"`
	Receipts      []Receipt `json:"receipts"`
}

// =======================
// Estado global (FIFO soberano)
// =======================

var (
	mu           sync.Mutex
	fifoSequence int
	lastHash     = "GENESIS"
)

// =======================
// Utils
// =======================

func utcNowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func mustLedgerPaths() (fullPath string, err error) {
	if err := os.MkdirAll(LedgerDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(LedgerDir, LedgerFull), nil
}

func dailyLedgerPath(acceptedAtUTC string) (string, error) {
	t, err := time.Parse(time.RFC3339, acceptedAtUTC)
	if err != nil {
		return "", err
	}
	day := t.UTC().Format("2006-01-02")
	return filepath.Join(LedgerDir, day+".txt"), nil
}

// Hash canônico do ledger entry (determinístico).
// Inclui seq + dados do evento + prev observado.
func computeEventHash(seq int, ev PoEEvent, observedPrev string) string {
	// NUNCA muda ordem/estrutura dessa string sem versionar (v0.2+)
	data := fmt.Sprintf(
		"%d|%s|%s|%s|%s|%s|%s",
		seq,
		ev.Version,
		ev.EventID,
		ev.PayloadHash,
		ev.SourceID,
		ev.LocalTimestamp,
		observedPrev,
	)
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func validateEvent(ev PoEEvent) error {
	if ev.Version == "" {
		ev.Version = Version
	}
	if ev.Version != Version {
		return errors.New("ERR_BAD_VERSION")
	}
	if ev.EventID == "" || ev.PayloadHash == "" || ev.PreviousEventHash == "" || ev.SourceID == "" || ev.LocalTimestamp == "" {
		return errors.New("ERR_BAD_FORMAT")
	}
	// Aceita payload_hash em hex "0x..." ou só hex, mas exige conteúdo mínimo
	h := strings.TrimPrefix(ev.PayloadHash, "0x")
	if len(h) < 16 {
		return errors.New("ERR_BAD_FORMAT")
	}
	return nil
}

// =======================
// Persistência: ledger append-only
// Formato por linha (full e diário):
// sequence|event_hash|observed_prev_hash|source_id|payload_hash|event_id|accepted_at_utc
// =======================

func appendLedger(entry LedgerEntry) error {
	fullPath, err := mustLedgerPaths()
	if err != nil {
		return err
	}

	line := fmt.Sprintf("%d|%s|%s|%s|%s|%s|%s\n",
		entry.Sequence,
		entry.EventHash,
		entry.ObservedPrev,
		entry.SourceID,
		entry.PayloadHash,
		entry.EventID,
		entry.AcceptedAtUTC,
	)

	// full.txt
	if err := appendFile(fullPath, line); err != nil {
		return err
	}

	// YYYY-MM-DD.txt (UTC)
	dayPath, err := dailyLedgerPath(entry.AcceptedAtUTC)
	if err != nil {
		return err
	}
	if err := appendFile(dayPath, line); err != nil {
		return err
	}

	return nil
}

func appendFile(path string, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(line)
	return err
}

// Carrega head do disco (última linha do full.txt)
func loadStateFromDisk() error {
	fullPath, err := mustLedgerPaths()
	if err != nil {
		return err
	}

	f, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			// sem ledger ainda
			fifoSequence = 0
			lastHash = "GENESIS"
			return nil
		}
		return err
	}
	defer f.Close()

	// lê última linha sem carregar tudo (scan inteiro é OK pro MVP; depois otimiza)
	sc := bufio.NewScanner(f)
	var lastLine string
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			lastLine = line
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	if lastLine == "" {
		fifoSequence = 0
		lastHash = "GENESIS"
		return nil
	}

	parts := strings.Split(lastLine, "|")
	if len(parts) < 2 {
		return fmt.Errorf("ledger corrupt: last line invalid")
	}

	seq, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ledger corrupt: bad seq")
	}
	evHash := parts[1]

	fifoSequence = seq
	lastHash = evHash
	return nil
}

// =======================
// HTTP helpers
// =======================

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func readBodyLimited(r *http.Request, limit int64) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(io.LimitReader(r.Body, limit))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ajusta se você quiser restringir.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// =======================
// Handlers
// =======================

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{
		"ok":        true,
		"version":   Version,
		"utc":       utcNowRFC3339(),
		"last_seq":  fifoSequence,
		"last_hash": lastHash,
	})
}

func headHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	writeJSON(w, 200, map[string]any{
		"last_sequence": fifoSequence,
		"last_event_hash": lastHash,
	})
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	raw, err := readBodyLimited(r, MaxJSONBytes)
	if err != nil {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}

	var ev PoEEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}
	if ev.Version == "" {
		ev.Version = Version
	}

	if err := validateEvent(ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// regra: prev_hash deve bater no head atual
	if ev.PreviousEventHash != lastHash {
		http.Error(w, "ERR_BAD_PREV_HASH", http.StatusConflict)
		return
	}

	observedPrev := lastHash
	acceptedAt := utcNowRFC3339()

	fifoSequence++
	evHash := computeEventHash(fifoSequence, ev, observedPrev)

	entry := LedgerEntry{
		Sequence:      fifoSequence,
		EventHash:     evHash,
		ObservedPrev:  observedPrev,
		SourceID:      ev.SourceID,
		PayloadHash:   ev.PayloadHash,
		EventID:       ev.EventID,
		AcceptedAtUTC: acceptedAt,
	}

	if err := appendLedger(entry); err != nil {
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}

	lastHash = evHash

	writeJSON(w, 200, Receipt{
		Accepted:      true,
		Sequence:      entry.Sequence,
		EventHash:     entry.EventHash,
		ObservedPrev:  entry.ObservedPrev,
		AcceptedAtUTC: entry.AcceptedAtUTC,
		EventID:       entry.EventID,
	})
}

func submitBatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	raw, err := readBodyLimited(r, int64(MaxJSONBytes)*32) // batch maior, mas ainda limitado
	if err != nil {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}

	// Aceita 2 formatos:
	// 1) {"events":[...]}
	// 2) [...]
	var req BatchRequest
	var events []PoEEvent

	// tenta {"events":...}
	if err := json.Unmarshal(raw, &req); err == nil && len(req.Events) > 0 {
		events = req.Events
	} else {
		// tenta array direto
		if err := json.Unmarshal(raw, &events); err != nil {
			http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
			return
		}
	}

	if len(events) == 0 {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}

	// valida fora do lock
	for i := range events {
		if events[i].Version == "" {
			events[i].Version = Version
		}
		if err := validateEvent(events[i]); err != nil {
			writeJSON(w, http.StatusBadRequest, BatchResponse{
				AcceptedCount: 0,
				Receipts: []Receipt{
					{Accepted: false, Error: err.Error(), EventID: events[i].EventID},
				},
			})
			return
		}
	}

	mu.Lock()
	defer mu.Unlock()

	resp := BatchResponse{
		AcceptedCount: 0,
		Receipts:      make([]Receipt, 0, len(events)),
	}

	// batch precisa encadear do head atual
	for i := 0; i < len(events); i++ {
		ev := events[i]

		if ev.PreviousEventHash != lastHash {
			// falhou aqui -> aborta resto
			resp.Receipts = append(resp.Receipts, Receipt{
				Accepted: false,
				EventID:  ev.EventID,
				Error:    "ERR_BAD_PREV_HASH",
			})
			for j := i + 1; j < len(events); j++ {
				resp.Receipts = append(resp.Receipts, Receipt{
					Accepted: false,
					EventID:  events[j].EventID,
					Error:    "BATCH_ABORTED",
				})
			}
			writeJSON(w, http.StatusConflict, resp)
			return
		}

		observedPrev := lastHash
		acceptedAt := utcNowRFC3339()

		fifoSequence++
		evHash := computeEventHash(fifoSequence, ev, observedPrev)

		entry := LedgerEntry{
			Sequence:      fifoSequence,
			EventHash:     evHash,
			ObservedPrev:  observedPrev,
			SourceID:      ev.SourceID,
			PayloadHash:   ev.PayloadHash,
			EventID:       ev.EventID,
			AcceptedAtUTC: acceptedAt,
		}

		if err := appendLedger(entry); err != nil {
			// erro interno -> aborta batch
			resp.Receipts = append(resp.Receipts, Receipt{
				Accepted: false,
				EventID:  ev.EventID,
				Error:    "ERR_INTERNAL",
			})
			for j := i + 1; j < len(events); j++ {
				resp.Receipts = append(resp.Receipts, Receipt{
					Accepted: false,
					EventID:  events[j].EventID,
					Error:    "BATCH_ABORTED",
				})
			}
			writeJSON(w, http.StatusInternalServerError, resp)
			return
		}

		lastHash = evHash

		resp.AcceptedCount++
		resp.Receipts = append(resp.Receipts, Receipt{
			Accepted:      true,
			Sequence:      entry.Sequence,
			EventHash:     entry.EventHash,
			ObservedPrev:  entry.ObservedPrev,
			AcceptedAtUTC: entry.AcceptedAtUTC,
			EventID:       entry.EventID,
		})
	}

	writeJSON(w, 200, resp)
}

func ledgerFullHandler(w http.ResponseWriter, r *http.Request) {
	fullPath, err := mustLedgerPaths()
	if err != nil {
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, fullPath)
}

func ledgerByDateHandler(w http.ResponseWriter, r *http.Request) {
	// /ledger/YYYY-MM-DD.txt
	// path esperado: /ledger/2026-01-28.txt
	p := strings.TrimPrefix(r.URL.Path, "/ledger/")
	if p == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	// segurança: sem ../
	if strings.Contains(p, "..") || strings.ContainsAny(p, "\\") {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	path := filepath.Join(LedgerDir, p)
	http.ServeFile(w, r, path)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("from")
	if q == "" {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}
	from, err := strconv.Atoi(q)
	if err != nil || from < 1 {
		http.Error(w, "ERR_BAD_FORMAT", http.StatusBadRequest)
		return
	}

	fullPath, err := mustLedgerPaths()
	if err != nil {
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}

	f, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			writeJSON(w, 200, []LedgerEntry{})
			return
		}
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	out := make([]LedgerEntry, 0, 128)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		// sequence|event_hash|observed_prev_hash|source_id|payload_hash|event_id|accepted_at_utc
		if len(parts) < 7 {
			continue
		}
		seq, err := strconv.Atoi(parts[0])
		if err != nil || seq < from {
			continue
		}
		out = append(out, LedgerEntry{
			Sequence:      seq,
			EventHash:     parts[1],
			ObservedPrev:  parts[2],
			SourceID:      parts[3],
			PayloadHash:   parts[4],
			EventID:       parts[5],
			AcceptedAtUTC: parts[6],
		})
	}

	if err := sc.Err(); err != nil {
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}

	writeJSON(w, 200, out)
}

// =======================
// Main
// =======================

func main() {
	if err := loadStateFromDisk(); err != nil {
		log.Fatalf("failed to load state: %v", err)
	}

	mux := http.NewServeMux()

	// Core
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/head", headHandler)
	mux.HandleFunc("/submit", submitHandler)
	mux.HandleFunc("/submit_batch", submitBatchHandler)
	mux.HandleFunc("/stream", streamHandler)

	// Ledger files
	mux.HandleFunc("/ledger/full.txt", ledgerFullHandler)
	mux.HandleFunc("/ledger/", ledgerByDateHandler) // /ledger/YYYY-MM-DD.txt

	fmt.Println("PoE FIFO Gateway rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}
