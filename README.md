# Proof of Event (PoE)

**A blockchain não decide. Ela testemunha.**

Proof of Event (PoE) é um protocolo de registro determinístico de eventos externos.
Seu objetivo é fornecer uma prova **imutável, auditável e reexecutável** de que
um evento ocorreu, **sem consenso, sem votação e sem julgamento on-chain**.

PoE não cria verdade.  
PoE não interpreta significado.  
PoE apenas **testemunha eventos cuja ocorrência já é aceita fora do sistema**.

---

## 🎯 Objetivo do Projeto

O PoE foi projetado para ambientes onde:

- o consenso sobre o evento **já existe fora da blockchain**
- auditoria, rastreabilidade e reexecução são mais importantes que governança
- mecanismos como PoW, PoS, staking ou votação são **indesejáveis**
- simplicidade, determinismo e compatibilidade institucional são requisitos

Exemplos de uso:

- eventos institucionais
- registros legais
- logs auditáveis
- sensores e sistemas industriais
- provas de execução e ocorrência

---

## 🧱 Arquitetura (Visão Geral)

O protocolo é dividido em camadas **estritamente desacopladas**:

### Camada 1 — Evento Externo
Onde o evento acontece (fora do escopo do PoE).  
A validação, verificação e responsabilidade ocorrem **antes** do registro no protocolo.

### Camada 2 — Ledger Determinístico (Núcleo PoE)
Ledger *append-only*, ordenado por **FIFO soberano**, encadeado por hash,
sem consenso, sem votação e sem forks.

### Camada 3 — Ledger Semântico (Opcional)
Camada de interpretação, contexto ou integração institucional.  
Pode enriquecer eventos, mas **nunca interfere** na validade, ordem ou integridade
do ledger PoE.

> ⚠️ Apenas a **Camada 2** faz parte do protocolo PoE.

---

## 💰 Modelo Econômico (Visão Geral)

O Proof of Event opera com uma **criptomoeda nativa de infraestrutura**,
utilizada para o pagamento do uso do protocolo e para a remuneração dos
participantes operacionais.

### Princípios Fundamentais

- O **cliente final** paga pelos serviços em moeda fiduciária.
- **Verificadores, tokenizadores e operadores** pagam o uso do protocolo
  (ex: acesso à fila FIFO) utilizando a criptomoeda.
- **Armazenadores, verificadores e a plataforma** recebem criptomoeda como
  remuneração por trabalho efetivamente executado.
- A criptomoeda possui **oferta fixa**, criada uma única vez, e **não é inflacionária**.
- O token **circula**: não é criado pelo FIFO, apenas redistribuído.
- O preço da criptomoeda é definido **exclusivamente pelo mercado**.
- O protocolo PoE **não garante retorno financeiro**, valorização ou rendimento
  associado à posse do token.

A criptomoeda funciona como **mecanismo de liquidação de custos de infraestrutura**
e **remuneração operacional**, não como instrumento de governança ou promessa
financeira.

---

## ❌ O que o PoE NÃO é

- Não é apenas uma criptomoeda especulativa
- Não é um protocolo de consenso
- Não é um sistema de governança
- Não é uma DAO
- Não é um árbitro de verdade ou significado
- Não recompensa usuários finais com tokens
- Não promete retorno financeiro ou valorização

---

## 📜 Especificação Técnica

A definição formal, completa e normativa do protocolo está em:

➡️ **`/SPEC.md`**

O SPEC é a **fonte de verdade técnica** do projeto.

---

## 🔬 Status do Projeto

- 🧠 Fundação conceitual: **consolidada**
- 📐 Especificação técnica: **em elaboração**
- ⚙️ Implementação de referência: **a definir**
- 💰 Modelo econômico: **definido em nível conceitual**

Este repositório começa pela **especificação**, não pela implementação.

---

## ⚖️ Licença

Apache License 2.0

Autor da especificação conceitual:  
**Armando José Freire de Melo**

---

> PoE existe para registrar eventos como fatos criptográficos,  
> não como decisões sociais.

