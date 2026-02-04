# Proof of Event (PoE)
**A blockchain não decide. Ela testemunha.**

**Proof of Event (PoE)** é um protocolo determinístico para **certificar eventos externos** como fatos criptográficos **ancorados no tempo**, **sem consenso**, **sem votação** e **sem interpretação semântica on-chain**.

- PoE **não cria verdade**.  
- PoE **não interpreta significado**.  
- PoE **testemunha** eventos cuja validade já existe **fora** do sistema.

---

## 🎯 Objetivo do Projeto
O PoE foi projetado para ambientes onde:

- o consenso sobre o evento **já existe fora** do sistema;
- auditoria, rastreabilidade e reexecução são mais importantes que governança;
- mecanismos como PoW/PoS/staking/votação são indesejáveis;
- simplicidade, determinismo e compatibilidade institucional são requisitos.

### Exemplos de uso
- registros institucionais
- eventos legais/contratuais
- logs auditáveis
- sensores e sistemas industriais
- prova de execução e ocorrência

---

## 🧱 Arquitetura (Visão Geral)
O Proof of Event é dividido em **camadas estritamente desacopladas**:

### Camada 1 — Evento Externo (Fora do Escopo)
Onde o evento ocorre e é validado.

- validação
- auditoria
- responsabilidade
- verificação

Tudo isso acontece **antes** do PoE.

### Camada 2 — Certificação Temporal Determinística (PoE)
Executada por **um Certificador PoE** (instância independente).

O Certificador:
- recebe um **hash do evento** (`payload_hash_512`) e metadados mínimos;
- atribui um **timestamp canônico** (microsegundos UTC em 16 dígitos);
- gera uma prova determinística e **encadeia** em uma **hash chain**;
- registra em um **ledger append-only** (linhas imutáveis por append);
- emite um **recibo verificável** (JSON) para o cliente.

⚠️ **Não existe rede PoE.**  
Cada Certificador opera de forma **independente**, como prestador de serviço de certificação temporal.

#### Por que múltiplos certificadores?
O PoE não opera como rede única nem como sistema de consenso. Certificadores:
- não precisam concordar entre si;
- respondem apenas pelos eventos que certificam;
- podem ser escolhidos por fatores externos (jurisdição, contrato, reputação, exigência regulatória).

O PoE registra **quem certificou** e **quando** — não que múltiplas entidades concordaram sobre o evento.

### Camada 3 — Camadas Semânticas (Opcional)
Camadas externas podem:
- interpretar eventos;
- integrar sistemas;
- aplicar regras de negócio;
- enriquecer metadados.

Essas camadas **nunca interferem** na prova PoE.

---

## 🔐 Modos de Certificação: SELF e VERIFIED
O PoE suporta **dois modos** (mesmo motor, responsabilidades diferentes):

### 1) PoE SELF (Cliente)
O evento é certificado **somente** pelo cliente.

**Ledger (SELF):**

client_address | payload_hash_512 | poe_timestamp_us



### 2) PoE VERIFIED (Cliente + Verificador)
O evento é certificado pelo cliente e inclui um **verificador externo** explicitamente registrado na prova.

**Ledger (VERIFIED):**


client_address | payload_hash_512 | verifier_address | poe_timestamp_us


> O PoE não “julga” o verificador. Ele apenas registra que **uma entidade específica** foi incluída como verificador naquele evento, naquele instante.

---

## 🧾 O que é a Prova PoE (no backend atual)
A prova emitida pelo Certificador é derivada de um **hash determinístico encadeado** (hash chain) que inclui:

- `event_id`
- `client_address`
- `verifier_address` (opcional)
- `payload_hash_512`
- `timestamp_canônico` (microsegundos UTC em 16 dígitos)
- `previous_hash` (hash anterior do ledger)

Isso forma uma cadeia imutável por encadeamento, com raiz **GENESIS** e sequência monotônica (`sequence`).

> A prova PoE demonstra:  
> **“Este evento (representado por um hash) foi certificado por esta instância, até este momento no tempo, dentro de uma cadeia criptográfica ordenada.”**

---

## 🛡️ Regras operacionais de integridade
O backend implementa proteções determinísticas e anti-abuso:

- **anti-replay por `event_id`** com TTL (janela de 24h);
- **rate limit intencional**: no máximo **1 submissão por segundo por IP** (fricção operacional);
- validação estrita de campos (tamanhos, caracteres, formatos);
- estado persistente em `state.json`.

---

## 💰 Modelo Econômico (Fuel de Comunidade)
O backend opera com **créditos de uso** (“community fuel”):

- **1 evento aceito = 1 crédito consumido**
- sem créditos, o `/submit` retorna **NO_COMMUNITY_FUEL**
- créditos são **creditados por pagamentos confirmados**

### Pacotes (créditos)
- 50.000 créditos
- 500.000 créditos
- 5.000.000 créditos

### Pagamentos e idempotência
Os pagamentos são registrados com **idempotência por `tx_id`** e auditados em ledger administrativo.

- **Asaas**: webhook autenticado por `authToken` no header `asaas-access-token`
- **Mercado Pago**: valida assinatura do webhook e consulta o pagamento (fonte de verdade)

---

## 📒 Ledgers e Auditoria
O Certificador escreve logs append-only em:

- `ledger/daily/self/AAAA-MM-DD.log`
- `ledger/daily/verified/AAAA-MM-DD.log`
- `ledger/monthly/self/AAAA-MM.log`
- `ledger/monthly/verified/AAAA-MM.log`

E também um ledger administrativo (pagamentos):

- `ledger/admin/payments-AAAA-MM-DD.log`

---

## 🔌 Integrações opcionais
- **Pinata**: upload opcional do recibo (proof JSON) e retorno de `cid`
- **Discord**: webhook opcional por evento + envio diário automático dos ledgers

---

## 🌐 Endpoints (Implementação de Referência)
- `POST /submit` — registra evento (consome 1 fuel)
- `GET /status` — status do sistema (sequência, fuel, paths, preços)
- `GET /health` — healthcheck
- `GET /debug` — protegido por token (amostra de payments + estado)
- `POST /webhook/asaas` — crédito de fuel (Asaas)
- `POST /webhook/payment` — alias do Asaas (compatibilidade)
- `POST /webhook/mercadopago` — crédito de fuel (Mercado Pago)
- `/` — UI estática servida de `public/`

---

## 📜 Especificação Técnica
A definição formal e normativa do protocolo está em:

➡️ `SPEC.md`

O SPEC é a fonte única de verdade técnica.

---

## 🔬 Status do Projeto
- 🧠 Fundação conceitual: consolidada  
- 📐 Especificação técnica: definida (v0.1)  
- ⚙️ Implementação de referência: **operacional (beta)**  
- 💰 Modelo econômico: implementado via **fuel por pagamento**  

---

### Limite Operacional de Submissão (Anti-abuso)

Para preservar a estabilidade do serviço e evitar abuso, um Certificador PoE
PODE aplicar limites operacionais à submissão de eventos.

Na implementação de referência, aplica-se:

- limite mínimo de 1 segundo entre submissões por origem.

Este limite:

- NÃO faz parte da prova PoE;
- NÃO altera o payload_hash;
- NÃO altera o timestamp canônico emitido;
- NÃO confere prioridade a nenhum participante;
- NÃO é comprável ou negociável.

Trata-se exclusivamente de uma política operacional anti-abuso.

## ⚖️ Licença
Apache License 2.0

**Autor da especificação conceitual:**  
Armando Freire

> PoE existe para registrar eventos como fatos criptográficos,  
> não como decisões sociais.



