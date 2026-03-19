package main

import (
	"bufio"
	"bytes"
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

const (
	Version         = "0.1"
	LedgerDir       = "ledger"
	LedgerFull      = "full.txt"
	LedgerProvas    = "provas.txt"
	MaxJSONBytes    = 8 * 1024
	BlockchainURL   = "http://localhost:8545"
	ContractAddress = "0xSEU_CONTRATO_AQUI"
	AnchorBatchSize = 1000
	AnchorInterval  = 5 * time.Minute
)

// ── Tipos ────────────────────────────────────────────────────────

type PoEEvent struct {
	Version           string `json:"version"`
	EventID           string `json:"event_id"`
	PreviousEventHash string `json:"previous_event_hash"`
	PayloadHash       string `json:"payload_hash"`
	SourceID          string `json:"source_id"`
	LocalTimestamp    string `json:"local_timestamp"`
	De                string `json:"de,omitempty"`
	Para              string `json:"para,omitempty"`
	Valor             string `json:"valor,omitempty"`
	Moeda             string `json:"moeda,omitempty"`
}

type LedgerEntry struct {
	Sequence      int    `json:"sequence"`
	EventHash     string `json:"event_hash"`
	ObservedPrev  string `json:"observed_prev_hash"`
	SourceID      string `json:"source_id"`
	PayloadHash   string `json:"payload_hash"`
	AcceptedAtUTC string `json:"accepted_at_utc"`
	EventID       string `json:"event_id"`
	De            string `json:"de,omitempty"`
	Para          string `json:"para,omitempty"`
	Valor         string `json:"valor,omitempty"`
	Moeda         string `json:"moeda,omitempty"`
	Liquidado     bool   `json:"liquidado"`
	LiquidadoTx   string `json:"liquidado_tx,omitempty"`
}

type Receipt struct {
	Accepted      bool   `json:"accepted"`
	Sequence      int    `json:"sequence,omitempty"`
	EventHash     string `json:"event_hash,omitempty"`
	ObservedPrev  string `json:"observed_prev_hash,omitempty"`
	AcceptedAtUTC string `json:"accepted_at_utc,omitempty"`
	EventID       string `json:"event_id,omitempty"`
	Error         string `json:"error,omitempty"`
	PreValidado   bool   `json:"pre_validado"`
	MensagemL2    string `json:"mensagem_l2,omitempty"`
}

// ── Prova Individual ─────────────────────────────────────────────
// Uma prova por transação — registrada no provas.txt
// Prova que a transação existiu, foi pré-validada
// e vincula ao txHash da liquidação na blockchain

type ProvaIndividual struct {
	Seq            int    `json:"seq"`
	EventID        string `json:"event_id"`
	SHA256Evento   string `json:"sha256_evento"`   // hash evento completo
	SHA256Payload  string `json:"sha256_payload"`  // hash payload original
	SHA256Conteudo string `json:"sha256_conteudo"` // hash de|para|valor|moeda
	SourceID       string `json:"source_id"`
	RegistradoUTC  string `json:"registrado_utc"`
	De             string `json:"de"`
	Para           string `json:"para"`
	Valor          string `json:"valor"`
	Moeda          string `json:"moeda"`
	Liquidado      bool   `json:"liquidado"`
	LiquidadoTx    string `json:"liquidado_tx,omitempty"`
	LiquidadoUTC   string `json:"liquidado_utc,omitempty"`
	EventHash      string `json:"event_hash"`
	ObservedPrev   string `json:"observed_prev"`
}

// ── Estado global ────────────────────────────────────────────────

var (
	mu                  sync.Mutex
	fifoSequence        int
	lastHash            = "GENESIS"
	pendentesLiquidacao []LedgerEntry
	muPendentes         sync.Mutex
)

// ── Prova Individual — funções ────────────────────────────────────

func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func gerarProva(entry LedgerEntry, ev PoEEvent) ProvaIndividual {
	evJSON, _ := json.Marshal(ev)
	sha256Evento := sha256Hex(string(evJSON))
	sha256Payload := sha256Hex(ev.PayloadHash)
	conteudo := fmt.Sprintf("%s|%s|%s|%s",
		ev.De, ev.Para, ev.Valor, ev.Moeda)
	sha256Conteudo := sha256Hex(conteudo)

	return ProvaIndividual{
		Seq:            entry.Sequence,
		EventID:        entry.EventID,
		SHA256Evento:   sha256Evento,
		SHA256Payload:  sha256Payload,
		SHA256Conteudo: sha256Conteudo,
		SourceID:       entry.SourceID,
		RegistradoUTC:  entry.AcceptedAtUTC,
		De:             ev.De,
		Para:           ev.Para,
		Valor:          ev.Valor,
		Moeda:          ev.Moeda,
		Liquidado:      false,
		EventHash:      entry.EventHash,
		ObservedPrev:   entry.ObservedPrev,
	}
}

// Persiste prova — append-only
// seq|sha256_evento|sha256_payload|sha256_conteudo|
// source_id|de|para|valor|moeda|liquidado|tx|liq_utc|reg_utc|event_hash
func appendProva(p ProvaIndividual) error {
	liquidado := "0"
	if p.Liquidado {
		liquidado = "1"
	}
	linha := fmt.Sprintf("%d|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n",
		p.Seq,
		p.SHA256Evento,
		p.SHA256Payload,
		p.SHA256Conteudo,
		p.SourceID,
		p.De,
		p.Para,
		p.Valor,
		p.Moeda,
		liquidado,
		p.LiquidadoTx,
		p.LiquidadoUTC,
		p.RegistradoUTC,
		p.EventHash,
	)
	path := filepath.Join(LedgerDir, LedgerProvas)
	return appendFile(path, linha)
}

// Marca prova como liquidada com txHash
func marcarProvasLiquidadas(entries []LedgerEntry, txHash string) {
	agora := utcNowRFC3339()
	path := filepath.Join(LedgerDir, LedgerProvas)

	f, err := os.Open(path)
	if err != nil {
		return
	}
	conteudo, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		return
	}

	// Mapa de seqs liquidados
	seqMap := make(map[int]bool)
	for _, e := range entries {
		seqMap[e.Sequence] = true
	}

	linhas := strings.Split(string(conteudo), "\n")
	for i, linha := range linhas {
		if linha == "" {
			continue
		}
		partes := strings.Split(linha, "|")
		if len(partes) < 14 {
			continue
		}
		seqLinha, err := strconv.Atoi(partes[0])
		if err != nil || !seqMap[seqLinha] {
			continue
		}
		partes[9] = "1"
		partes[10] = txHash
		partes[11] = agora
		linhas[i] = strings.Join(partes, "|")
	}

	os.WriteFile(path, []byte(strings.Join(linhas, "\n")), 0644)
	log.Printf("[PROVA] %d provas marcadas como liquidadas | tx=%s",
		len(entries), txHash)
}

// ── Pré-validação L2 ─────────────────────────────────────────────

type ResultadoPreValidacao struct {
	Aprovado bool
	Motivo   string
}

func preValidar(ev PoEEvent) ResultadoPreValidacao {
	if ev.De == "" || ev.Para == "" || ev.Valor == "" {
		return ResultadoPreValidacao{false, "ERR_CAMPOS_LIQUIDACAO_AUSENTES"}
	}
	valor, err := strconv.ParseInt(ev.Valor, 10, 64)
	if err != nil || valor <= 0 {
		return ResultadoPreValidacao{false, "ERR_VALOR_INVALIDO"}
	}
	if ev.De == ev.Para {
		return ResultadoPreValidacao{false, "ERR_DE_IGUAL_PARA"}
	}
	if !enderecoValido(ev.De) || !enderecoValido(ev.Para) {
		return ResultadoPreValidacao{false, "ERR_ENDERECO_INVALIDO"}
	}
	moeda := ev.Moeda
	if moeda == "" {
		moeda = "DREX"
	}
	if moeda != "DREX" && moeda != "BRL" && moeda != "BCA" {
		return ResultadoPreValidacao{false, "ERR_MOEDA_NAO_SUPORTADA"}
	}
	return ResultadoPreValidacao{true, "PRE_VALIDADO_OK"}
}

func enderecoValido(addr string) bool {
	if !strings.HasPrefix(addr, "0x") {
		return false
	}
	h := strings.TrimPrefix(addr, "0x")
	if len(h) != 40 {
		return false
	}
	_, err := hex.DecodeString(h)
	return err == nil
}

// ── Merkle root ───────────────────────────────────────────────────

func computeMerkleRoot(entries []LedgerEntry) string {
	if len(entries) == 0 {
		return strings.Repeat("0", 64)
	}
	layer := make([]string, len(entries))
	for i, e := range entries {
		data := fmt.Sprintf("%d|%s|%s|%s|%s",
			e.Sequence, e.EventHash, e.PayloadHash, e.De, e.Para)
		sum := sha256.Sum256([]byte(data))
		layer[i] = hex.EncodeToString(sum[:])
	}
	for len(layer) > 1 {
		if len(layer)%2 != 0 {
			layer = append(layer, layer[len(layer)-1])
		}
		next := make([]string, len(layer)/2)
		for i := 0; i < len(layer); i += 2 {
			combined := layer[i] + layer[i+1]
			sum := sha256.Sum256([]byte(combined))
			next[i/2] = hex.EncodeToString(sum[:])
		}
		layer = next
	}
	return layer[0]
}

// ── Envio para blockchain ─────────────────────────────────────────
//
// Duas chamadas atomicas numa unica transacao:
//   Chamada 1 -> PoEAnchor.receberPreValidacao()
//     "ancora a prova — esses eventos existiram"
//   Chamada 2 -> SequencerExecution.executeBatch()
//     "liquida os tokens — move o dinheiro de verdade"
//
// Atomico: ou as duas acontecem juntas ou nenhuma acontece.

type InstrucaoBC struct {
	Seq         int    `json:"seq"`
	De          string `json:"de"`
	Para        string `json:"para"`
	Valor       string `json:"valor"`     // centavos
	ValorWei    string `json:"valor_wei"` // wei para o contrato
	Moeda       string `json:"moeda"`
	Hash        string `json:"hash"`
	PayloadHash string `json:"payload_hash"`
	Deadline    int64  `json:"deadline"`
}

// Chamada 1 — PoEAnchor: ancora a prova
type PayloadAnchor struct {
	MerkleRoot string `json:"merkle_root"`
	FromSeq    int    `json:"from_seq"`
	ToSeq      int    `json:"to_seq"`
	EntryCount int    `json:"entry_count"`
	StarkProof string `json:"stark_proof"`
}

// Chamada 2 — SequencerExecution: liquida os tokens
type PayloadExecucao struct {
	Instrucoes []InstrucaoBC `json:"instrucoes"`
	MerkleRoot string        `json:"merkle_root"`
}

// Payload atomico — as duas chamadas juntas
type PayloadAtomico struct {
	Anchor   PayloadAnchor   `json:"anchor"`
	Execucao PayloadExecucao `json:"execucao"`
}

func enviarParaBlockchain(entries []LedgerEntry) (string, error) {
	if len(entries) == 0 {
		return "", fmt.Errorf("batch vazio")
	}

	merkleRoot := computeMerkleRoot(entries)

	// Monta instrucoes com valores reais para execucao
	instrucoes := make([]InstrucaoBC, 0, len(entries))
	for _, e := range entries {
		if e.De == "" || e.Para == "" || e.Valor == "" {
			continue
		}
		valorCentavos, err := strconv.ParseInt(e.Valor, 10, 64)
		if err != nil {
			log.Printf("[EXEC] seq %d: valor invalido %s", e.Sequence, e.Valor)
			continue
		}
		// 1 centavo = 1e16 wei no DREX
		valorWei := fmt.Sprintf("%d0000000000000000", valorCentavos)

		instrucoes = append(instrucoes, InstrucaoBC{
			Seq:         e.Sequence,
			De:          e.De,
			Para:        e.Para,
			Valor:       e.Valor,
			ValorWei:    valorWei,
			Moeda:       e.Moeda,
			Hash:        e.EventHash,
			PayloadHash: e.PayloadHash,
			Deadline:    time.Now().Add(10 * time.Minute).Unix(),
		})
	}

	// Payload atomico — prova + liquidacao juntas
	payload := PayloadAtomico{
		// Ancora a prova no PoEAnchor
		Anchor: PayloadAnchor{
			MerkleRoot: merkleRoot,
			FromSeq:    entries[0].Sequence,
			ToSeq:      entries[len(entries)-1].Sequence,
			EntryCount: len(entries),
			StarkProof: gerarStarkProofStub(merkleRoot, entries),
		},
		// Executa as transferencias no SequencerExecution
		Execucao: PayloadExecucao{
			Instrucoes: instrucoes,
			MerkleRoot: merkleRoot,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}

	// Envia para o no Besu/DREX
	// Em producao: go-ethereum assina e envia tx atomica
	resp, err := http.Post(
		BlockchainURL+"/executar_atomico",
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("blockchain indisponivel: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	txHash := result["tx_hash"]

	log.Printf("[L2->DREX] ANCHOR+EXEC atomico | seq %d->%d | %d instrucoes | merkle=%s... | tx=%s",
		payload.Anchor.FromSeq,
		payload.Anchor.ToSeq,
		len(instrucoes),
		merkleRoot[:16],
		txHash,
	)

	return txHash, nil
}

// Stub do mini STARK — substitui pelo gnark em producao
func gerarStarkProofStub(merkleRoot string, entries []LedgerEntry) string {
	data := fmt.Sprintf("%s|%d|%d|%d",
		merkleRoot,
		entries[0].Sequence,
		entries[len(entries)-1].Sequence,
		len(entries),
	)
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

// ── Scheduler de liquidação ───────────────────────────────────────

func iniciarSchedulerLiquidacao() {
	go func() {
		ticker := time.NewTicker(AnchorInterval)
		defer ticker.Stop()
		for range ticker.C {
			liquidarPendentes()
		}
	}()
}

func liquidarPendentes() {
	muPendentes.Lock()
	if len(pendentesLiquidacao) == 0 {
		muPendentes.Unlock()
		return
	}
	batch := make([]LedgerEntry, len(pendentesLiquidacao))
	copy(batch, pendentesLiquidacao)
	pendentesLiquidacao = nil
	muPendentes.Unlock()

	txHash, err := enviarParaBlockchain(batch)
	if err != nil {
		log.Printf("[L2->DREX] erro: %v", err)
		muPendentes.Lock()
		pendentesLiquidacao = append(batch, pendentesLiquidacao...)
		muPendentes.Unlock()
		return
	}

	// Atualiza provas individuais com txHash
	marcarProvasLiquidadas(batch, txHash)

	log.Printf("[L2->DREX] %d transacoes liquidadas | tx=%s",
		len(batch), txHash)
}

// ── Utils ─────────────────────────────────────────────────────────

func utcNowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func mustLedgerPaths() (string, error) {
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
	return filepath.Join(LedgerDir, t.UTC().Format("2006-01-02")+".txt"), nil
}

func computeEventHash(seq int, ev PoEEvent, observedPrev string) string {
	data := fmt.Sprintf("%d|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		seq, ev.Version, ev.EventID, ev.PayloadHash,
		ev.SourceID, ev.LocalTimestamp,
		ev.De, ev.Para, ev.Valor, observedPrev)
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
	if ev.EventID == "" || ev.PayloadHash == "" ||
		ev.PreviousEventHash == "" || ev.SourceID == "" ||
		ev.LocalTimestamp == "" {
		return errors.New("ERR_BAD_FORMAT")
	}
	h := strings.TrimPrefix(ev.PayloadHash, "0x")
	if len(h) < 16 {
		return errors.New("ERR_BAD_FORMAT")
	}
	return nil
}

// ── Persistência ─────────────────────────────────────────────────

func appendLedger(entry LedgerEntry) error {
	fullPath, err := mustLedgerPaths()
	if err != nil {
		return err
	}
	line := fmt.Sprintf("%d|%s|%s|%s|%s|%s|%s|%s|%s|%s\n",
		entry.Sequence, entry.EventHash, entry.ObservedPrev,
		entry.SourceID, entry.PayloadHash, entry.EventID,
		entry.AcceptedAtUTC, entry.De, entry.Para, entry.Valor)
	if err := appendFile(fullPath, line); err != nil {
		return err
	}
	dayPath, err := dailyLedgerPath(entry.AcceptedAtUTC)
	if err != nil {
		return err
	}
	return appendFile(dayPath, line)
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

func loadStateFromDisk() error {
	fullPath, err := mustLedgerPaths()
	if err != nil {
		return err
	}
	f, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			fifoSequence = 0
			lastHash = "GENESIS"
			return nil
		}
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var lastLine string
	for sc.Scan() {
		if line := strings.TrimSpace(sc.Text()); line != "" {
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
		return fmt.Errorf("ledger corrupt")
	}
	seq, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ledger corrupt: bad seq")
	}
	fifoSequence = seq
	lastHash = parts[1]
	return nil
}

// ── HTTP helpers ──────────────────────────────────────────────────

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func readBodyLimited(r *http.Request, limit int64) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(io.LimitReader(r.Body, limit))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// ── Handlers ──────────────────────────────────────────────────────

func healthHandler(w http.ResponseWriter, r *http.Request) {
	muPendentes.Lock()
	pendentes := len(pendentesLiquidacao)
	muPendentes.Unlock()
	writeJSON(w, 200, map[string]any{
		"ok":                  true,
		"version":             Version,
		"utc":                 utcNowRFC3339(),
		"last_seq":            fifoSequence,
		"last_hash":           lastHash,
		"pendentes_liquidacao": pendentes,
	})
}

func headHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	writeJSON(w, 200, map[string]any{
		"last_sequence":   fifoSequence,
		"last_event_hash": lastHash,
	})
}

// submitHandler — pré-valida, gera prova, enfileira liquidação
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

	// Pré-validação L2
	resultado := preValidar(ev)
	if !resultado.Aprovado {
		http.Error(w, resultado.Motivo, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

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
		De:            ev.De,
		Para:          ev.Para,
		Valor:         ev.Valor,
		Moeda:         ev.Moeda,
		Liquidado:     false,
	}

	// 1. Grava no ledger principal
	if err := appendLedger(entry); err != nil {
		http.Error(w, "ERR_INTERNAL", http.StatusInternalServerError)
		return
	}

	// 2. Gera e grava prova individual
	prova := gerarProva(entry, ev)
	if err := appendProva(prova); err != nil {
		log.Printf("[PROVA] erro ao gravar prova seq=%d: %v",
			entry.Sequence, err)
	}

	lastHash = evHash

	// 3. Enfileira para liquidação na blockchain
	muPendentes.Lock()
	pendentesLiquidacao = append(pendentesLiquidacao, entry)
	nPendentes := len(pendentesLiquidacao)
	muPendentes.Unlock()

	if nPendentes >= AnchorBatchSize {
		go liquidarPendentes()
	}

	// Resposta imediata — L2 confirmou
	valor, _ := strconv.ParseInt(ev.Valor, 10, 64)
	writeJSON(w, 200, Receipt{
		Accepted:      true,
		Sequence:      entry.Sequence,
		EventHash:     entry.EventHash,
		ObservedPrev:  entry.ObservedPrev,
		AcceptedAtUTC: entry.AcceptedAtUTC,
		EventID:       entry.EventID,
		PreValidado:   true,
		MensagemL2: fmt.Sprintf(
			"R$%.2f pre-validado | prova registrada | liquidacao na blockchain em andamento",
			float64(valor)/100),
	})
}

// provaHandler — consulta prova individual de um seq
// GET /prova?seq=42
func provaHandler(w http.ResponseWriter, r *http.Request) {
	seqStr := r.URL.Query().Get("seq")
	if seqStr == "" {
		http.Error(w, "ERR_SEQ_AUSENTE", http.StatusBadRequest)
		return
	}
	seq, err := strconv.Atoi(seqStr)
	if err != nil {
		http.Error(w, "ERR_SEQ_INVALIDO", http.StatusBadRequest)
		return
	}

	path := filepath.Join(LedgerDir, LedgerProvas)
	f, err := os.Open(path)
	if err != nil {
		http.Error(w, "ERR_PROVAS_NAO_ENCONTRADO", http.StatusNotFound)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		linha := strings.TrimSpace(sc.Text())
		if linha == "" {
			continue
		}
		p := strings.Split(linha, "|")
		if len(p) < 14 {
			continue
		}
		seqLinha, err := strconv.Atoi(p[0])
		if err != nil || seqLinha != seq {
			continue
		}
		writeJSON(w, 200, map[string]any{
			"seq":             seqLinha,
			"sha256_evento":   p[1],
			"sha256_payload":  p[2],
			"sha256_conteudo": p[3],
			"source_id":       p[4],
			"de":              p[5],
			"para":            p[6],
			"valor":           p[7],
			"moeda":           p[8],
			"liquidado":       p[9] == "1",
			"liquidado_tx":    p[10],
			"liquidado_utc":   p[11],
			"registrado_utc":  p[12],
			"event_hash":      p[13],
		})
		return
	}
	http.Error(w, "ERR_PROVA_NAO_ENCONTRADA", http.StatusNotFound)
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
	writeJSON(w, 200, out)
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
	p := strings.TrimPrefix(r.URL.Path, "/ledger/")
	if p == "" || strings.Contains(p, "..") || strings.ContainsAny(p, "\\") {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filepath.Join(LedgerDir, p))
}

// ── Main ──────────────────────────────────────────────────────────

func main() {
	if err := loadStateFromDisk(); err != nil {
		log.Fatalf("failed to load state: %v", err)
	}

	iniciarSchedulerLiquidacao()

	mux := http.NewServeMux()
	mux.HandleFunc("/health",          healthHandler)
	mux.HandleFunc("/head",            headHandler)
	mux.HandleFunc("/submit",          submitHandler)
	mux.HandleFunc("/stream",          streamHandler)
	mux.HandleFunc("/prova",           provaHandler)        // ← prova individual
	mux.HandleFunc("/ledger/full.txt", ledgerFullHandler)
	mux.HandleFunc("/ledger/",         ledgerByDateHandler)

	fmt.Println("PoE L2 Gateway rodando em http://localhost:8080")
	fmt.Println("Endpoints:")
	fmt.Println("  POST /submit          — envia transacao")
	fmt.Println("  GET  /prova?seq=N     — consulta prova individual")
	fmt.Println("  GET  /stream?from=N   — stream do ledger")
	fmt.Println("  GET  /health          — status do gateway")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}
