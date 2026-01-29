
# Proof of Event (PoE)

**A blockchain não decide. Ela testemunha.**

Proof of Event (PoE) é um protocolo determinístico para registrar eventos externos como **fatos criptográficos ancorados no tempo**, sem consenso, sem votação e sem interpretação semântica on-chain.

PoE **não cria verdade**.  
PoE **não interpreta significado**.  
PoE **testemunha eventos cuja validade já existe fora do sistema**.

---

## 🎯 Objetivo do Projeto

O PoE foi projetado para ambientes onde:

- o consenso sobre o evento já existe fora do sistema;
- auditoria, rastreabilidade e reexecução são mais importantes que governança;
- mecanismos como PoW, PoS, staking ou votação são indesejáveis;
- simplicidade, determinismo e compatibilidade institucional são requisitos.

### Exemplos de uso

- registros institucionais
- eventos legais
- logs auditáveis
- sensores e sistemas industriais
- provas de execução e ocorrência

---

## 🧱 Arquitetura (Visão Geral)

O Proof of Event é dividido em **camadas estritamente desacopladas**:

### Camada 1 — Evento Externo (Fora do Escopo)

Onde o evento ocorre.

- validação
- auditoria
- responsabilidade
- verificação

Tudo acontece **antes** do PoE.

---

### Camada 2 — Certificação Temporal Determinística (PoE)

Executada por **Certificadores PoE**.

O Certificador:
- recebe o hash do evento;
- atribui um **timestamp canônico**;
- gera uma **prova PoE**;
- registra a prova em um ledger append-only;
- emite um recibo verificável.

> ⚠️ **Não existe rede PoE**.  
> Cada certificador opera de forma independente.


### Por que múltiplos certificadores?

O Proof of Event não opera como uma rede única nem como um sistema de consenso.
Cada Certificador PoE atua de forma independente, como prestador de serviço de
certificação temporal.

Certificadores são responsáveis apenas pelos eventos que certificam e não
precisam concordar entre si. A escolha de um certificador é externa ao protocolo
e depende de fatores como confiança institucional, relação contratual,
jurisdição, reputação ou exigências regulatórias.

O PoE registra o fato criptográfico de que um evento foi certificado por uma
entidade específica em um determinado momento — não que múltiplas entidades
concordaram sobre ele.

---

### Camada 3 — Camadas Semânticas (Opcional)

Camadas externas podem:
- interpretar eventos;
- integrar sistemas;
- aplicar regras de negócio;
- enriquecer metadados.

Essas camadas **NUNCA** interferem na prova PoE.

---

## 🔐 O que é a Prova PoE?

A prova PoE é definida por:

PoE_Proof = HASH(payload_hash || timestamp_canônico)


Ela prova que:

> “Este evento existia **até** este momento no tempo.”

Nada mais. Nada menos.

---

## 🧠 O que o PoE NÃO é

- não é uma blockchain
- não é um protocolo de consenso
- não é uma DAO
- não é um sistema de governança
- não é um árbitro de verdade
- não promete retorno financeiro
- não recompensa usuários finais

---

## 💰 Modelo Econômico (Visão Geral)

O PoE pode operar com uma unidade econômica opcional (Token PoE) para:

- pagamento pelo uso do serviço de certificação;
- liquidação de custos operacionais.

Princípios:

- o token **não faz parte da prova**;
- o PoE não emite tokens;
- o preço é definido externamente;
- não existe promessa de valorização;
- o token não confere governança.

---

## 📜 Especificação Técnica

A definição formal, normativa e completa do protocolo está em:

➡️ **[`SPEC.md`](./SPEC.md)**

O SPEC é a **fonte única de verdade técnica**.

---

## 🔬 Status do Projeto

- 🧠 Fundação conceitual: consolidada
- 📐 Especificação técnica: definida (v0.1)
- ⚙️ Implementação de referência: em desenvolvimento
- 💰 Modelo econômico: definido em nível conceitual


---

## ⚖️ Licença

Apache License 2.0

Autor da especificação conceitual:  
**Armando Freire**

---

> PoE existe para registrar eventos como fatos criptográficos,  
> não como decisões sociais.

