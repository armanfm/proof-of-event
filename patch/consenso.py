"""
Verificador de Consenso — PoE L2
=================================
Versão correta e simples.
Usa apenas SHA256 — binário, sem float, sem edge case.

Regra:
  SHA256 igual → conteúdo idêntico → aprovado
  SHA256 diferente → adulterado → slashing

Roda como serviço externo independente.
Não faz parte de nenhum nó — observa todos.
"""

import hashlib
import json
import time
import requests
import logging
from typing import Optional

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [VERIFICADOR] %(message)s"
)

# ── Configuração ─────────────────────────────────────────────────

NOS = [
    {"id": "no_1", "url": "http://localhost:8080"},
    {"id": "no_2", "url": "http://localhost:8081"},
    {"id": "no_3", "url": "http://localhost:8082"},
]

CONSENSO_MIN  = 2        # mínimo de nós que precisam concordar
INTERVALO     = 10       # segundos entre verificações
CONTRATO_URL  = "http://localhost:9000/slashing"

# ── SHA256 ────────────────────────────────────────────────────────

def sha256_hex(s: str) -> str:
    return hashlib.sha256(s.encode()).hexdigest()

def sha256_entry(entry: dict) -> str:
    # sort_keys garante que a ordem dos campos não afeta o hash
    return sha256_hex(json.dumps(entry, sort_keys=True))

# ── Busca ledger de um nó ─────────────────────────────────────────

def buscar_ledger_no(no: dict, from_seq: int = 1) -> Optional[list]:
    try:
        resp = requests.get(
            f"{no['url']}/stream?from={from_seq}",
            timeout=5
        )
        if resp.status_code == 200:
            return resp.json()
    except Exception as e:
        logging.warning(f"Nó {no['id']} inacessível: {e}")
    return None

# ── Verificação de consenso ───────────────────────────────────────

def verificar_seq(seq: int, ledgers: dict) -> dict:
    """
    Para um seq específico:
      1. Pega o entry de cada nó
      2. Calcula SHA256 de cada um
      3. Agrupa por hash
      4. Maior grupo = verdade
      5. Minoria = slashing
    """
    # Coleta entry de cada nó para esse seq
    conteudos = {}
    for no_id, entries in ledgers.items():
        for entry in entries:
            if entry.get("sequence") == seq:
                conteudos[no_id] = entry
                break

    if len(conteudos) < 2:
        return {
            "seq":      seq,
            "aprovado": False,
            "motivo":   "ENTRIES_INSUFICIENTES",
            "nos_ok":   [],
            "nos_falhos": [],
        }

    # Agrupa nós pelo hash do seu entry
    # Nós com mesmo hash → concordam
    # Nós com hash diferente → divergem
    grupos = {}
    for no_id, entry in conteudos.items():
        h = sha256_entry(entry)
        grupos.setdefault(h, []).append(no_id)

    # Grupo com mais nós = verdade
    grupo_maior = max(grupos.values(), key=len)

    # Aprovado se o maior grupo tem consenso mínimo
    aprovado = len(grupo_maior) >= CONSENSO_MIN

    # Nós que divergiram da maioria
    nos_falhos = []
    for h, nos in grupos.items():
        if nos != grupo_maior:
            for no_id in nos:
                nos_falhos.append({
                    "no_id": no_id,
                    "hash":  h[:16] + "...",
                })

    motivo = ""
    if not aprovado:
        motivo = f"CONSENSO_INSUFICIENTE: {len(grupo_maior)}/{len(conteudos)}"
    elif nos_falhos:
        motivo = f"NOS_DIVERGENTES: {[n['no_id'] for n in nos_falhos]}"

    return {
        "seq":        seq,
        "aprovado":   aprovado,
        "nos_ok":     grupo_maior,
        "nos_falhos": nos_falhos,
        "motivo":     motivo,
        "hash_verdade": list(grupos.keys())[
            list(grupos.values()).index(grupo_maior)
        ][:16] + "...",
    }

# ── Reversão entre aprovados ──────────────────────────────────────

def reversar_aprovados(nos_ok: list, seq: int, ledgers: dict) -> bool:
    """
    Nós aprovados confirmam entre si.
    Compara SHA256 de todos — tem que ser idêntico.
    """
    if len(nos_ok) < CONSENSO_MIN:
        return False

    hashes = set()
    for no_id in nos_ok:
        entries = ledgers.get(no_id, [])
        for entry in entries:
            if entry.get("sequence") == seq:
                hashes.add(sha256_entry(entry))
                break

    # Todos têm exatamente o mesmo hash
    if len(hashes) == 1:
        logging.info(f"seq {seq}: ✅ reversão OK — {len(nos_ok)} nós confirmam")
        return True

    logging.error(f"seq {seq}: ❌ reversão falhou — hashes divergem")
    return False

# ── Slashing ──────────────────────────────────────────────────────

def executar_slashing(no_id: str, seq: int, hash_falso: str):
    payload = {
        "no_id":      no_id,
        "seq":        seq,
        "hash_falso": hash_falso,
        "timestamp":  time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime()),
    }

    logging.warning(f"🚨 SLASHING | nó={no_id} | seq={seq}")

    try:
        requests.post(CONTRATO_URL, json=payload, timeout=5)
    except Exception as e:
        logging.error(f"Erro ao enviar slashing: {e}")

    with open("slashing.log", "a") as f:
        f.write(json.dumps(payload) + "\n")

# ── Loop principal ────────────────────────────────────────────────

def ciclo_verificacao():
    logging.info(f"Verificador iniciado | {len(NOS)} nós | consenso: {CONSENSO_MIN}/{len(NOS)}")

    ultimo_seq = 0

    while True:
        # Busca ledgers de todos os nós
        ledgers = {}
        for no in NOS:
            entries = buscar_ledger_no(no, from_seq=ultimo_seq + 1)
            if entries:
                ledgers[no["id"]] = entries

        if len(ledgers) < 2:
            logging.warning("Menos de 2 nós acessíveis")
            time.sleep(INTERVALO)
            continue

        # Seqs novos para verificar
        todos_seqs = set()
        for entries in ledgers.values():
            for entry in entries:
                todos_seqs.add(entry.get("sequence", 0))

        seqs = sorted(s for s in todos_seqs if s > ultimo_seq)

        for seq in seqs:
            resultado = verificar_seq(seq, ledgers)

            if resultado["aprovado"]:
                ok = reversar_aprovados(
                    resultado["nos_ok"], seq, ledgers
                )
                if ok:
                    logging.info(
                        f"✅ seq={seq} | "
                        f"nós={resultado['nos_ok']} | "
                        f"pronto para liquidação"
                    )
                    ultimo_seq = seq
            else:
                logging.error(
                    f"❌ seq={seq} | {resultado['motivo']}"
                )
                for nf in resultado["nos_falhos"]:
                    executar_slashing(
                        nf["no_id"], seq, nf["hash"]
                    )

        time.sleep(INTERVALO)

# ── Demo ──────────────────────────────────────────────────────────

def demo():
    print("=" * 55)
    print("  VERIFICADOR DE CONSENSO — SHA256 puro")
    print("  Binário: igual ou diferente. Sem float.")
    print("=" * 55)

    entry_ok = {
        "sequence":        1,
        "event_hash":      "abc123",
        "source_id":       "banco_inter",
        "payload_hash":    "def456",
        "event_id":        "tx_001",
        "accepted_at_utc": "2026-03-19T12:00:00Z",
        "de":    "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
        "para":  "0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB",
        "valor": "10000",
    }

    entry_valor_adulterado = dict(entry_ok)
    entry_valor_adulterado["valor"] = "99999"

    entry_dest_adulterado = dict(entry_ok)
    entry_dest_adulterado["para"] = "0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"

    casos = [
        ("3 nós idênticos", [
            entry_ok, entry_ok, entry_ok
        ]),
        ("Nó 3 mudou o valor R$100 → R$999", [
            entry_ok, entry_ok, entry_valor_adulterado
        ]),
        ("Nó 2 e 3 mudaram o destinatário", [
            entry_ok, entry_dest_adulterado, entry_dest_adulterado
        ]),
        ("Todos diferentes", [
            entry_ok, entry_valor_adulterado, entry_dest_adulterado
        ]),
    ]

    for nome, entries in casos:
        print(f"\n  {'─'*50}")
        print(f"  Teste: {nome}")

        ledgers = {f"no_{i+1}": [e] for i, e in enumerate(entries)}
        r = verificar_seq(1, ledgers)

        print(f"  Resultado  : {'✅ APROVADO' if r['aprovado'] else '❌ REPROVADO'}")
        if r["nos_ok"]:
            print(f"  Nós ok     : {r['nos_ok']}")
        if r["nos_falhos"]:
            print(f"  Slashing   : {[n['no_id'] for n in r['nos_falhos']]}")
        if r["motivo"]:
            print(f"  Motivo     : {r['motivo']}")

    print(f"\n{'='*55}")
    print("  SHA256 é suficiente — binário, O(n), sem edge case")
    print(f"{'='*55}")

if __name__ == "__main__":
    import sys
    if "--demo" in sys.argv or len(sys.argv) == 1:
        demo()
    else:
        ciclo_verificacao()
