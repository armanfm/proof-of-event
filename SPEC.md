# Proof of Event (PoE)
## Especificação Técnica Oficial — Camada 2 (Ledger Determinístico)

**Versão:** 0.1 (alinhada com implementação beta)  
**Status:** Fundação Técnica  
**Autor:** Armando Freire  
**Licença:** Apache License 2.0  

---

## 1. Escopo do Protocolo
O Proof of Event (PoE) é um protocolo determinístico para certificação temporal de eventos externos, cujo objetivo é produzir um registro criptográfico **imutável, auditável e reexecutável** de existência temporal, operando **sem consenso**, **sem votação** e **sem interpretação semântica on-chain**.

O PoE **não valida significado**, **não decide verdade** e **não resolve disputas**.  
Ele registra eventos cuja ocorrência e validação foram tratadas fora do protocolo (Camada 1).

Esta especificação define exclusivamente a **Camada 2 — Ledger Determinístico PoE** e o papel do **Certificador PoE**.

---

## 2. Princípios Fundamentais

### 2.1 Determinismo Absoluto
Dada a mesma entrada válida, qualquer implementação compatível do PoE **DEVE** produzir exatamente o mesmo resultado criptográfico.

Não existe aleatoriedade, votação ou interpretação subjetiva.

### 2.2 Ausência de Consenso
O PoE não implementa consenso distribuído.

Não existem:
- votação
- mineração
- staking
- slashing
- forks
- governança on-chain

O protocolo não resolve conflitos sociais, jurídicos ou semânticos.

### 2.3 Ancoragem Temporal Canônica
O PoE atribui um **timestamp canônico** no momento da aceitação do evento, gerado pelo Certificador PoE.

A ordenação global entre eventos **não é objetivo do PoE**.  
A única ordenação relevante é **operacional e local ao Certificador**, derivada do ledger append-only e do encadeamento criptográfico.

### 2.4 Imutabilidade Append-Only
Eventos aceitos **NUNCA** são alterados, removidos ou reescritos.

O ledger é estritamente **append-only** dentro de cada Certificador PoE.

### 2.5 Neutralidade do Protocolo
O PoE é cego a:
- identidade social
- reputação
- valor econômico
- conteúdo semântico do evento

O protocolo aplica apenas regras mecânicas e determinísticas.

---

## 3. Definições

- **Evento Externo (Camada 1):** fato ocorrido fora do PoE.
- **Certificador PoE:** entidade que executa a Camada 2, recebe eventos validados off-chain e emite provas PoE.
- **Evento Canônico:** representação determinística reduzida do evento externo.
- **payload_hash_512:** hash SHA-512 do conteúdo do evento externo, em hexadecimal lowercase (128 caracteres).
- **event_id:** identificador textual do evento (anti-replay e referência).
- **client_address:** endereço do cliente (formato Ethereum `0x` + 40 hex, lowercase).
- **verifier_address (opcional):** endereço do verificador externo (mesmo formato do client_address).
- **poe_timestamp_us:** timestamp canônico em **microsegundos UTC**, serializado como **string decimal de 16 dígitos**.
- **sequence:** contador monotônico local ao Certificador (inteiro crescente).
- **previous_hash:** hash anterior da cadeia (ou sentinela GENESIS).
- **poe_hash:** hash determinístico da aceitação do evento (encadeado).

> Observação: esta versão do PoE não valida identidade social; endereços apenas identificam participantes no registro.

---

## 4. Arquitetura do Sistema

### 4.1 Camada 1 — Evento Externo (Fora do Escopo)
A validação, auditoria e responsabilidade legal do evento ocorrem fora do PoE.

O protocolo assume que:
- houve verificação off-chain; e
- o evento foi reduzido a um `payload_hash_512`.

### 4.2 Camada 2 — Certificador PoE (Este Protocolo)
A Camada 2 é executada por um Certificador PoE, responsável por:
- receber eventos canônicos;
- atribuir `poe_timestamp_us`;
- gerar `poe_hash` encadeado;
- registrar em ledger append-only;
- emitir recibos verificáveis.

O PoE **não define**:
- rede distribuída;
- replicação entre certificadores;
- consenso entre certificadores;
- topologia de infraestrutura.

Cada certificador opera de forma soberana e independente.

---

## 5. Modos de Certificação: SELF e VERIFIED

O PoE suporta dois modos compatíveis:

### 5.1 PoE SELF
Registro certificado apenas pelo cliente.

**Campos mínimos:**
- event_id
- client_address
- payload_hash_512
- poe_timestamp_us

### 5.2 PoE VERIFIED
Registro certificado pelo cliente com verificador externo explícito.

**Campos mínimos:**
- event_id
- client_address
- verifier_address
- payload_hash_512
- poe_timestamp_us

O PoE não “julga” o verificador. Ele apenas registra que um verificador específico foi incluído naquele evento.

---

## 6. Prova Criptográfica (poe_hash) — Definição Formal

### 6.1 Hash de Encadeamento
A unidade criptográfica emitida pelo Certificador é um hash encadeado:



poe_hash = SHA-256(
event_id |
client_address |
verifier_address |
payload_hash_512 |
poe_timestamp_us |
previous_hash
)



Regras:
- o separador lógico entre campos é o caractere `|` (pipe);
- `verifier_address` deve ser string vazia quando ausente (modo SELF);
- `previous_hash` é o hash anterior da cadeia do Certificador;
- o algoritmo de encadeamento nesta versão é **SHA-256**;
- o resultado é hexadecimal lowercase (64 caracteres).

### 6.2 GENESIS (Sentinela)
A cadeia inicia com:

- `previous_hash = "GENESIS"`
- `sequence = 0` antes do primeiro evento aceito

O primeiro evento aceito deve ter:
- `sequence = 1`
- `previous_hash = "GENESIS"`

---

## 7. Timestamp Canônico (poe_timestamp_us)

O timestamp canônico:
- deve ser gerado por fonte de tempo do Certificador;
- deve ser expresso em UTC;
- deve usar microsegundos desde Unix epoch;
- deve ser serializado como **string decimal de 16 dígitos**.

O Certificador pode incluir também um campo auxiliar `recorded_at` (RFC3339) apenas como metadado humano; ele não participa da prova criptográfica.

---

## 8. Formato do Evento Canônico

Todo evento submetido ao Certificador deve conter, no mínimo:
- event_id
- client_address
- payload_hash_512
- verifier_address (opcional)

### 8.1 Regras do payload_hash_512
O `payload_hash_512` deve:
- ser SHA-512 do conteúdo do evento externo;
- ser hexadecimal lowercase;
- possuir comprimento fixo de 128 caracteres.

Eventos com `payload_hash_512` idêntico são tratados como representando o mesmo conteúdo externo.

---

## 9. Regras de Aceitação (Normativo)

Um evento é aceito se, e somente se:
1. os campos mínimos do modo escolhido (SELF ou VERIFIED) forem válidos;
2. o Certificador atribuir `poe_timestamp_us`;
3. o Certificador computar `poe_hash` com o `previous_hash` atual;
4. o Certificador registrar o evento em ledger append-only;
5. requisitos operacionais do Certificador forem atendidos (ex.: disponibilidade de créditos, quando aplicável).

Não existe rejeição baseada em conteúdo semântico.

---

## 10. Reexecução e Verificação

Para verificar um evento aceito dentro de um Certificador:
1. obter os campos:
   - event_id
   - client_address
   - verifier_address (ou vazio)
   - payload_hash_512
   - poe_timestamp_us
   - previous_hash
2. recomputar:
   - `poe_hash' = SHA-256(event_id|client|verifier|payload|ts|previous)`
3. comparar `poe_hash'` com `poe_hash` registrado/emitido.

Se coincidir, a prova é válida e reexecutável.

---

## 11. Ledger Determinístico (Append-only)

O Certificador mantém ledger em texto, append-only, com duas categorias:

### 11.1 Ledger SELF
Linha por evento:



### 11.2 Ledger VERIFIED
Linha por evento:


client_address|payload_hash_512|verifier_address|poe_timestamp_us



O Certificador pode manter arquivos por dia e por mês (política operacional), desde que preserve o histórico.

---

## 12. Modelo Econômico Operacional (Créditos / Fuel)

Quando o Certificador adota cobrança:
- o pagamento deve ocorrer **antes** da aceitação do evento;
- a falta de créditos pode causar rejeição operacional;
- dados econômicos podem constar como metadados e ledger administrativo.

O modelo econômico **não participa** do `poe_hash` e **não altera** a prova criptográfica.

> Nesta implementação de referência, o Certificador usa um saldo de créditos (“community_fuel”) e consome 1 crédito por evento aceito.

---

## 13. Anti-replay e Rate Limit (Operacional Recomendado)

Para reduzir abuso e replay:
- recomenda-se anti-replay por `event_id` com janela temporal (TTL);
- recomenda-se limitar frequência por origem (ex.: 1 submissão por segundo por IP).

Essas regras são operacionais e não alteram o formato de prova.

---

## 14. Retenção do Ledger
Um Certificador PoE deve manter todo o histórico desde o início operacional (GENESIS sentinela).

Retenção parcial é proibida para certificadores que se declarem auditáveis.

---

## 15. Versionamento
Mudanças:
- devem incrementar versão;
- nunca devem invalidar provas antigas.

Algoritmos e formatos podem evoluir, mas devem manter compatibilidade de verificação por versão.

---

## 16. Encerramento
O Proof of Event existe para registrar eventos como fatos criptográficos,
não como decisões sociais.

**A blockchain não decide. Ela testemunha.**


