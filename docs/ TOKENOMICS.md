# TOKENOMICS v0.1 — Proof of Event (PoE)
Implementação Econômica Normativa (Contabilidade Interna) + Congestion Fee

**Versão:** 0.1  
**Status:** Normativo (regras mecânicas executáveis)  
**Compatível com:** SPEC.md v0.1 + protocol/v0.1.md  
**Autor:** Armando Freire  
**Licença:** Apache 2.0  

---

## 0. Princípio

O PoE **não decide conteúdo** e **não interpreta eventos**.  
O PoE **testemunha**: ordem FIFO + ledger append-only.

A economia do PoE é **contabilidade interna determinística**:  
cobrança e distribuição de POE obedecem regras fixas, auditáveis e reexecutáveis.

---

## 1. Ativo Nativo

- **Nome:** Token PoE  
- **Símbolo:** POE  
- **Decimais:** 18  
- **Governança:** nenhuma  

---

## 2. Oferta (Supply)

- **Supply total (genesis):** `1.000.000.000.000 POE`  
- **Mint adicional:** proibido em v0.1  

Em v0.1, a emissão é **fechada**:  
não existe criação infinita por evento.

---

## 3. Identidades Mecânicas (Sem Identidade Civil)

Para contabilidade interna automática:

- **storer_id:** `bytes32`  
  Identifica o nó armazenador que realizou o append do ledger.

- **payer_id:** `bytes32`  
  Identifica quem consome POE para submeter eventos  
  (verificador, cliente ou Plataforma).

Esses IDs são **puramente mecânicos**, não representam identidade civil.

---

## 4. Regra Central — Cobrança e Distribuição por Evento

### 4.1 Taxas fixas por evento aceito (v0.1)

- **FEE_PLATFORM = 1 POE**  
- **FEE_STORER   = 1 POE**


TOTAL_FEE_BASE = 2 POE

## Normativo

- O verificador **NÃO** recebe **POE**.
- O verificador é **consumidor do protocolo**, não agente remunerado.

---

## 4.2 Momento exato da liquidação

A liquidação ocorre **somente quando**:

1. o evento foi aceito pelo **FIFO**  
   *(formato + `previous_event_hash` + ordem FIFO)*, **e**
2. o evento foi efetivamente gravado no **ledger** *(append confirmado)*.

Se falhar **antes disso**:

- não há débito;
- não há crédito;
- não há consumo de POE.

---

## 4.3 Transferências determinísticas (contabilidade interna)

Ao aceitar e gravar um evento:

- debitar `payer_id` em `TOTAL_FEE_BASE`

- creditar:
  - `platform_id` com **1 POE**
  - `storer_id` com **1 POE**

### Normativo

Se `payer_id` não tiver saldo suficiente, a submissão **DEVE** ser recusada **antes do append**, com:

- `ERR_INSUFFICIENT_BALANCE`

---

## 5. Congestion Fee (Taxa de Congestionamento)

A **congestion fee** é uma taxa adicional **exclusiva da Plataforma**.

### Normativo

- congestion fee **não altera ordem**
- congestion fee **não compra prioridade**
- congestion fee vai **100% para a Plataforma**

---

## 5.1 Cálculo

A implementação pode usar função baseada em `queue_size`, desde que:

- `queue_size` seja o número de submissões pendentes no FIFO no instante da recepção;
- o cálculo seja determinístico e auditável.

---

## 5.2 Liquidação com congestion fee

`TOTAL_FEE = TOTAL_FEE_BASE + CONGESTION_FEE`

Repasse:

- `platform_id += 1 POE + CONGESTION_FEE`
- `storer_id += 1 POE`

---

## 6. Limites de Payload (Hard Limits)

Mesmo pagando, existem limites técnicos:

- `MAX_EVENT_JSON_BYTES = 8.192` (8 KiB)
- `MAX_BLOB_BYTES = 5.242.880` (5 MiB)

Acima disso:

- `ERR_PAYLOAD_TOO_LARGE`

---

## 7. Auditoria Pública (Obrigatória)

A Plataforma **DEVE** publicar relatório auditável contendo:

- número de eventos aceitos
- total debitado
- total creditado para:
  - Plataforma
  - Armazenadores
- hash SHA-256 do relatório

### Normativo

O hash do relatório **DEVE** ser registrado como evento no PoE *(payload_hash do relatório)*.

---

## 8. Parâmetros e Mudanças de Regra

Em v0.1:

- `FEE_PLATFORM = 1 POE`
- `FEE_STORER = 1 POE`
- congestion fee **não compra prioridade**
- congestion fee é **100% Plataforma**

Qualquer mudança:

- exige versionamento (**v0.2+**)
- exige registro do hash do novo arquivo de parâmetros no PoE

---

## 9. Erros Mínimos

- `ERR_BAD_FORMAT`
- `ERR_BAD_PREV_HASH`
- `ERR_PAYLOAD_TOO_LARGE`
- `ERR_INSUFFICIENT_BALANCE`
- `ERR_INTERNAL`

---

## 10. Avisos (Sem Promessa)

O token **POE**:

- não garante valorização
- não é governança
- não é dividendo
- não opera exchange

---

## 11. Separação entre Pagamento e Consumo

O POE é exclusivamente uma unidade técnica de consumo do protocolo.

Pagamentos pelo uso do PoE:

- ocorrem fora do protocolo
- podem ser feitos em fiat ou cripto
- não criam tokens
- não concedem direitos on-chain

A Plataforma:

- converte pagamento externo em consumo de POE
- assume risco operacional e cambial

---

## 12. Encerramento

Em v0.1:

- quem usa, paga **2 POE** (+ congestion fee se houver)
- quem trabalha, recebe automaticamente (Plataforma e Armazenador)
- congestionamento protege a infraestrutura
- o PoE permanece minimalista: testemunha e registra

**Quem usa, paga.**  
**Quem trabalha, recebe.**  
**Congestionamento protege a infraestrutura — não compra prioridade.**




