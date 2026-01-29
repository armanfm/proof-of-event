# Proof of Event (PoE)
## Especificação Técnica Oficial  
### Camada 2 — Ledger Determinístico

- **Versão:** 0.1  
- **Status:** Fundação Técnica  
- **Autor:** Armando José Freire de Melo  
- **Licença:** Apache License 2.0  

---

## 1. Escopo do Protocolo

O Proof of Event (PoE) é um protocolo determinístico para registro de eventos externos, cujo objetivo é produzir uma **prova criptográfica imutável, auditável e reexecutável de existência temporal**.

O PoE **não valida significado**, **não decide verdade** e **não resolve disputas**.  
Ele registra eventos cuja ocorrência e validação **já foram tratadas fora do protocolo** (Camada 1).

Esta especificação define **exclusivamente a Camada 2 — Ledger Determinístico PoE** e o papel dos **Certificadores PoE**.

---

## 2. Princípios Fundamentais

### 2.1 Determinismo Absoluto

Dada a mesma entrada válida, qualquer implementação compatível do PoE **DEVE** produzir exatamente a mesma prova criptográfica.

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

O protocolo atribui um **timestamp canônico**, gerado pelo próprio sistema, no momento da aceitação do evento.

A ordenação entre eventos **não é objetivo do PoE** e **não possui valor probatório**.

---

### 2.4 Imutabilidade Append-Only

Eventos aceitos **NUNCA** são alterados, removidos ou reescritos.

O ledger é estritamente append-only.

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
- **Evento Canônico:** Representação determinística do evento externo.
- **Timestamp Canônico:** Marca temporal gerada pelo Certificador PoE no momento da aceitação.
- **Prova PoE:** Associação criptográfica entre o hash do evento e o timestamp canônico.
- **Ledger PoE:** Registro append-only das provas temporais emitidas por um Certificador PoE.
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

## 5. Prova Canônica PoE

### 5.1 Definição Formal

A prova PoE é definida exclusivamente por:

PoE_Proof = HASH(payload_hash || timestamp_canônico)


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

## 6. Formato do Evento Canônico

Todo evento submetido ao PoE **DEVE** conter, no mínimo:

- `payload_hash`;
- metadados de versão (quando aplicável).

Campos adicionais podem existir, desde que **NÃO interfiram** na definição da prova canônica.

---

## 7. Timestamp Canônico

- O timestamp canônico é gerado exclusivamente pelo Certificador PoE.
- Ele é a **única referência temporal válida** da prova.
- Timestamps externos são apenas informativos.

O nível de precisão (milissegundos, nanossegundos, etc.) é decisão do certificador, desde que determinística.

---

## 8. Aceitação de Eventos (Normativo)

Um evento é aceito pelo PoE se, e somente se:

- o formato canônico é respeitado;
- o timestamp canônico é atribuído pelo Certificador PoE;
- requisitos operacionais (ex: pagamento, se aplicável) são atendidos.

Não existe rejeição baseada em conteúdo semântico.

---

## 9. Eventos Arbitrários e Conteúdo “Errado”

O PoE **não diferencia** eventos de produção, teste, erro ou experimento.

Qualquer evento que:
- respeite o formato;
- cumpra as regras mecânicas;

**DEVE** ser aceito pelo certificador.

Correções ou invalidações ocorrem **exclusivamente por novos eventos**, preservando a trilha de auditoria.

---

## 10. Token PoE (Camada Operacional Opcional)

O Token PoE pode ser utilizado para:
- pagamento pelo uso do serviço de certificação;
- liquidação de custos operacionais.

O token:
- não é emitido pelo núcleo PoE;
- não participa da prova criptográfica;
- não confere governança;
- não promete retorno financeiro.

---

## 11. Segurança

A segurança do PoE deriva de:
- determinismo;
- reexecução independente;
- imutabilidade do ledger;
- verificabilidade criptográfica da prova.

Confiança é substituída por verificação.

---

## 12. Versionamento

Mudanças no protocolo:
- **DEVEM** incrementar a versão;
- **DEVEM** declarar compatibilidade;
- **NUNCA** alteram provas já emitidas.

---

## 13. Retenção do Ledger (Normativo)

Um Certificador PoE **DEVE** manter a integridade completa de seu ledger determinístico desde o GENESIS até a prova mais recente.

A retenção parcial do histórico **é proibida** dentro de uma mesma instância certificadora.

---

## 14. Encerramento

O Proof of Event existe para registrar eventos como **fatos criptográficos**,  
não como decisões sociais.

**A blockchain não decide. Ela testemunha.**
