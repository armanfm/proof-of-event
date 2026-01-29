# Proof of Event (PoE)
## Especificação Técnica Oficial  
### Camada 2 — Ledger Determinístico

- **Versão:** 0.1  
- **Status:** Fundação Técnica  
- **Autor:** Armando Freire  
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
- **Certificador_ID:** Identificador único, estável e imutável de um Certificador PoE.
- **Token PoE (Opcional):** Unidade econômica utilizada exclusivamente para liquidação operacional do uso do protocolo.

---

### 3.1 Formato do Certificador_ID

O `certificador_id` **DEVE**:
- Ser determinístico, único e imutável
- Ser computável exclusivamente a partir do evento GENESIS
- Permanecer constante durante toda a existência do certificador

O formato **RECOMENDADO** é:

certificador_id = "poe:" || SHA-512(genesis_serializado)


Onde:
- `genesis_serializado` é a serialização determinística completa do evento GENESIS
- O hash SHA-512 resulta em **128 caracteres hexadecimais**

Exemplo:

poe:a1b2c3d4e5f6...<128 hex>


Outros formatos **SÃO PERMITIDOS**, desde que sejam:
- determinísticos;
- imutáveis após a criação;
- computáveis a partir do GENESIS;
- únicos sem colisões práticas.

O protocolo PoE **não valida identidade**, apenas registra o identificador declarado.

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

## 5. Prova Canônica PoE

### 5.1 Definição Formal

A prova PoE é definida exclusivamente por:

PoE_Proof = SHA-512(payload_hash || timestamp_canônico)


Essa é a **unidade mínima e suficiente de prova**.

---

### 5.2 Propriedades da Prova

A prova PoE é:
- imutável;
- verificável independentemente;
- reexecutável;
- semanticamente neutra.

---

## 6. Verificação da Prova PoE

A verificação de uma prova PoE consiste em:

1. Obter o `payload_hash`
2. Obter o `timestamp_canônico`
3. Recomputar:

PoE_Proof' = SHA-512(payload_hash || timestamp_canônico)

4. Comparar com a `poe_proof` apresentada

Se os valores coincidirem, a prova é válida.

---

## 7. Formato do Evento Canônico

Todo evento submetido ao PoE **DEVE** conter, no mínimo:

- `payload_hash`

Campos adicionais podem existir, desde que **NÃO interfiram**
na definição da prova canônica.

### 7.1 Formato do `payload_hash`

O `payload_hash` **DEVE**:
- ser gerado utilizando **SHA-512**;
- ser codificado em **hexadecimal lowercase**;
- possuir comprimento fixo de **128 caracteres hexadecimais**;
- representar exclusivamente o conteúdo do evento externo.

O protocolo assume a resistência a colisões do algoritmo SHA-512.
Eventos com `payload_hash` idêntico são tratados como o mesmo evento.

---

## 8. Timestamp Canônico

O timestamp canônico **DEVE**:
- ser gerado por fonte de tempo confiável do certificador;
- ser expresso em UTC;
- utilizar o formato ISO 8601 estendido:
`YYYY-MM-DDTHH:MM:SS.sssZ`;
- possuir precisão mínima de milissegundos;
- representar o instante de aceitação do evento.

---

## 9. Aceitação de Eventos (Normativo)

Um evento é aceito pelo PoE se, e somente se:
- o formato canônico é respeitado;
- o timestamp canônico é atribuído;
- requisitos operacionais (ex: pagamento, se aplicável) são atendidos.

Não existe rejeição baseada em conteúdo semântico.

---

## 10. Eventos Arbitrários e Correções

O PoE não diferencia eventos de produção, teste ou erro.

Eventos aceitos **NUNCA** são removidos.
Correções ocorrem apenas por **novos eventos**.

---

## 11. Token PoE (Camada Operacional Opcional)

Quando utilizado:
- o pagamento **DEVE** ocorrer antes da aceitação;
- a falha no pagamento **É** motivo válido de rejeição;
- dados econômicos **PODEM** constar como metadata.

O token não participa da prova criptográfica.

---

## 12. GENESIS

O evento GENESIS representa a inicialização do Certificador PoE.

O GENESIS **DEVE**:
- ser o primeiro registro do ledger;
- possuir `payload_hash` constante (ex: SHA-512("POE_GENESIS"));
- possuir timestamp canônico inicial;
- PODE conter metadata do certificador.

Na versão 0.1, o algoritmo de hash é **FIXO em SHA-512**.

---

## 13. Recibo PoE (Opcional)

Um Certificador PoE **PODE** emitir um recibo verificável.

O recibo **DEVE** conter:
- `poe_proof`
- `payload_hash`
- `timestamp_canônico`
- `certificador_id`
- `version`

### 13.1 Assinatura Criptográfica

A assinatura é **OPCIONAL**.

Quando presente:
- DEVE usar algoritmo PQC;
- autentica o certificador;
- NÃO altera a prova.

---

### 13.2 Serialização Determinística

Objetos que participem de hash ou assinatura **DEVEM** ser serializados
deterministicamente:

- UTF-8
- Campos ordenados alfabeticamente
- JSON compacto (sem espaços)
- Números como inteiros decimais

---

## 14. Retenção do Ledger

Um Certificador PoE **DEVE** manter todo o histórico desde o GENESIS.

Retenção parcial é **PROIBIDA**.

---

## 15. Versionamento

Mudanças:
- **DEVEM** incrementar versão;
- **NUNCA** invalidam provas antigas.

---

## 16. Encerramento

O Proof of Event existe para registrar eventos como **fatos criptográficos**,
não como decisões sociais.

**A blockchain não decide. Ela testemunha.**
