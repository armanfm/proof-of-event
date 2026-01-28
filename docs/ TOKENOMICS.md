# TOKENOMICS v0.1 — Proof of Event (PoE)
Implementação Econômica Normativa (Contabilidade Interna) + Congestion Fee

**Versão:** 0.1  
**Status:** Normativo (regras mecânicas executáveis)  
**Compatível com:** `SPEC.md` v0.1 + `protocol/v0.1.md`  
**Autor:** Armando Freire  
**Licença:** Apache 2.0  

---

## 0. Princípio
O PoE **não decide conteúdo** e **não interpreta eventos**.  
O PoE **testemunha**: ordem FIFO + ledger append-only.

A economia do PoE é **contabilidade interna determinística**: cobrança e distribuição de POE obedecem regras fixas e auditáveis.

---

## 1. Ativo Nativo
- **Nome:** Token PoE  
- **Símbolo:** `POE`  
- **Decimais:** `18`  
- **Governança:** nenhuma  

---

## 2. Oferta (Supply)
- **Supply total (genesis):** `1.000.000.000.000 POE`  
- **Mint adicional:** proibido em v0.1  

> Em v0.1, a emissão é “fechada”: não existe criação infinita por evento.

---

## 3. Identidades Mecânicas (Sem Identidade Civil)
Para distribuição automática:
- **verifier_id:** `bytes32` (ex.: hash da chave/identidade do verificador) — em v0.1, é o `source_id` do evento.
- **storer_id:** `bytes32` (hash da identidade do nó armazenador) — em v0.1, é o ID do nó que realizou o append do ledger.
- **payer_id:** `bytes32` (quem paga para submeter; pode ser o próprio verificador ou um cliente).

> Isso é ID mecânico para contabilidade, não “identidade civil”.

---

## 4. Regra Central — Cobrança e Distribuição por Evento
Em v0.1, **cada evento aceito no FIFO** executa **uma liquidação determinística**:

### 4.1 Taxas fixas por evento aceito
- **FEE_PLATFORM = 1 POE**
- **FEE_STORER   = 1 POE**
- **FEE_VERIFIER = 1 POE**

Logo:

`TOTAL_FEE_BASE = 3 POE`

### 4.2 Momento exato da liquidação
A liquidação **só acontece** quando:
1) o evento foi **aceito** (passou em formato + prev_hash + FIFO), e  
2) o evento foi **gravado no ledger** (append confirmado no nó).

Se falhar antes disso, **não existe débito nem crédito**.

### 4.3 Transferências determinísticas (contabilidade interna)
Ao aceitar e gravar um evento:
- debitar `payer_id` em `TOTAL_FEE_BASE` (mais congestion fee se aplicável)
- creditar:
  - `platform_id` com `1 POE`
  - `storer_id` com `1 POE`
  - `verifier_id` com `1 POE`

**Normativo:** se `payer_id` não tiver saldo suficiente, a submissão **DEVE** ser recusada antes do append, com `ERR_INSUFFICIENT_BALANCE`.

---

## 5. Congestion Fee (Taxa de Congestionamento)
A congestion fee é uma taxa adicional **exclusiva da Plataforma**, aplicada quando a fila está congestionada.

**Normativo:**
- A congestion fee **não altera ordem**.
- A congestion fee **não compra prioridade**.
- A congestion fee **vai 100% para a Plataforma**.

### 5.1 Cálculo (mecânico)
A implementação pode usar uma função baseada em `queue_size` (backlog pendente), respeitando:
- `queue_size` é o número de submissões **pendentes de aceitação** no FIFO no instante de recebimento.
- o valor final deve ser determinístico e auditável.

### 5.2 Liquidação com congestion fee
`TOTAL_FEE = TOTAL_FEE_BASE + CONGESTION_FEE`

E o repasse fica:
- `platform_id += 1 POE + CONGESTION_FEE`
- `storer_id   += 1 POE`
- `verifier_id += 1 POE`

---

## 6. Limites de Payload (Hard Limits)
Mesmo pagando, existem limites para evitar abuso:

- `MAX_EVENT_JSON_BYTES = 8.192` (8 KiB)
- `MAX_BLOB_BYTES = 5.242.880` (5 MiB), se houver blob

Acima disso: rejeitar com **`ERR_PAYLOAD_TOO_LARGE`**.

---

## 7. Auditoria Pública (Obrigatória)
A Plataforma **DEVE** publicar relatório auditável periódico contendo:
- número de eventos aceitos no período
- total debitado
- total creditado para `platform_id`, `storers`, `verifiers`
- saldo agregado por classe (pode ser por ID se você quiser transparência total)
- hash `SHA-256` do relatório

**Normativo:** o hash do relatório **DEVE** ser registrado como evento no PoE (`payload_hash` do relatório).

---

## 8. Parâmetros e Mudanças de Regra
Em v0.1:
- `FEE_PLATFORM = 1 POE`
- `FEE_STORER   = 1 POE`
- `FEE_VERIFIER = 1 POE`
- congestion fee não compra prioridade e é 100% plataforma

**Normativo (anti-mudança silenciosa):**
- Se qualquer parâmetro econômico mudar, isso **exige versionamento** (`v0.2+`), e a Plataforma **DEVE** registrar no PoE o hash do arquivo de parâmetros ativo.

---

## 9. Erros mínimos (consistência com protocol)
- `ERR_BAD_FORMAT`
- `ERR_BAD_PREV_HASH`
- `ERR_PAYLOAD_TOO_LARGE`
- `ERR_INSUFFICIENT_BALANCE`
- `ERR_INTERNAL`

---

## 10. Avisos (Sem promessa)
O token:
- não garante valorização, retorno, liquidez
- não é governança
- não é dividendo
- não opera exchange

---

## 11. Encerramento
Em v0.1:
- quem usa, paga `3 POE (+ congestion fee se houver)`
- quem trabalha, recebe automaticamente (plataforma, armazenador, verificador)
- congestionamento protege infraestrutura e **não compra prioridade**
- o PoE continua minimalista: **testemunha e registra**


**Quem usa, paga.**  
**Quem trabalha, recebe.**  
**Congestionamento protege a infraestrutura — não compra prioridade.**


