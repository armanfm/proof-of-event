# Security Policy — Proof of Event (PoE)

## 1. Objetivo deste Documento
Este documento descreve o modelo de segurança do Proof of Event (PoE), incluindo princípios de design, ameaças consideradas, limitações explícitas e diretrizes para divulgação responsável de vulnerabilidades.

O PoE é um protocolo determinístico de **certificação temporal de eventos**.  
Ele **não cria confiança**, **não valida conteúdo** e **não substitui verificação externa**.

A confiança em um evento **existe antes** de sua submissão ao PoE.

---

## 2. Princípio Fundamental de Confiança

No Proof of Event:

> **A verificação ocorre antes.**  
> **O registro ocorre depois.**

O protocolo assume que:
- o evento já foi verificado por um agente externo (instituição, sistema, autoridade, oráculo);
- a responsabilidade pela verificação **não pertence ao PoE**.

O PoE registra apenas um fato criptográfico:

> “Este evento, representado por este hash, foi certificado por esta instância, neste instante canônico.”

---

## 3. Superfície de Segurança Deliberadamente Reduzida

O PoE reduz sua superfície de ataque **por design**:

- não utiliza consenso distribuído;
- não utiliza mineração ou staking;
- não executa verificação semântica;
- não depende de inteligência artificial;
- não possui governança on-chain;
- não depende de redes P2P obrigatórias.

A segurança do sistema deriva de:
- simplicidade;
- determinismo;
- auditabilidade;
- reexecução independente.

---

## 4. O Que o PoE Protege

O PoE protege:

- a **integridade temporal** do registro;
- a **imutabilidade append-only** do ledger;
- o **encadeamento criptográfico** entre eventos;
- a **reexecução determinística** das provas;
- a **verificabilidade independente** por terceiros.

Uma vez registrado, um evento não pode ser alterado sem que a divergência seja **criptograficamente detectável**.

---

## 5. O Que o PoE NÃO Protege

O PoE **não protege**:

- a veracidade semântica de eventos;
- a legitimidade jurídica de documentos;
- a identidade ou intenção de participantes;
- a correção da verificação externa;
- dados armazenados fora do ledger do certificador;
- a disponibilidade de serviços externos (IPFS, Pinata, Discord, blockchains).

Esses aspectos estão **explicitamente fora do escopo** do protocolo.

---

## 6. Criptografia e Resiliência

### 6.1 Funções Hash Utilizadas

O PoE utiliza hashes criptográficos **sem chaves privadas**:

- **SHA-512** para `payload_hash_512` (conteúdo do evento externo);
- **SHA-256** para o encadeamento determinístico do ledger (`poe_hash`).

Propriedades:
- funções one-way;
- sem dependência de segredos;
- verificáveis por qualquer terceiro.

Eventos com `payload_hash_512` idêntico são tratados como representando o mesmo conteúdo externo.

### 6.2 Assinaturas Criptográficas

Assinaturas **não fazem parte do núcleo** do protocolo.

Quando utilizadas:
- são opcionais;
- são externas à prova PoE;
- podem usar algoritmos clássicos ou pós-quânticos (ex.: Dilithium);
- **não alteram** o `poe_hash`.

A quebra de um esquema de assinatura **não invalida** o histórico do ledger PoE.

---

## 7. Modelo de Ameaças

### 7.1 Spam e Abuso
**Ameaça:** submissão massiva de eventos.

**Mitigações operacionais:**
- custo econômico por evento (créditos / fuel);
- consumo unitário de crédito por aceitação;
- rate limit intencional (ex.: 1 submissão por segundo por IP).

---

### 7.2 Ataques Sybil
**Ameaça:** criação de múltiplas identidades.

**Mitigação:**
- identidades adicionais não reduzem custo;
- identidades não alteram a prova criptográfica;
- o PoE não atribui privilégios por identidade.

---

### 7.3 Certificador Malicioso
**Ameaça:** certificador negar eventos, atrasar registros ou agir de má-fé.

**Mitigação:**
- auditoria externa;
- reexecução independente das provas;
- responsabilidade institucional fora do protocolo.

> O PoE não impede má-fé institucional —  
> ele a torna **detectável**.

---

### 7.4 Perda de Dados Externos (IPFS / Pinata / Outros)
**Ameaça:** indisponibilidade de storage externo.

**Mitigação:**
- o ledger local do certificador é a fonte de verdade;
- a prova PoE depende apenas de hashes e timestamps;
- dados externos podem ser republicados sem alterar a prova.

A indisponibilidade de storage externo **não invalida** provas PoE.

---

## 8. Replay e Reuso de Eventos

O PoE é resistente a replay **por design criptográfico e operacional**:

- cada prova associa `payload_hash_512` a um `poe_timestamp_us`;
- o encadeamento inclui `previous_hash`, impedindo reutilização fora da cadeia;
- implementações podem adotar anti-replay por `event_id` com TTL.

Não são necessários:
- nonces do cliente;
- timestamps confiáveis do cliente;
- listas globais de replay.

---

## 9. Ferramentas Off-chain

Ferramentas off-chain podem:
- armazenar dados completos;
- executar verificação criptográfica;
- aplicar análise semântica;
- utilizar IA ou sistemas especialistas.

Falhas nessas ferramentas **não comprometem**:
- a validade das provas PoE;
- o ledger do certificador;
- a reexecução determinística.

---

## 10. Divulgação Responsável de Vulnerabilidades

### 10.1 Processo
- **Reporte:** contato privado com o mantenedor
- **Confirmação:** até 72h
- **Correção:**
  - críticas: até 30 dias
  - médias: até 90 dias
- **Divulgação:** após correção disponível

### 10.2 Escopo
Inclui:
- geração incorreta de provas;
- quebra de determinismo;
- corrupção do ledger;
- falhas na serialização canônica;
- inconsistências no encadeamento hash.

Exclui:
- bugs em implementações de terceiros;
- falhas em ferramentas off-chain;
- ataques econômicos de mercado;
- erros de verificação externa.

---

## 11. Defesa em Profundidade

O PoE aplica múltiplas camadas independentes:

### Camada Criptográfica
- SHA-512 (conteúdo);
- SHA-256 (encadeamento);
- hashes determinísticos;
- independência de assinaturas.

### Camada Operacional
- ledger soberano do certificador;
- append-only;
- reexecução independente;
- transparência total.

### Camada Institucional
- responsabilidade explícita;
- separação clara de papéis;
- auditoria externa.

A falha de uma camada **não compromete as demais**.

---

## 12. Limitações Explícitas

O Proof of Event **não promete**:
- inviolabilidade absoluta;
- proteção contra má-fé institucional;
- garantias econômicas;
- verdade semântica.

A segurança do sistema decorre da **clareza de seus limites**, não de promessas irreais.

---

## 13. Encerramento

O Proof of Event foi projetado para ser:
- simples;
- verificável;
- auditável;
- honesto quanto às suas responsabilidades.

O protocolo não cria confiança.  
Ele registra, de forma imutável, eventos que já foram confiados fora dele.

**A blockchain não decide.  
Ela testemunha.**



