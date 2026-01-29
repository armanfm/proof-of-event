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

PoE_Proof' = SHA-512(payload_hash || timestamp_canônico)

4. Comparar `PoE_Proof'` com a `poe_proof` apresentada

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
Eventos com `payload_hash` idêntico são tratados como o mesmo evento
do ponto de vista criptográfico.

---

## 8. Timestamp Canônico

O timestamp canônico **DEVE**:
- ser gerado por uma fonte de tempo confiável controlada pelo Certificador PoE;
- ser expresso em UTC;
- utilizar o formato ISO 8601 estendido:
`YYYY-MM-DDTHH:MM:SS.sssZ`;
- possuir precisão mínima de milissegundos;
- representar o instante de aceitação do evento pelo certificador.

---

## 9. Aceitação de Eventos (Normativo)

Um evento é aceito pelo PoE se, e somente se:
- o formato canônico é respeitado;
- o timestamp canônico é atribuído pelo Certificador PoE;
- requisitos operacionais (ex: pagamento, se aplicável) são atendidos.

Não existe rejeição baseada em conteúdo semântico.

---

## 10. Eventos Arbitrários e Conteúdo “Errado”

O PoE **não diferencia** eventos de produção, teste, erro ou experimento.

Qualquer evento que:
- respeite o formato;
- cumpra as regras mecânicas;

**DEVE** ser aceito pelo certificador.

Correções ou invalidações ocorrem **exclusivamente por novos eventos**,
preservando a trilha de auditoria.

---

## 11. Token PoE (Camada Operacional Opcional)

Quando utilizado, o Token PoE pode servir como requisito operacional
para submissão de eventos.

Nesse caso:
- o pagamento **DEVE** ser concluído antes da aceitação;
- a falha no pagamento **É** motivo válido de rejeição;
- informações econômicas **PODEM** ser registradas como metadata opcional.

O token:
- não é emitido pelo núcleo PoE;
- não participa da prova criptográfica;
- não confere governança;
- não promete retorno financeiro.

---

## 12. GENESIS

O evento GENESIS representa a inicialização de um Certificador PoE.

O GENESIS **DEVE**:
- ser o primeiro registro do ledger do certificador;
- possuir um `payload_hash` constante
(ex: SHA-512("POE_GENESIS"));
- possuir um timestamp canônico correspondente ao início de operação;
- PODE conter metadata identificando o certificador,
sua versão inicial e parâmetros operacionais.

Na versão 0.1 do protocolo, o algoritmo de hash do payload
é **FIXO em SHA-512**, independentemente de declarações no GENESIS.

---

## 13. Recibo PoE (Opcional)

Um Certificador PoE **PODE** emitir um recibo verificável
associado a uma prova PoE.

Um recibo PoE **DEVE** conter:
- `poe_proof`
- `timestamp_canônico`
- `payload_hash`
- `certificador_id`
- `version`

Um recibo PoE **PODE** conter:
- `metadata` (dados auxiliares, fora da prova)

### 13.1 Assinatura Criptográfica

A assinatura criptográfica de um recibo PoE é **OPCIONAL**
e externa ao núcleo do protocolo.

Quando presente:
- o recibo **DEVE** ser assinado utilizando
um algoritmo de criptografia pós-quântica (PQC);
- o algoritmo específico é escolha do certificador;
- a assinatura autentica o certificador,
**não altera nem integra** a PoE_Proof.

A ausência de assinatura **NÃO invalida** a prova PoE.

---

## 14. Retenção do Ledger (Normativo)

Um Certificador PoE **DEVE** manter a integridade completa
de seu ledger determinístico desde o GENESIS
até a prova mais recente.

A retenção parcial do histórico **É PROIBIDA**
dentro de uma mesma instância certificadora.

---

## 15. Versionamento e Compatibilidade

Mudanças no protocolo:
- **DEVEM** incrementar a versão;
- **DEVEM** declarar compatibilidade ou incompatibilidade;
- **NUNCA** invalidam provas já emitidas.

Versões são compatíveis quando:
- mantêm a mesma definição de `PoE_Proof`;
- não alteram regras de aceitação;
- não invalidam provas anteriores.

Mudanças incompatíveis **EXIGEM**
novo identificador de protocolo.

---

## 16. Encerramento

O Proof of Event existe para registrar eventos como
**fatos criptográficos**, não como decisões sociais.

**A blockchain não decide. Ela testemunha.**
