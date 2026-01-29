# Proof of Event (PoE)
## Protocolo v0.1 — Fluxo Operacional (Certificação → Append)

**Versão:** 0.1  
**Status:** Operacional (executável no papel)  
**Compatibilidade:** Implementa integralmente `SPEC.md` v0.1, sem extensões  
**Escopo:** Documento procedural (operacional), não normativo  
**Autor:** Armando Freire  

---

## Estrutura de Documentação

**TIER 1 — README.md**  
→ O que é o PoE (visão geral, stakeholders, motivação)

**TIER 2 — SPEC.md**  
→ Regras normativas do protocolo (o que DEVE / NÃO DEVE)

**TIER 3 — Fluxo Operacional (este documento)**  
→ Como o protocolo roda na prática (operadores, implementação)

---

## 1. Objetivo deste Documento

Este documento define o **fluxo operacional** do PoE v0.1: como um evento,
produzido por um verificador/oráculo, é recebido por um **Certificador PoE**,
recebe um **timestamp canônico**, gera uma **prova PoE** e é persistido em um
ledger **append-only**.

Este documento **não altera** o `SPEC.md`.
Ele apenas remove ambiguidades de implementação:
ordem de chamadas, estados operacionais, erros e comportamento sob falhas.

---

## 2. Entidades

- **Client (Verificador / Oráculo):**
  Prepara o evento, valida off-chain e submete ao Certificador PoE.

- **Certificador PoE:**
  Executa o protocolo PoE, atribui timestamp canônico, gera a prova e mantém
  o ledger determinístico.

- **Ledger do Certificador:**
  Registro append-only das provas emitidas por aquele certificador.

---

## 3. Artefatos

### 3.1 Evento Canônico

Conforme definido no `SPEC.md` v0.1, o evento canônico contém, no mínimo:

- `payload_hash`
- metadados de versão (quando aplicável)

O significado do evento **não é avaliado** pelo PoE.

---

### 3.2 Recibo PoE (Receipt)

Ao aceitar um evento, o Certificador PoE **DEVE** retornar um recibo contendo:

- `poe_proof`
- `payload_hash`
- `timestamp_canônico`
- `certificador_id`
- `version`
- `metadata` (opcional, fora da prova)

O recibo é suficiente para verificação independente.

---

## 4. Estados Operacionais do Certificador

O Certificador PoE pode operar nos seguintes estados:

- **ONLINE:** aceita novos eventos normalmente.
- **DEGRADED:** aceita eventos, mas com limitações operacionais.
- **OFFLINE:** não aceita novos eventos.

**Observação:**  
Estados operacionais **NÃO alteram** a validade das provas já emitidas.

---

## 5. Códigos de Erro (Normativo)

O Certificador **DEVE** responder com erros determinísticos:

- `ERR_BAD_FORMAT` — formato inválido
- `ERR_BAD_VERSION` — versão não suportada
- `ERR_NO_TOKEN` — pagamento exigido e não efetuado
- `ERR_DUPLICATE_EVENT` — evento duplicado (deduplicação)
- `ERR_CERTIFIER_OFFLINE` — certificador indisponível
- `ERR_INTERNAL` — erro interno (uso mínimo)

---

## 6. Fluxo Operacional  
### Submit → Certify → Append → Receipt

### 6.1 Preparação do Evento (Client)

O Client:

- valida o evento fora do PoE (Camada 1);
- gera `payload_hash` (ex.: SHA-512);
- prepara o evento canônico.

---

### 6.2 Submissão ao Certificador (Client → Certificador)

O Client envia:

- evento canônico;
- pagamento, se exigido pela operação.

O Certificador valida:

- formato canônico;
- versão suportada;
- deduplicação por `payload_hash` ou `event_id` (se aplicável);
- requisitos operacionais (ex.: pagamento).

Se aprovado, o evento é aceito.

---

### 6.3 Certificação (Certificador)

O Certificador:

1. gera o **timestamp canônico** (UTC);
2. calcula a prova PoE:

PoE_Proof = SHA-512(payload_hash || timestamp_canônico)

3. registra a prova no ledger append-only;
4. emite o recibo PoE.

---

### 6.4 Append no Ledger

O ledger do Certificador:

- é append-only;
- nunca remove eventos;
- nunca reescreve histórico.

#### Ordem do Ledger (Esclarecimento Importante)

A ordem no ledger:

- **NÃO** representa ordem de ocorrência no mundo real;
- **NÃO** implica causalidade;
- **NÃO** estabelece precedência jurídica externa.

Ela representa **exclusivamente a ordem local de aceitação**
por aquele Certificador PoE.

---

## 7. Armazenamento Externo (Opcional)

Dados completos podem ser armazenados externamente
(ex.: IPFS, Pinata, bancos institucionais).

Esses sistemas:

- são opcionais;
- não fazem parte do protocolo PoE;
- não interferem na prova criptográfica.

### 7.1 Ledger vs. Armazenamento Externo (Normativo)

O ledger do Certificador **É A FONTE PRIMÁRIA E AUTORITATIVA**.

Armazenamento externo é apenas:

- replicação;
- conveniência ao Client;
- distribuição de dados completos.

Se o armazenamento externo falhar:

- ✅ Provas PoE permanecem válidas
- ✅ Verificação pode ser feita via ledger
- ❌ Apenas metadados opcionais podem ser perdidos

---

## 8. Falhas Realistas

### 8.1 Certificador Offline

- não aceita novos eventos;
- histórico existente permanece válido;
- nenhuma prova é invalidada.

### 8.2 Client Offline

- problema do Client;
- pode reenviar quando retornar;
- deduplicação evita cobrança dupla.

**Não há promessa de 100% uptime.**

---

## 9. Eventos de Teste e Conteúdo “Errado”

O PoE aceita qualquer evento que:

- respeite o formato;
- cumpra requisitos operacionais.

Se pagou e seguiu o formato, **é aceito**.

Neutralidade é mecânica, não discursiva.

Correções ocorrem apenas por **novo evento**.

---

## 10. Interface Mínima (Sugerida)

- `POST /submit` — submete evento e retorna recibo ou erro
- `GET /head` — retorna estado atual do ledger
- `GET /ledger` — leitura pública do histórico

---

## 11. Encerramento

Este fluxo operacional define um sistema simples e honesto:

**Clients verificam. Certificadores testemunham. Ledgers preservam.**

O Proof of Event não decide.
Ele registra fatos criptográficos.

**A blockchain não decide. Ela testemunha.**

---

## Apêndice A — Sessão de Certificação Completa

### A.1 Submissão (Client)

Evento externo:  
“Contrato XYZ assinado”

Hash off-chain:  
SHA-512("Contrato XYZ") = `a1b2c3...`

---

### A.2 Certificação

Timestamp canônico:  
`2024-01-01T12:00:00.000Z`

Prova:

SHA-512(a1b2c3... || 2024-01-01T12:00:00.000Z)


---

### A.3 Recibo


{
  "poe_proof": "d8e7f2a1...",
  "payload_hash": "a1b2c3...",
  "timestamp_canônico": "2024-01-01T12:00:00.000Z",
  "certificador_id": "poe:c6c1024c...",
  "version": "0.1"
}

