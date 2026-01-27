# TOKENOMICS-IMPL-v0.1 — Proof of Event (PoE)
Implementação Econômica Normativa (com taxa de congestionamento)

**Versão:** 0.1  
**Status:** Normativo (regras executáveis)  
**Compatível com:** SPEC.md v0.1 + protocol/v0.1.md  
**Autor:** Armando José Freire de Melo  
**Licença:** Apache License 2.0  

---

## 0. Escopo

Este documento define **parâmetros numéricos** e **regras mecânicas** para:
- cobrança de taxas (fees) no FIFO;
- queima (burn) obrigatória;
- redistribuição operacional;
- taxa adicional de congestionamento destinada exclusivamente à Plataforma;
- regras de medição para pagamentos por trabalho.

Este documento **não promete lucro**, **não define preço de mercado** e **não opera exchange**.

---

## 1. Ativo Nativo

- **Nome:** Token PoE  
- **Símbolo:** `POE`  
- **Decimais:** `18`  
- **Tipo:** criptomoeda / ativo digital de liquidação operacional

---

## 2. Oferta (Supply)

### 2.1 Supply Total (Fixado no Genesis)
- **Supply total:** `1.000.000.000 POE` (um bilhão)
- Criado **uma única vez** no genesis.
- **Nenhum mint adicional** é permitido em v0.1.

### 2.2 Queima Reduz Supply
- A queima é **destruição irreversível** de tokens.
- O supply total **diminui** com o uso do protocolo.
- **Não existe “nascer mais token”** após queima.

---

## 3. Alocação Genesis (Distribuição Inicial)

A alocação genesis é definida em 4 buckets fixos (percentual do supply inicial):

1. **STORER_FOUNDERS:** `40%` = `400.000.000 POE`
2. **PLATFORM_RESERVE:** `30%` = `300.000.000 POE`
3. **VERIFIER_FOUNDERS:** `20%` = `200.000.000 POE`
4. **DEV_FUND:** `10%` = `100.000.000 POE`

### 3.1 Política de Liberação (Lock/Vesting) — v0.1
Para reduzir risco de “dump” e manter previsibilidade institucional:

- **STORER_FOUNDERS (400M):**
  - `25%` liberado no genesis
  - `75%` liberado linearmente por `12 meses` (liberação diária)

- **VERIFIER_FOUNDERS (200M):**
  - `25%` liberado no genesis
  - `75%` liberado linearmente por `12 meses` (liberação diária)

- **PLATFORM_RESERVE (300M):**
  - `40%` liberado no genesis (infra e bootstrap)
  - `60%` liberado linearmente por `36 meses` (liberação diária)

- **DEV_FUND (100M):**
  - `0%` liberado no genesis
  - `100%` liberado linearmente por `24 meses` (liberação diária)

> Regra: tokens travados (locked) **não contam** como disponíveis para pagamentos até liberados.

---

## 4. Identidades de Trabalho (Sem Governança)

O token PoE **não concede poder político**.  
Mas a redistribuição exige identificação de “quem trabalhou”.

Em v0.1, as identidades são definidas por IDs imutáveis:

- **verifier_id:** `bytes32` (hash de chave pública/identidade do verificador)
- **storer_id:** `bytes32` (hash de chave pública/identidade do armazenador)

Esses IDs aparecem:
- no evento (`source_id`) para o verificador/oráculo;
- na `Commitment_Proof` para armazenadores.

---

## 5. Estrutura de Taxas (Fee Model)

A taxa total para entrada no FIFO é composta por 4 parcelas:

1. **Base Fee (registro no FIFO)**
2. **Blob Fee (peso físico de anexos/assinaturas, por bytes)**
3. **Retention Fee (retenção do blob por tempo)**
4. **Congestion Fee (pico de demanda / fila cheia)**

### 5.1 Unidades
- Todas as taxas são cobradas em **POE** (unidade mínima: `10^-18 POE`).
- Arredondamento: **sempre para cima** (ceil) na unidade mínima.

---

## 6. Parâmetros Numéricos Fixos (v0.1)

### 6.1 Base Fee
- **BASE_FEE:** `0,10 POE` por evento aceito no FIFO.

### 6.2 Blob Fee (por tamanho)
Usado quando o evento referencia um blob (ex.: assinatura PQC completa).

- **BLOB_FEE_PER_KIB:** `0,002 POE` por `1 KiB` (1024 bytes)
- Cálculo:
  - `blob_kib = ceil(blob_bytes / 1024)`
  - `blob_fee = blob_kib * 0,002 POE`

### 6.3 Retention Fee (multiplicador por retenção)
A retenção define por quanto tempo o blob deve ser mantido disponível.

- `RETENTION_30D_MULT = 1x`
- `RETENTION_365D_MULT = 2x`
- `RETENTION_PERMANENT_MULT = 6x`

Cálculo:
- `retained_blob_fee = blob_fee * retention_multiplier`

> Se não existir blob, `blob_fee = 0` e `retained_blob_fee = 0`.

### 6.4 Hard Limits (mesmo pagando)
Para impedir DoS por payload gigante, v0.1 define limites duros:

- `MAX_BLOB_BYTES = 5.242.880` (5 MiB)
- `MAX_EVENT_JSON_BYTES = 8.192` (8 KiB)

Acima desses limites:
- **rejeitar** com `ERR_PAYLOAD_TOO_LARGE`.

### 6.5 Congestion Fee (Taxa de Congestionamento do FIFO)
A taxa de congestionamento é mecânica, baseada no tamanho da fila no instante de submissão.

Parâmetros:
- `Q0 = 100` (threshold)
- `QMAX = 10.000` (cap de medição)
- `CONGESTION_CAP_MULT = 20x` (limite de multiplicador)

Regra:
- Se `queue_size <= Q0` ⇒ `congestion_fee = 0`
- Se `queue_size > Q0` ⇒

q = min(queue_size, QMAX)
mult = ceil_div( (q - Q0), Q0 ) // 1x, 2x, 3x...
mult = min(mult, CONGESTION_CAP_MULT)
congestion_fee = BASE_FEE * mult


**Importantíssimo (normativo):**
- A congestion fee **NÃO altera ordem**.
- A congestion fee **NÃO compra prioridade**.
- A congestion fee **vai 100% para a Plataforma**.

---

## 7. Taxa Total a Pagar

Defina:
- `protocol_fee = BASE_FEE + retained_blob_fee`
- `total_fee = protocol_fee + congestion_fee`

O FIFO **DEVE** rejeitar se o pagador não possuir saldo para `total_fee`.

---

## 8. Liquidação Econômica (Burn + Redistribuição)

### 8.1 Separação Obrigatória (Normativa)
O FIFO separa a taxa em duas partes:

1. **Protocol Fee** (base + blob/retention)  
2. **Congestion Fee** (pico)

A **queima e redistribuição** são aplicadas **apenas** sobre `protocol_fee`.

A `congestion_fee` vai integralmente para a Plataforma.

### 8.2 Burn
- **BURN_RATE:** `10%` sobre `protocol_fee`
- `burn_amount = protocol_fee * 0,10`
- `distributable = protocol_fee - burn_amount`  (ou seja, `90%`)

A queima é:
- automática;
- obrigatória;
- irrevogável.

### 8.3 Redistribuição do `distributable` (Normativa)
Sobre `distributable`:

- **STORERS:** `40%`
- **VERIFIERS:** `30%`
- **PLATFORM:** `20%`
- **(sem sobra)** — totaliza 90% do protocol_fee após burn.

Cálculo:

- `to_storers   = distributable * 0,40`
- `to_verifiers = distributable * 0,30`
- `to_platform  = distributable * 0,20`

Além disso:
- `to_platform += congestion_fee`  (100% da congestion fee)

---

## 9. Medição de Trabalho (Como paga “quem fez”)

Pagamentos acontecem por **épocas** (epochs) para reduzir custo operacional.

### 9.1 Epoch
- **EPOCH_LENGTH:** `24h`
- Fuso: **UTC**
- A época é identificada por:
  - `epoch_id = YYYY-MM-DD (UTC)`.

### 9.2 Unidades de Trabalho — Verificadores
- Cada evento aceito no FIFO conta **1 unidade** para o `verifier_id` do evento.
- `verifier_units(verifier_id, epoch) = número de eventos aceitos no epoch`.

### 9.3 Unidades de Trabalho — Armazenadores
Um armazenador ganha unidade quando:

1. recebeu o evento do FIFO (sequência canônica),
2. realizou append no ledger local,
3. emitiu `Commitment_Proof` válida (opcional no SPEC, obrigatória para pagamento).

- Cada `Commitment_Proof` para um evento conta **1 unidade**.
- `storer_units(storer_id, epoch) = número de proofs válidas no epoch`.

### 9.4 Condição de Elegibilidade — Retenção Completa
Para receber pagamentos de armazenador, o nó **DEVE** manter:
- ledger completo desde GENESIS até head (segmentado por dia, mas contínuo);
- e estar sincronizado no hash canônico.

Nós fora de sincronia:
- recebem `0` para o período fora.

---

## 10. Distribuição dos Buckets (pagamento por epoch)

Para cada epoch:

1. calcula-se `pool_storers(epoch)` = soma de `to_storers` dos eventos do epoch
2. calcula-se `pool_verifiers(epoch)` = soma de `to_verifiers` dos eventos do epoch
3. calcula-se `pool_platform(epoch)` = soma de:
   - `to_platform` dos eventos do epoch
   - + todas as `congestion_fee` do epoch

### 10.1 Payout Pro-rata
- Pagamento por pro-rata de unidades no epoch:

Para armazenadores:
- `payout(storer_id) = pool_storers(epoch) * storer_units(storer_id) / total_storer_units(epoch)`

Para verificadores:
- `payout(verifier_id) = pool_verifiers(epoch) * verifier_units(verifier_id) / total_verifier_units(epoch)`

Plataforma:
- recebe `pool_platform(epoch)` integral.

### 10.2 Prova e Auditoria Pública
Ao fechar um epoch, a Plataforma **DEVE** publicar um relatório auditável:

- `epoch_id`
- `total_events`
- `burn_total`
- `pool_storers / pool_verifiers / pool_platform`
- lista `(id, units, payout)` para storers e verifiers
- hash SHA-256 do arquivo de relatório

Esse hash **DEVE** ser registrado como evento no PoE (payload_hash do relatório).

---

## 11. Regras de Acúmulo e “Token ficar escasso”

- O supply é fixo e reduz com queima.
- Se houver alta demanda, tokens podem ficar caros/escassos no mercado.
- Isso **não é bug**: é consequência de uso intenso.
- A Plataforma possui `PLATFORM_RESERVE` para sustentar operação e onboarding.

O protocolo **não** promete “sempre barato”.  
Ele promete **mecânica previsível**.

---

## 12. Combinação de Papéis (permitido)

Uma mesma entidade pode ser:
- verificador + armazenador,
- armazenador + plataforma,
- verificador + plataforma,
- tudo junto.

Isso é permitido.

Mesmo assim:
- o burn de 10% sempre ocorre;
- a ordem FIFO não muda;
- a congestion fee sempre vai para a plataforma.

---

## 13. Proibição de Política Monetária Ajustável (v0.1)

Em v0.1:
- supply é fixo;
- burn rate é fixo;
- splits são fixos;
- BASE_FEE e tarifas são fixas.

Mudanças só podem ocorrer via:
- **novo versionamento do protocolo** (v0.2, v0.3, …),
- com documentação pública e migração explícita.

---

## 14. Exemplo Numérico (concreto)

Cenário:
- `BASE_FEE = 0,10 POE`
- blob: `200 KiB` → `blob_fee = 200 * 0,002 = 0,40 POE`
- retention: `365D` → `retained_blob_fee = 0,40 * 2 = 0,80 POE`
- fila: `queue_size = 350`  
  - Q0=100 ⇒ mult = ceil((350-100)/100)=ceil(2,5)=3  
  - congestion_fee = 0,10 * 3 = 0,30 POE

Taxas:
- `protocol_fee = 0,10 + 0,80 = 0,90 POE`
- `total_fee = 0,90 + 0,30 = 1,20 POE`

Liquidação:
- burn = 10% de 0,90 = `0,09 POE`
- distributable = `0,81 POE`

Redistribuição:
- storers = 40% de 0,81 = `0,324 POE`
- verifiers = 30% de 0,81 = `0,243 POE`
- platform = 20% de 0,81 = `0,162 POE`
- + congestion_fee = `0,30 POE`
- platform total = `0,462 POE`

---

## 15. Encerramento

Esta implementação econômica v0.1 é desenhada para:

- tornar spam economicamente caro;
- impedir DoS por payload gigante (fee por bytes + limites duros);
- reduzir supply com uso (burn);
- remunerar trabalho real (storers e verifiers);
- proteger infraestrutura em picos (congestion fee 100% plataforma);
- manter previsibilidade institucional (parâmetros fixos em v0.1).

**Quem usa, paga.  
Quem trabalha, recebe.  
Congestionamento não beneficia terceiros — protege a infraestrutura.**
