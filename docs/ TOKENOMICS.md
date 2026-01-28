# TOKENOMICS v0.1 — Proof of Event (PoE)
Implementação Econômica Normativa (Off-chain) com Congestion Fee

**Versão:** 0.1  
**Status:** Normativo (regras mecânicas executáveis)  
**Compatível com:** `SPEC.md` v0.1 + `protocol/v0.1.md`  
**Autor:** Armando Freire 
**Licença:** Apache 2.0  

---

## 0. Escopo e Separação de Responsabilidades

Este documento define **parâmetros numéricos** e **regras mecânicas** para:
- cobrança de taxas (fees) para submissão no FIFO;
- queima (burn) obrigatória;
- redistribuição operacional;
- taxa adicional de congestionamento destinada exclusivamente à Plataforma;
- regras de medição para pagamentos por trabalho (epochs e units).

### 0.1 O que é PoE Core (Camada 2)
O **PoE Core** (FIFO + ledger append-only) tem responsabilidade mínima:
- exigir **formato** + **encadeamento** (`previous_event_hash`) + **ordem FIFO**;
- registrar e distribuir eventos em ordem;
- armazenadores replicam e mantêm o ledger.

O PoE Core **não executa** queima, redistribuição, payout, exchange ou preço.

### 0.2 O que é a Camada Econômica (Off-chain)
A **camada econômica off-chain** (implementação/serviço externo) é quem:
- verifica saldo e prova de pagamento para entrada no FIFO;
- executa burn;
- calcula pools e payouts por epoch;
- publica relatórios auditáveis e registra o hash do relatório no PoE.

> O token não promete lucro, não define preço e não opera exchange.

---

## 1. Ativo Nativo

- **Nome:** Token PoE  
- **Símbolo:** `POE`  
- **Decimais:** `18`  
- **Tipo:** criptomoeda / ativo digital de liquidação operacional  
- **Governança:** nenhuma

---

## 2. Oferta (Supply)

### 2.1 Supply Total (Fixado no Genesis)
- **Supply total:** `1.000.000.000 POE`
- Criado **uma única vez** no genesis.
- **Nenhum mint adicional** é permitido em v0.1.

### 2.2 Queima Reduz Supply
- Burn é destruição irreversível de tokens (supply diminui).
- Não existe “nascer mais token” após queima.

---

## 3. Alocação Genesis (Distribuição Inicial)

Buckets fixos (percentual do supply inicial):

## 2. Oferta (Supply)

### 2.1 Supply Total (Fixado no Genesis)
- **Supply total:** `1.000.000.000.000 POE`
- Criado **uma única vez** no genesis.
- **Nenhum mint adicional** é permitido em v0.1.

---

## 3. Alocação Genesis (Distribuição Inicial)

Buckets fixos (percentual do supply inicial):

1. **STORER_FOUNDERS:** `50%` = `500.000.000.000 POE`  
2. **PLATFORM_RESERVE:** `30%` = `300.000.000.000 POE`  
3. **VERIFIER_FOUNDERS:** `20%` = `200.000.000.000 POE`

### 3.1 Vesting (v0.1)
- **STORER_FOUNDERS (500B):** 25% no genesis + 75% linear em 12 meses (diário)  
- **VERIFIER_FOUNDERS (200B):** 25% no genesis + 75% linear em 12 meses (diário)  
- **PLATFORM_RESERVE (300B):** 40% no genesis + 60% linear em 36 meses (diário)

Regra: tokens travados **não contam** como disponíveis para pagamentos até liberados.

> Em v0.1 não existe bucket DEV_FUND separado (se você quiser DEV_FUND, tem que tirar percentual de alguém).



Regra: tokens travados **não contam** como disponíveis para pagamentos até liberados.

---

## 4. Identidades de Trabalho (Sem Governança)

A redistribuição exige identificar “quem trabalhou”.

Em v0.1:
- **verifier_id:** `bytes32` (hash de chave pública/identidade do verificador)  
- **storer_id:** `bytes32` (hash de chave pública/identidade do armazenador)

Uso:
- `verifier_id` aparece no evento (ex.: `source_id`)  
- `storer_id` aparece na `Commitment_Proof`

> Isso não é “identidade civil”. É apenas ID mecânico de remuneração.

---

## 5. Estrutura de Taxas (Fee Model)

A taxa total para entrada no FIFO:

1. **Base Fee** (registro no FIFO)  
2. **Blob Fee** (peso físico por bytes, se houver blob)  
3. **Retention Fee** (multiplicador por tempo de retenção do blob)  
4. **Congestion Fee** (pico de demanda / fila cheia)

### 5.1 Unidades
- Todas as taxas são cobradas em **POE** (`10^-18 POE`).
- Arredondamento: **sempre para cima** (ceil) na unidade mínima.

---

## 6. Parâmetros Numéricos Fixos (v0.1)

### 6.1 Base Fee
- **BASE_FEE:** `0,10 POE` por evento aceito no FIFO.

### 6.2 Blob Fee (por tamanho)
- **BLOB_FEE_PER_KIB:** `0,002 POE` por `1 KiB` (1024 bytes)
- `blob_kib = ceil(blob_bytes / 1024)`
- `blob_fee = blob_kib * 0,002 POE`

### 6.3 Retention Fee (multiplicador)
- `RETENTION_30D_MULT = 1x`
- `RETENTION_365D_MULT = 2x`
- `RETENTION_PERMANENT_MULT = 6x`

`retained_blob_fee = blob_fee * retention_multiplier`

### 6.4 Hard Limits (mesmo pagando)
- `MAX_BLOB_BYTES = 5.242.880` (5 MiB)
- `MAX_EVENT_JSON_BYTES = 8.192` (8 KiB)

Acima disso: rejeitar com **`ERR_PAYLOAD_TOO_LARGE`**.

> **Normativo:** adicione `ERR_PAYLOAD_TOO_LARGE` na lista de erros do `protocol/v0.1.md` para ficar consistente.

### 6.5 Congestion Fee (Taxa de Congestionamento)
Parâmetros:
- `Q0 = 100`
- `QMAX = 10.000`
- `CONGESTION_CAP_MULT = 20x`

**Definição normativa de `queue_size`:**  
`queue_size` é o número de submissões **pendentes de aceitação** no FIFO no instante em que a submissão é recebida (backlog atual da fila do gateway).

Regra:
- Se `queue_size <= Q0` ⇒ `congestion_fee = 0`
- Se `queue_size > Q0` ⇒


q = min(queue_size, QMAX)
mult = ceil_div((q - Q0), Q0)      // 1x, 2x, 3x...
mult = min(mult, CONGESTION_CAP_MULT)
congestion_fee = BASE_FEE * mult

## Normativo (Congestion Fee)

- A congestion fee **não altera ordem**.
- A congestion fee **não compra prioridade**.
- A congestion fee **vai 100% para a Plataforma**.

---

## 7. Taxa Total a Pagar (Entrada no FIFO)

Defina:

- `protocol_fee = BASE_FEE + retained_blob_fee`
- `total_fee = protocol_fee + congestion_fee`

A camada econômica **DEVE** recusar a submissão se não houver saldo/prova para `total_fee`.

O mecanismo de pagamento é **off-chain**. O FIFO só recebe “prova/ok” da camada econômica conforme implantação.

---

## 8. Liquidação Econômica (Burn + Redistribuição) — Off-chain

### 8.1 Separação Obrigatória

A liquidação separa a taxa em duas partes:

- **Protocol Fee** (`protocol_fee`)
- **Congestion Fee** (`congestion_fee`)

Burn e redistribuição aplicam somente sobre `protocol_fee`.  
`congestion_fee` vai integralmente para Plataforma.

### 8.2 Burn

- `BURN_RATE = 10%` sobre `protocol_fee`
- `burn_amount = protocol_fee * 0,10`
- `distributable = protocol_fee - burn_amount`

Burn é obrigatório/irrevogável na camada econômica.

### 8.3 Redistribuição (sobre `distributable`)

- **STORERS:** `40%`
- **VERIFIERS:** `30%`
- **PLATFORM:** `20%`

Cálculo:

- `to_storers   = distributable * 0,40`
- `to_verifiers = distributable * 0,30`
- `to_platform  = distributable * 0,20`
- `to_platform += congestion_fee`

---

## 9. Medição de Trabalho (Pagamentos por Epoch)

Pagamentos ocorrem por **epochs** para reduzir custo operacional.

### 9.1 Epoch

- `EPOCH_LENGTH = 24h`
- Fuso: `UTC`
- `epoch_id = YYYY-MM-DD (UTC)`

### 9.2 Unidades — Verificadores

Cada evento aceito no FIFO conta **1 unidade** para o `verifier_id` do evento.

### 9.3 Unidades — Armazenadores

Um armazenador ganha unidade quando:

1. recebeu o evento do FIFO (sequência canônica),
2. realizou append no ledger local,
3. emitiu `Commitment_Proof` válida.

`Commitment_Proof` é opcional no PoE Core, mas obrigatória para remuneração nesta tokenomics v0.1.

### 9.4 Elegibilidade — Retenção Completa

Para receber pagamento como armazenador, o nó deve:

- manter ledger completo desde GENESIS até head (segmentado por dia, contínuo);
- estar sincronizado no hash canônico do ledger.

---

## 10. Pools e Payouts por Epoch

Para cada epoch:

- `pool_storers(epoch)` = soma de `to_storers`
- `pool_verifiers(epoch)` = soma de `to_verifiers`
- `pool_platform(epoch)` = soma de `to_platform` + todas as `congestion_fee`

### 10.1 Payout Pro-rata

- `payout(storer_id) = pool_storers * units(storer_id) / total_units_storers`
- `payout(verifier_id) = pool_verifiers * units(verifier_id) / total_units_verifiers`

Plataforma recebe `pool_platform` integral.

### 10.2 Auditoria Pública (Normativo)

Ao fechar o epoch, a Plataforma deve publicar relatório auditável contendo:

- `epoch_id`, `total_events`, `burn_total`
- `pool_storers`, `pool_verifiers`, `pool_platform`
- lista `(id, units, payout)` para storers e verifiers
- hash `SHA-256` do relatório

O hash do relatório **DEVE** ser registrado como evento no PoE (`payload_hash` do relatório).

---

## 11. Proibição de Política Monetária Ajustável (v0.1)

Em v0.1:

- supply fixo
- burn rate fixo
- splits fixos
- `BASE_FEE` e parâmetros fixos

Mudanças só via versionamento (`v0.2+`), com migração explícita.

---

## 12. Riscos e Avisos (Sem Promessa)

O token:

- não garante valorização, retorno, liquidez
- não é governança
- não é dividendos

O cliente final paga em fiat fora do protocolo; operadores compram `POE` no mercado para usar o FIFO.

---

## 13. Encerramento

Este modelo econômico v0.1 é desenhado para:

- tornar spam caro;
- impedir DoS por payload gigante (fees + limites);
- reduzir supply com uso (burn);
- remunerar trabalho real (storers e verifiers);
- proteger infraestrutura em picos (congestion fee 100% plataforma);
- manter previsibilidade (parâmetros fixos em v0.1).

**Quem usa, paga.**  
**Quem trabalha, recebe.**  
**Congestionamento protege a infraestrutura — não compra prioridade.**


