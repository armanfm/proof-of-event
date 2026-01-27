# Security Policy — Proof of Event (PoE)

## 1. Objetivo deste Documento

Este documento descreve o modelo de segurança do Proof of Event (PoE),
incluindo princípios de design, ameaças consideradas, limitações explícitas
e diretrizes para divulgação responsável de vulnerabilidades.

O PoE é um protocolo de **registro determinístico de eventos**.
Ele **não cria confiança**, **não valida conteúdo** e **não substitui verificação externa**.

A confiança em um evento **existe antes** de sua submissão ao PoE.

---

## 2. Princípio Fundamental de Confiança

No Proof of Event:

> **A verificação ocorre antes.  
> O registro ocorre depois.**

O protocolo assume que:
- o evento já foi verificado por um agente externo (verificador, oráculo, sistema);
- a responsabilidade pela verificação **não pertence ao PoE**.

O PoE apenas registra um **fato criptográfico**:
> “este evento, representado por este hash, foi registrado nesta ordem”.

---

## 3. Superfície de Segurança Deliberadamente Reduzida

O PoE reduz sua superfície de ataque por design:

- não utiliza consenso distribuído;
- não utiliza criptografia assimétrica no ledger;
- não executa verificação semântica;
- não depende de inteligência artificial;
- não possui governança on-chain.

A segurança do sistema deriva da **simplicidade, determinismo e auditabilidade**.

---

## 4. O Que o PoE Protege

O PoE protege:

- a **integridade temporal** do registro;
- a **imutabilidade do histórico**;
- a **ordem canônica de eventos (FIFO)**;
- a **reexecução determinística** do ledger;
- a **verificabilidade independente** por terceiros.

Uma vez registrado, um evento não pode ser alterado, removido ou reordenado
sem que a divergência seja detectável.

---

## 5. O Que o PoE NÃO Protege

O PoE **não protege**:

- a veracidade semântica de eventos;
- a legitimidade legal de documentos;
- a identidade ou intenção de participantes;
- a correção da verificação externa;
- dados armazenados fora do ledger;
- resultados produzidos por ferramentas off-chain.

Esses aspectos estão **explicitamente fora do escopo** do protocolo.

---

## 6. Criptografia e Resiliência Pós-Quântica

### 6.1 Justificativa Técnica

O ledger PoE armazena exclusivamente **hashes criptográficos SHA-256**.

Mesmo na presença de computadores quânticos capazes de quebrar
criptografia assimétrica clássica (RSA, ECDSA):

1. **Hashes são one-way**  
   Algoritmos quânticos conhecidos não permitem reverter SHA-256 de forma prática.

2. **A ordem é imutável por design**  
   A sequência FIFO não depende de chaves privadas ou assinaturas.

3. **Assinaturas são externas ao protocolo**  
   Esquemas criptográficos, inclusive pós-quânticos (PQC), podem ser utilizados
   na verificação off-chain sem impactar o ledger histórico.

Exemplo ilustrativo:

2015: Evento #100 → SHA256(dados) → 0xabc123...
2035: Computador quântico disponível
2035: Evento #100 → SHA256(dados) → 0xabc123... (inalterado)


O PoE é **resiliente por design** a avanços criptográficos, pois não fixa
algoritmos de identidade no núcleo do protocolo.

---

## 7. Modelo de Ameaças (Resumo)

### 7.1 Spam e Abuso
- **Ameaça:** Submissão massiva de eventos.
- **Mitigação:** Custo econômico obrigatório por evento (Token PoE).

### 7.2 Sybil
- **Ameaça:** Criação de múltiplas identidades.
- **Mitigação:** Identidades adicionais não reduzem custo nem alteram ordem.

### 7.3 Nós Maliciosos
- **Ameaça:** Armazenadores tentam alterar histórico local.
- **Mitigação:** Divergência de hash é detectável; nó deve resincronizar.

### 7.4 Falha de Armazenadores
- **Ameaça:** Nós offline ou instáveis.
- **Mitigação:** Perda de remuneração e sincronização obrigatória ao retorno.

### 7.5 Falha do FIFO
- **Ameaça:** Indisponibilidade temporária.
- **Mitigação:** O ledger permanece válido; apenas novas entradas são suspensas.

---

## 8. Ameaças Específicas

### 8.1 Reorganização de Histórico (History Replay)

**Ameaça:**  
Um operador apaga um ledger local e tenta recriar uma versão alternativa.

**Mitigações possíveis:**  
- múltiplos armazenadores independentes;
- checkpoints públicos periódicos;
- ancoragem opcional de hashes em blockchains públicas.

Essas mitigações são **operacionais** e podem evoluir sem alterar o protocolo.

---

## 9. Ferramentas Off-chain e Inteligência

Ferramentas off-chain (ex.: `mind.bin`) podem:

- armazenar dados completos;
- executar verificação criptográfica;
- aplicar análise semântica;
- utilizar algoritmos clássicos ou pós-quânticos.

Falhas nessas ferramentas **não comprometem**:
- a validade do ledger;
- a ordem FIFO;
- a imutabilidade do histórico.

O PoE não depende dessas ferramentas para sua segurança.

---

## 10. Divulgação Responsável de Vulnerabilidades

### 10.1 Processo

1. **Reportar:** contato privado com o mantenedor do projeto.
2. **Confirmar:** até 72h para confirmação de recebimento.
3. **Correção:**  
   - até 30 dias para vulnerabilidades críticas;  
   - até 90 dias para vulnerabilidades médias.
4. **Divulgação:** após correção disponível.

### 10.2 Escopo

**Inclui:**
- manipulação da ordem FIFO;
- bypass de pagamento obrigatório;
- corrupção do ledger canônico.

**Exclui:**
- bugs em implementações específicas;
- falhas em sistemas off-chain;
- ataques econômicos de mercado;
- erros de verificação externa.

---

## 11. Defesa em Profundidade

O PoE aplica múltiplas camadas independentes:

### Camada Criptográfica
- hashes SHA-256;
- encadeamento imutável;
- independência de criptografia assimétrica.

### Camada Econômica
- custo real por evento;
- desincentivo a spam;
- alinhamento por redistribuição.

### Camada Operacional
- replicação independente;
- sincronização determinística;
- ordenação centralizada e auditável.

### Camada Institucional
- transparência total;
- leitura pública do ledger;
- separação clara de responsabilidades.

A falha de uma camada **não compromete as demais**.

---

## 12. Limitações Explícitas

O Proof of Event **não promete**:

- inviolabilidade absoluta;
- resistência total a todos os ataques;
- proteção contra má-fé na verificação externa;
- garantias econômicas ou financeiras.

A segurança do sistema decorre da **clareza de seus limites**, não de promessas.

---

## 13. Encerramento

O Proof of Event foi projetado para ser:

- simples;
- verificável;
- auditável;
- honesto quanto às suas responsabilidades.

O protocolo não cria confiança.
Ele registra, de forma imutável, eventos que **já foram confiados** fora dele.

A blockchain não decide.  
Ela testemunha.
