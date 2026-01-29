
Onde:

- `payload_hash` é o hash da informação fornecida externamente
- `timestamp_canônico` é gerado pelo próprio PoE no momento da aceitação

Qualquer ordenação interna (fila, FIFO, etc.) é **estritamente operacional** e **não possui valor semântico probatório**.

---

### Camada 3 — Camadas Semânticas (Opcional)

Camadas externas e opcionais que podem:

- interpretar eventos
- enriquecer metadados
- integrar sistemas institucionais
- aplicar regras de negócio
- gerar visualizações ou relatórios

Essas camadas:

- **não interferem** na prova
- **não alteram** o ledger PoE
- **não participam** da validade criptográfica

---

## 🔐 Prova Criptográfica

O Proof of Event não tenta responder *o que* um evento significa.  
Ele responde apenas:

> **“Esta informação existia a partir deste momento.”**

A prova é:

- determinística
- reexecutável
- verificável independentemente
- resistente a interpretação subjetiva

Assinaturas digitais, identidades, certificados ou criptografia pós-quântica **não fazem parte do núcleo probatório**.  
Quando utilizadas, pertencem a **camadas auxiliares**, fora do hash canônico.

---

## 💰 Modelo Econômico (Visão Geral)

O Proof of Event pode operar com uma criptomoeda nativa de infraestrutura, utilizada exclusivamente para:

- pagamento pelo uso do protocolo
- liquidação de custos operacionais
- remuneração de participantes técnicos

### Princípios fundamentais

- o cliente final paga pelos serviços em moeda fiduciária
- operadores técnicos utilizam a criptomoeda para acessar o protocolo
- armazenadores, verificadores e operadores recebem criptomoeda por trabalho executado
- a criptomoeda **não é criada pelo núcleo PoE**
- o token apenas circula e é redistribuído
- não há promessa de retorno financeiro
- não há governança on-chain
- o protocolo não incentiva especulação

O modelo econômico é **operacional**, **desacoplado da prova criptográfica** e **não faz parte do núcleo conceitual do PoE**.

---

## ❌ O que o PoE NÃO é

- não é um protocolo de consenso
- não é uma blockchain tradicional
- não é uma DAO
- não é um sistema de governança
- não é um árbitro de verdade
- não é um sistema de votação
- não é um mecanismo de recompensa ao usuário final
- não promete retorno financeiro
- não cria significado social

---

## 📜 Especificação Técnica

A definição formal, normativa e técnica do protocolo está em:

➡️ **`/SPEC.md`**

O arquivo `SPEC.md` é a **fonte de verdade técnica** do projeto.

---

## 🔬 Status do Projeto

- 🧠 Fundação conceitual: **consolidada**
- 📐 Especificação técnica: **em elaboração**
- ⚙️ Implementação de referência: **a definir**
- 💰 Modelo econômico: **definido em nível conceitual**

Este repositório começa pela **especificação**, não pela implementação.

---

## ✍️ Autor

**Armando Freire**

---

> **Proof of Event existe para registrar eventos como fatos criptográficos,  
> não como decisões sociais.**

---

## ⚖️ Licença

Apache License 2.0

Autor da especificação conceitual:  
**Armando Freire**

---

> PoE existe para registrar eventos como fatos criptográficos,  
> não como decisões sociais.

