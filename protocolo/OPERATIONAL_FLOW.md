# Proof of Event (PoE)
## Protocol v0.1 — Fluxo Operacional do Certificador
(Event → Timestamp → Append)

**Versão:** 0.1  
**Status:** Operacional (executável no papel)  
**Compatível com:** `SPEC.md` v0.1  
**Autor:** Armando Freire  

---

## 1. Objetivo deste Documento

Este documento define o **fluxo operacional do Proof of Event (PoE) v0.1**,
descrevendo como um evento externo, já verificado fora do protocolo,
é submetido a um **Certificador PoE**, recebe um **timestamp canônico**
e é registrado em um **ledger determinístico append-only**.

Este documento **NÃO altera** o `SPEC.md`.
Ele apenas remove ambiguidades de implementação,
explicitando a sequência de operações,
os artefatos produzidos e o comportamento esperado sob falhas.

---

## 2. Entidades

O PoE v0.1 envolve apenas as seguintes entidades:

- **Client (Verificador / Sistema Externo)**  
  Entidade responsável por verificar o evento fora do PoE e submeter seu hash.

- **Certificador PoE**  
  Entidade que executa o protocolo PoE:
  atribui timestamp canônico, gera a prova e mantém o ledger soberano.

- **Ledger do Certificador**  
  Registro local, determinístico e append-only das provas PoE emitidas.

Não existe consenso, coordenação ou ordenação global entre certificadores.

---

## 3. Artefatos

### 3.1 Evento Canônico (Entrada)

O evento submetido ao PoE consiste, no mínimo, em:

- `payload_hash` — hash criptográfico do evento externo
- metadados opcionais (fora do núcleo)

O conteúdo semântico do evento **não é avaliado** pelo PoE.

---

### 3.2 Timestamp Canônico

O timestamp canônico:

- é gerado exclusivamente pelo Certificador PoE;
- representa o instante de aceitação do evento;
- é expresso em UTC;
- possui precisão mínima de milissegundos;
- segue o formato ISO 8601:  
  `YYYY-MM-DDTHH:MM:SS.sssZ`

Timestamps fornecidos pelo Client são apenas informativos
e **não possuem valor probatório** no PoE.

---

### 3.3 Prova PoE (PoE_Proof)

A prova PoE é definida como:



PoE_Proof = SHA-512(payload_hash || timestamp_canônico)

Esta é a **única unidade mínima de prova** do protocolo.

---

### 3.4 Recibo PoE (Saída)

Após a aceitação do evento, o Certificador PoE **DEVE** retornar um recibo
contendo, no mínimo:

- `poe_proof`
- `payload_hash`
- `timestamp_canônico`
- `certificador_id`
- `version`

Metadados adicionais (ex.: assinatura PQC, custo pago, referências externas)
são opcionais e **não fazem parte da prova criptográfica**.

---

## 4. Fluxo Operacional

### 4.1 Preparação do Evento (Client)

O Client:

1. verifica o evento fora do PoE (Camada 1);
2. gera o `payload_hash` de forma determinística;
3. prepara a submissão ao Certificador PoE.

O PoE **não participa** dessa verificação.

---

### 4.2 Submissão ao Certificador

O Client envia ao Certificador PoE:

- `payload_hash`;
- metadados opcionais;
- comprovação de pagamento, se aplicável (fora do núcleo).

---

### 4.3 Aceitação do Evento (Certificador)

Ao receber uma submissão válida, o Certificador PoE:

1. valida o formato mínimo exigido;
2. verifica requisitos operacionais (ex.: pagamento, se aplicável);
3. gera o timestamp canônico;
4. calcula a PoE_Proof;
5. registra o evento no ledger append-only;
6. emite o Recibo PoE.

Não existe rejeição baseada em conteúdo semântico.

---

### 4.4 Append no Ledger

O ledger do Certificador:

- é estritamente append-only;
- mantém registros imutáveis;
- preserva a ordem local de aceitação;
- não sofre reescrita ou exclusão.

A ordem do ledger **não possui valor semântico ou probatório global**;
ela reflete apenas a sequência local de aceitação.

---

## 5. Pagamento (Opcional)

Quando utilizado, o Token PoE:

- é requisito operacional para aceitação do evento;
- deve ser liquidado antes da emissão da prova;
- **não participa** do cálculo da PoE_Proof;
- **não interfere** no determinismo do ledger.

A falha no pagamento **PODE** resultar em rejeição.

---

## 6. Assinaturas Criptográficas (Opcional)

Assinaturas criptográficas:

- são externas ao núcleo do PoE;
- podem utilizar algoritmos clássicos ou pós-quânticos (PQC);
- autenticam o Certificador, não a prova.

A ausência de assinatura **NÃO invalida** uma prova PoE.

Assinaturas **não fazem parte do caminho crítico**
de aceitação de eventos.

---

## 7. Armazenamento Externo (Opcional)

Dados completos do evento podem ser armazenados externamente
(ex.: IPFS, Pinata, bancos de dados institucionais).

Esses sistemas:

- são opcionais;
- não fazem parte do protocolo PoE;
- não afetam a validade da prova.

A indisponibilidade de armazenamento externo
**NÃO invalida** provas PoE já emitidas.

O ledger do Certificador é a fonte de verdade.

---

## 8. Falhas Operacionais

### 8.1 Certificador Indisponível

Se um Certificador estiver indisponível:

- novas submissões não são aceitas;
- provas já emitidas permanecem válidas;
- ao retornar, o ledger continua a partir do último estado.

---

### 8.2 Falhas de Rede ou Cliente

Falhas de conexão do Client:

- não invalidam provas já emitidas;
- exigem nova submissão, se necessário;
- não geram deduplicação automática no núcleo
  (políticas de retry são responsabilidade do Client).

---

## 9. Eventos Incorretos ou de Teste

O PoE aceita eventos de teste ou conteúdo incorreto se:

- o formato mínimo for respeitado;
- os requisitos operacionais forem atendidos.

Eventos aceitos **NUNCA** são apagados.

Correções ou invalidações devem ocorrer
por **novos eventos**, preservando a trilha histórica.

---

## 10. Garantias do Protocolo

O PoE v0.1 garante:

- determinismo;
- imutabilidade;
- verificabilidade independente;
- separação clara de responsabilidades.

O PoE **não garante**:

- verdade semântica;
- legitimidade jurídica;
- consenso entre entidades;
- ordenação global de eventos.

---

## 11. Encerramento

O fluxo operacional do Proof of Event é intencionalmente simples:

**Sistemas verificam.  
Certificadores testemunham.  
Ledgers preservam.**

O PoE não decide.
Ele apenas registra, de forma imutável,
eventos que já foram aceitos fora dele.

**A blockchain não decide.  
Ela testemunha.**




