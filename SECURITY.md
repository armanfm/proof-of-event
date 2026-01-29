# Security Policy — Proof of Event (PoE)

## 1. Objetivo deste Documento

Este documento descreve o modelo de segurança do Proof of Event (PoE),
incluindo princípios de design, ameaças consideradas, limitações explícitas
e diretrizes para divulgação responsável de vulnerabilidades.

O PoE é um protocolo determinístico de registro temporal de eventos.
Ele **não cria confiança**, **não valida conteúdo** e **não substitui
verificação externa**.

A confiança em um evento **existe antes** de sua submissão ao PoE.

---

## 2. Princípio Fundamental de Confiança

No Proof of Event:

**A verificação ocorre antes.  
O registro ocorre depois.**

O protocolo assume que:

- o evento já foi verificado por um agente externo
  (instituição, oráculo, sistema, autoridade);
- a responsabilidade pela verificação **não pertence ao PoE**.

O PoE registra apenas um fato criptográfico:

> “Este evento, representado por este hash,
> foi registrado por este certificador,
> neste instante canônico.”

---

## 3. Superfície de Segurança Deliberadamente Reduzida

O PoE reduz sua superfície de ataque por design:

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
- a **imutabilidade** do histórico;
- a **associação criptográfica** entre evento e timestamp;
- a **reexecução determinística** das provas;
- a **verificabilidade independente** por terceiros.

Uma vez registrado, um evento **não pode ser alterado**
sem que a divergência seja detectável.

---

## 5. O Que o PoE NÃO Protege

O PoE **não protege**:

- a veracidade semântica de eventos;
- a legitimidade jurídica de documentos;
- a identidade ou intenção de participantes;
- a correção da verificação externa;
- dados armazenados fora do ledger do certificador;
- a disponibilidade de serviços externos (IPFS, Pinata, blockchains).

Esses aspectos estão **explicitamente fora do escopo** do protocolo.

---

## 6. Criptografia e Resiliência Pós-Quântica

### 6.1 Hashes Criptográficos

O núcleo do PoE utiliza exclusivamente:

- **SHA-512**
- hashes one-way
- sem dependência de chaves privadas

O protocolo assume a resistência a colisões do SHA-512.

Eventos com `payload_hash` idêntico são tratados como o mesmo evento.

---

### 6.2 Assinaturas Criptográficas

Assinaturas **NÃO fazem parte do núcleo do protocolo**.

Quando utilizadas:
- são opcionais;
- são externas à prova PoE;
- podem utilizar algoritmos clássicos ou pós-quânticos (ex: Dilithium).

A quebra de um esquema de assinatura **não invalida**
o histórico do ledger PoE.

---

## 7. Modelo de Ameaças (Resumo)

### 7.1 Spam e Abuso

**Ameaça:** Submissão massiva de eventos.  
**Mitigação:** Custo econômico por evento (Token PoE, quando aplicável).

---

### 7.2 Ataques Sybil

**Ameaça:** Criação de múltiplas identidades.  
**Mitigação:** Identidades adicionais não reduzem custo nem alteram provas.

---

### 7.3 Certificador Malicioso

**Ameaça:** Certificador tenta negar eventos ou emitir provas incorretas.  
**Mitigação:**  
- auditoria externa;
- reexecução independente;
- responsabilidade institucional fora do protocolo.

O PoE não impede má-fé institucional — ele a torna **detectável**.

---

### 7.4 Perda de Dados Externos (IPFS / Pinata)

**Ameaça:** Indisponibilidade de storage externo.  
**Mitigação:**  
- o ledger do certificador é a fonte de verdade;
- a prova PoE depende apenas de hashes e timestamps;
- dados externos podem ser republicados sem alterar a prova.

A queda do IPFS **não invalida** provas PoE.

---

## 8. Ataques de Replay

O PoE é **intrinsecamente resistente a replay**.

Cada prova:
- associa um `payload_hash` a um `timestamp_canônico`;
- produz um hash único (PoE_Proof);
- não pode ser reutilizada para um instante diferente.

Não são necessários:
- nonces;
- timestamps confiáveis do cliente;
- listas de replay.

---

## 9. Ferramentas Off-chain

Ferramentas off-chain podem:

- armazenar dados completos;
- executar verificação criptográfica;
- aplicar análise semântica;
- utilizar IA ou sistemas especialistas.

Falhas nessas ferramentas **não comprometem**:

- a validade das provas;
- o ledger do certificador;
- a reexecução determinística.

---

## 10. Divulgação Responsável de Vulnerabilidades

### 10.1 Processo

- **Reportar:** contato privado com o mantenedor
- **Confirmação:** até 72h
- **Correção:**
  - críticas: até 30 dias
  - médias: até 90 dias
- **Divulgação:** após correção disponível

---

### 10.2 Escopo

Inclui:
- geração incorreta de provas;
- quebra de determinismo;
- corrupção do ledger;
- falhas na serialização canônica.

Exclui:
- bugs em implementações específicas;
- falhas em ferramentas off-chain;
- ataques econômicos de mercado;
- erros de verificação externa.

---

## 11. Defesa em Profundidade

O PoE aplica múltiplas camadas independentes:

### Camada Criptográfica
- SHA-512;
- hashes determinísticos;
- independência de assinaturas.

### Camada Operacional
- ledger soberano do certificador;
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

A segurança do sistema decorre da clareza de seus limites,
não de promessas.

---

## 13. Encerramento

O Proof of Event foi projetado para ser:

- simples;
- verificável;
- auditável;
- honesto quanto às suas responsabilidades.

O protocolo não cria confiança.
Ele registra, de forma imutável, eventos
que já foram confiados fora dele.

**A blockchain não decide.  
Ela testemunha.**


