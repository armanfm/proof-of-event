# Proof of Event (PoE)
## Protocolo v0.1 — Fluxo Operacional (Certificação → Append)

**Versão:** 0.1  
**Status:** Operacional (executável no papel)  
**Compatibilidade:** Implementa integralmente SPEC.md v0.1  
**Escopo:** Documento procedural (operacional), não normativo  
**Autor:** Armando Freire  

---

## Estrutura de Documentação

**TIER 1 — README.md**  
→ Visão geral, motivação, stakeholders

**TIER 2 — SPEC.md**  
→ Regras normativas do protocolo (DEVE / NÃO DEVE)

**TIER 3 — Fluxo Operacional (este documento)**  
→ Como o protocolo roda na prática (ordem, estados, falhas)

---

## 1. Objetivo deste Documento

Este documento descreve o **fluxo operacional** do PoE v0.1: como um evento, previamente validado fora do protocolo, é submetido a um Certificador PoE, recebe um timestamp canônico, gera uma prova criptográfica encadeada e é persistido em um ledger append-only.

Este documento **não altera** o SPEC.md.  
Ele apenas remove ambiguidades práticas de implementação.

---

## 2. Entidades

- **Client (Verificador / Oráculo):** prepara e valida o evento fora do PoE.
- **Certificador PoE:** executa o protocolo, certifica temporalmente e mantém o ledger.
- **Ledger do Certificador:** registro append-only das provas emitidas por aquele certificador.

---

## 3. Artefatos

### 3.1 Evento Canônico

Conforme SPEC.md v0.1, contém no mínimo:

- `event_id`
- `client_address`
- `payload_hash_512`
- `verifier_address` (opcional)

O significado do evento **não é avaliado** pelo PoE.

---

### 3.2 Recibo PoE (Receipt)

Ao aceitar um evento, o Certificador retorna um recibo contendo:

- `poe_hash`
- `previous_hash`
- `sequence`
- `payload_hash_512`
- `poe_timestamp_us`
- `version`
- metadados operacionais (opcional)

O recibo é suficiente para **verificação independente**.

---

## 4. Estados Operacionais do Certificador

- **ONLINE:** aceita eventos normalmente  
- **DEGRADED:** aceita eventos com limitações  
- **OFFLINE:** não aceita novos eventos  

Estados operacionais **não afetam** provas já emitidas.

---

## 5. Erros Operacionais (Indicativo)

- `BAD_EVENT`
- `INVALID_EVENT_ID`
- `INVALID_HASH_512`
- `NO_COMMUNITY_FUEL`
- `DUPLICATE_EVENT`
- `RATE_LIMIT`
- `INTERNAL_ERROR`

Os erros são **determinísticos** e não semânticos.

---

## 6. Fluxo Operacional

### Submit → Certify → Append → Receipt

### 6.1 Preparação do Evento (Client)

O Client:
- valida o evento fora do PoE;
- calcula `payload_hash_512` (SHA-512);
- prepara o evento canônico.

---

### 6.2 Submissão (Client → Certificador)

O Client envia:
- evento canônico;
- credenciais/pagamento, se exigido.

O Certificador valida:
- formato;
- deduplicação (`event_id`);
- requisitos operacionais (ex.: fuel).

---

### 6.3 Certificação (Certificador)

O Certificador:
- gera `poe_timestamp_us` (microsegundos UTC);
- computa o hash encadeado:

poe_hash = SHA-256(
event_id |
client_address |
verifier_address |
payload_hash_512 |
poe_timestamp_us |
previous_hash
)

- incrementa `sequence`;
- atualiza `previous_hash`.

---

### 6.4 Append no Ledger

O ledger é:
- append-only;
- imutável;
- local ao Certificador.

#### Formato SELF

client_address|payload_hash_512|verifier_address|poe_timestamp_us


---

## 7. Ordem do Ledger (Esclarecimento)

A ordem no ledger:
- NÃO representa ordem real;
- NÃO implica causalidade;
- NÃO estabelece precedência jurídica.

Ela representa **exclusivamente a ordem de aceitação local**.

---

## 8. Armazenamento Externo (Opcional)

Dados completos podem ser armazenados externamente (IPFS, Pinata etc.).

- não fazem parte do protocolo;
- não interferem na prova;
- o ledger é a fonte autoritativa.

---

## 9. Falhas Realistas

### 9.1 Certificador Offline
- não aceita novos eventos;
- histórico permanece válido.

### 9.2 Client Offline
- pode reenviar;
- deduplicação evita cobrança dupla.

---

## 10. Eventos de Teste e Conteúdo Inválido

Se:
- formato é válido;
- requisitos operacionais são atendidos;

→ o evento é aceito.

Correções ocorrem **apenas por novos eventos**.

---

## 11. Encerramento

Clients verificam.  
Certificadores testemunham.  
Ledgers preservam.

O Proof of Event não decide.  
Ele registra fatos criptográficos.

**A blockchain não decide. Ela testemunha.**



