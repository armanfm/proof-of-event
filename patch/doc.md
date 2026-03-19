# PoE L2 — Deterministic Quorum Consensus

> **Sequenciador L2 determinístico para liquidação institucional**
> Pré-valida off-chain, ancora on-chain, liquida tokens com prova imutável.

---

## 📌 Visão Geral

O PoE L2 é um sistema de liquidação baseado em execução off-chain determinística, onde múltiplos nós independentes processam os mesmos eventos e produzem um estado idêntico.

A validação **não ocorre na blockchain**.
A blockchain é utilizada apenas para:

- Liquidação final (movimentação de tokens)
- Ancoragem criptográfica (Merkle root)
- Auditoria imutável

---

## 🧠 Conceito Central

O sistema utiliza:

> **Consenso Determinístico por Quorum**

Isso significa:

- Todos os nós executam a mesma lógica
- O resultado deve ser **idêntico byte a byte**
- A validação ocorre comparando hashes (SHA256)

---

## ⚙️ Arquitetura

### 1. Nós L2 (Go)

Responsáveis por:

- Ordenação FIFO dos eventos
- Pré-validação (saldo, formato, regras)
- Execução das transações
- Geração de ledger determinístico
- Construção da Merkle Tree
- Envio para liquidação on-chain

Cada nó produz exatamente o mesmo resultado se estiver correto.

---

### 2. Verificador de Consenso (Python)

Serviço externo independente que:

- Coleta o ledger de múltiplos nós
- Calcula SHA256 de cada entrada
- Agrupa por igualdade de hash
- Define a verdade pelo quorum

Regra:

```
SHA256 igual     → válido
SHA256 diferente → inválido → slashing
```

> **Por que só SHA256?**
> SHA256 é determinístico e binário — ou o conteúdo é idêntico ou não é.
> Não existe "95% igual". Não há edge case com float. O(n) puro.

---

### 3. Smart Contract (Solidity)

Responsável por:

- Executar transferências de tokens (DREX)
- Registrar Merkle root (âncora imutável)
- Emitir eventos para auditoria

O contrato **não valida lógica de negócio** — recebe instruções pré-validadas e executa.

---

## 🔄 Fluxo do Sistema

```
Cliente envia transação
        ↓
[Nós L2 — Go]
  Pré-validação (saldo, endereço, valor)
  Ordenação FIFO soberana
  Ledger append-only
  Prova individual (SHA256)
        ↓
[Verificador — Python]
  Compara SHA256 de todos os nós
  Agrupa por quorum
  Divergência → slashing
        ↓
[Reversão entre aprovados]
  Nós do quorum confirmam entre si
        ↓
[Smart Contract — Solidity]
  Ancora Merkle root (prova imutável)
  Executa transferências de tokens
  Emite eventos auditáveis
```

---

## 🔐 Regra de Consenso

Para cada evento `seq`:

1. Cada nó produz um `entry`
2. O verificador calcula:

```python
hash = SHA256(json.dumps(entry, sort_keys=True))
```

3. Agrupa nós por hash idêntico

4. Determina:
   - Maior grupo = verdade
   - Minoria = divergente → slashing

---

## 📊 Quorum

```
quorum_minimo = floor(N / 2) + 1
falhos_tolerados = floor((N - 1) / 3)
```

| Nós | Quorum mínimo | Falhos tolerados |
|-----|---------------|-----------------|
| 3   | 2             | 1               |
| 5   | 3             | 1               |
| 7   | 4             | 2               |
| 9   | 5             | 2               |

> **Recomendação prática:**
> 3 nós = mínimo viável | 5 nós = bom | 7 nós = forte

---

## ⚖️ Propriedades

### ✔️ Determinismo

Mesma entrada → mesma lógica → mesmo hash

O determinismo é **garantido pelo SHA256**, independente do número de nós.
Mais nós não melhora o determinismo — ele já é 100%.
Mais nós melhora: tolerância a falhas e resistência a ataques.

---

### ✔️ Auditabilidade

- Ledger append-only (`ledger/full.txt`)
- Provas individuais por transação (`ledger/provas.txt`)
- Merkle root ancorado na blockchain para sempre

---

### ✔️ Simplicidade

- Sem consenso complexo (PBFT, PoS, PoW)
- Sem ZK obrigatório na versão atual
- Sem comunicação direta entre nós — verificador é externo
- SHA256 puro como árbitro

---

### ✔️ Escalabilidade

- Execução off-chain (Go — 5.000+ TPS)
- Batch + Merkle root
- Liquidação agregada: N transações em 1 tx on-chain

---

## 🚨 Divergência e Slashing

Se um nó produzir resultado diferente:

1. Verificador detecta via SHA256
2. Nó excluído do quorum
3. Slashing registrado em `slashing.log`
4. Notificação enviada ao smart contract

```python
# Lógica central do verificador
grupos = {}
for no_id, entry in conteudos.items():
    h = SHA256(entry)
    grupos.setdefault(h, []).append(no_id)

grupo_maior  = max(grupos.values(), key=len)
aprovado     = len(grupo_maior) >= QUORUM_MIN
```

---

## 📁 Estrutura de Provas

Cada transação gera três hashes independentes:

| Hash | Cobre |
|------|-------|
| `sha256_evento` | Evento completo |
| `sha256_payload` | Payload original |
| `sha256_conteudo` | `de\|para\|valor\|moeda` |

Arquivo: `ledger/provas.txt`

Formato de cada linha:
```
seq|sha256_evento|sha256_payload|sha256_conteudo|
source_id|de|para|valor|moeda|
liquidado|tx_hash|liquidado_utc|registrado_utc|event_hash
```

---

## 🔗 Liquidação

A liquidação ocorre apenas após:

1. Quorum atingido pelo verificador
2. Reversão confirmada entre nós válidos

Execução atômica no smart contract:

```
Chamada 1 → PoEAnchor.receberELiquidar()
  → ancora Merkle root (prova imutável)
  → executa transferFrom(de, para, valor) para cada instrução
```

Ou ancora + liquida juntos, ou nenhum acontece.

---

## 🧠 Diferença para Blockchain Tradicional

| Aspecto         | PoE L2                  | Blockchain tradicional |
|-----------------|-------------------------|------------------------|
| Execução        | Off-chain (Go)          | On-chain (EVM)         |
| Consenso        | Determinístico + quorum | BFT / PoS / PoW        |
| Validação       | Externa (Python)        | Interna                |
| Throughput      | 5.000+ TPS              | 100–500 TPS            |
| Custo           | Mínimo (1 tx por batch) | Alto (1 tx por evento) |
| Auditabilidade  | Ledger + blockchain     | Blockchain             |

---

## 🚀 Como Rodar

### Dependências

```bash
# Go >= 1.21
go version

# Python >= 3.9
python3 --version

# Node.js >= 18 (para o contrato)
node --version
```

### Gateway L2

```bash
go run main.go
# Rodando em http://localhost:8080
```

### Verificador de Consenso

```bash
# Demo local (sem nós)
python3 verificador_v2.py --demo

# Produção (com nós rodando)
python3 verificador_v2.py
```

### Smart Contract

```bash
cd poe-anchor
npm install
npx hardhat test
npx hardhat run scripts/deploy.js --network localhost
```

### Exemplo de transação

```bash
curl -X POST http://localhost:8080/submit \
  -H "Content-Type: application/json" \
  -d '{
    "version": "0.1",
    "event_id": "tx_001",
    "previous_event_hash": "GENESIS",
    "payload_hash": "0xabc123def456abc123def456abc123def456abc123def456abc123def456abc1",
    "source_id": "banco_inter",
    "local_timestamp": "2026-03-19T12:00:00Z",
    "de":    "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
    "para":  "0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB",
    "valor": "10000",
    "moeda": "DREX"
  }'
```

### Consultar prova de uma transação

```bash
curl http://localhost:8080/prova?seq=1
```

---

## 📂 Estrutura do Repositório

```
.
├── main.go                  # Gateway L2 — Go
├── fiscal.go                # Fiscal externo de integridade
├── verificador_v2.py        # Verificador de consenso — Python
├── poe-anchor/
│   ├── contracts/
│   │   └── PoEAnchor.sol    # Smart contract — Solidity
│   ├── scripts/
│   │   └── deploy.js
│   └── test/
│       └── PoEAnchor.test.js
└── ledger/
    ├── full.txt             # Ledger principal (append-only)
    ├── provas.txt           # Provas individuais
    ├── anchors.txt          # Registro de ancoragens
    └── YYYY-MM-DD.txt       # Ledger diário
```

---

## 🎯 Objetivo

Fornecer infraestrutura de liquidação:

- **Determinística** — mesmo input, mesmo output, em qualquer nó
- **Auditável** — qualquer auditor verifica sem depender de ninguém
- **Simples** — SHA256 como árbitro, sem complexidade desnecessária
- **Institucional** — adequado para banco central, fintechs, DREX

---

## 📌 Status

- ✅ Gateway L2 determinístico (Go)
- ✅ Ledger append-only com provas individuais
- ✅ Verificador de consenso por quorum (Python)
- ✅ Smart contract com ancora + liquidação atômica (Solidity)
- ✅ Fiscal externo de integridade de código
- ⏳ Integração completa com DREX sandbox
- ⏳ Slashing on-chain (lista dinâmica de validadores)
- ⏳ Mini STARK (gnark) substituindo stub

---

## 🔮 Roadmap

- Stake e slashing on-chain
- Lista dinâmica de validadores no contrato
- Integração com ZK Proof (mini STARK via gnark)
- Integração com Open Finance BR (NF-e como prova de comércio)
- Auditoria automática por terceiros

---

## 🔗 Relacionado

- [BCA Manuscript v31](https://github.com/armanfm/brics-elastic-monetary-model) — Arquitetura Monetária Elástica para o BRICS
- [PoE 15.0](https://github.com/armanfm/proof-of-event) — Sequenciamento Determinístico Auditável

---

## ⚠️ Observações de Segurança

- Segurança cresce com o número de validadores independentes
- Nós **devem** ser determinísticos — qualquer diferença de implementação quebra o consenso
- Chave privada do owner nunca deve ser hardcoded
- `slashing.log` deve ser monitorado em produção

---

## 👤 Autor

**Armando Freire**
[armanfm@github.com](mailto:armanfm@github.com)
Março 2026

---

## 📄 Licença

MIT
