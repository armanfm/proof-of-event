# Proof of Event (PoE)
## Especificação Técnica Oficial  
### Camada 2 — Ledger Determinístico

- **Versão:** 0.1  
- **Status:** Fundação Técnica  
- **Autor:** Armando José Freire de Melo  
- **Licença:** Apache License 2.0  

---

## 1. Escopo do Protocolo

O Proof of Event (PoE) é um protocolo determinístico para registro de eventos externos,
cujo objetivo é produzir uma **prova criptográfica imutável, auditável e reexecutável
de existência temporal**.

O PoE **não valida significado**, **não decide verdade** e **não resolve disputas**.  
Ele registra eventos cuja ocorrência e validação **já foram tratadas fora do protocolo**
(Camada 1).

Esta especificação define **exclusivamente a Camada 2 — Ledger Determinístico PoE**
e o papel dos **Certificadores PoE**.

---

## 2. Princípios Fundamentais

### 2.1 Determinismo Absoluto

Dada a mesma entrada válida, qualquer implementação compatível do PoE **DEVE**
produzir exatamente a mesma prova criptográfica.

Não existe aleatoriedade, votação ou interpretação subjetiva.

---

### 2.2 Ausência de Consenso

O PoE **não implementa consenso distribuído**.

Não existem:
- votação
- mineração
- staking
- slashing
- forks
- governança on-chain

O protocolo não resolve conflitos sociais, jurídicos ou semânticos.

---

### 2.3 Ancoragem Temporal Canônica

A prova produzida pelo PoE baseia-se **exclusivamente na ancoragem temporal**.

O protocolo atribui um **timestamp canônico**, gerado pelo Certificador PoE
no momento da aceitação do evento.

A ordenação entre eventos **não é objetivo do PoE** e **não possui valor probatório**.
Qualquer relação de ordem é consequência operacional do certificador,
não propriedade do protocolo.

---

### 2.4 Imutabilidade Append-Only

Eventos aceitos **NUNCA** são alterados, removidos ou reescritos.

O ledger é estritamente append-only dentro de cada certificador.

---

### 2.5 Neutralidade do Protocolo

O PoE é cego a:
- identidade social
- reputação
- valor econômico
- conteúdo semântico do evento

O protocolo aplica apenas regras **mecânicas e determinísticas**.

---

## 3. Definições

- **Evento Externo (Camada 1):** Fato ocorrido fora do PoE.
- **Certificador PoE:** Entidade que executa o protocolo PoE, recebe eventos validados off-chain e emite provas PoE.
- **Evento Canônico:** Representação determinística reduzida do evento externo.
- **Timestamp Canônico:** Marca temporal gerada pelo Certificador PoE no momento da aceitação.
- **Prova PoE (PoE_Proof):** Associação criptográfica entre o hash do evento e o timestamp canônico.
- **Ledger PoE:** Registro append-only das provas temporais emitidas por um Certificador PoE.
- **Certificador_ID:** Identificador estável e imutável do certificador.
- **Token PoE (Opcional):** Unidade econômica utilizada exclusivamente para liquidação operacional do uso do protocolo.

---

## 4. Arquitetura do Sistema

### 4.1 Camada 1 — Evento Externo (Fora do Escopo)

A validação, auditoria e responsabilidade legal do evento **ocorrem fora do PoE**.

O protocolo assume que:
- houve verificação off-chain; e
- o evento foi reduzido a um hash determinístico.

---

### 4.2 Camada 2 — Certificador PoE (Este Protocolo)

A Camada 2 é executada por um **Certificador PoE**, responsável por:

- receber eventos canônicos;
- atribuir timestamp canônico;
- gerar a prova PoE;
- registrar a prova em um ledger determinístico append-only;
- emitir recibos verificáveis.

O PoE **não define nem impõe**:
- rede distribuída;
- replicação entre certificadores;
- consenso entre certificadores;
- topologia de infraestrutura.

Cada certificador opera de forma **soberana e independente**.

---

### 4.3 Certificadores PoE

Certificadores PoE são entidades responsáveis por prestar o serviço de
certificação temporal de eventos externos.

Eles podem ser instituições, oráculos, empresas, órgãos públicos ou operadores
autorizados, conforme o contexto de uso.

O protocolo PoE **não define critérios de escolha, reputação ou autoridade**
dos certificadores — essas relações existem fora do sistema.

Cada certificador:
- certifica apenas eventos sob sua responsabilidade;
- responde legal e operacionalmente pelos eventos que aceita;
- não participa de consenso;
- não valida ou invalida eventos de outros certificadores.

Múltiplos certificadores existem para refletir a realidade do mundo externo:
diferentes eventos exigem diferentes entidades responsáveis.

---

## 5. Prova Canônica PoE

### 5.1 Definição Formal

A prova PoE é definida exclusivamente por:

PoE_Proof = SHA-512(payload_hash || timestamp_canônico)


Onde:
- `payload_hash` é fornecido externamente;
- `timestamp_canônico` é atribuído pelo Certificador PoE.

Essa é a **unidade mínima e suficiente de prova**.

---

### 5.2 Propriedades da Prova

A prova PoE é:
- imutável;
- verificável independentemente;
- reexecutável;
- resistente a interpretação semântica.

---

## 6. Verificação da Prova PoE

A verificação de uma prova PoE consiste em:

1. Obter o `payload_hash`
2. Obter o `timestamp_canônico`
3. Recomputar:

