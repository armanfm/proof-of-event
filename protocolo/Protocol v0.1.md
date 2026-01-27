# Proof of Event (PoE)
## Protocol v0.1 — Fluxo Operacional (FIFO → Broadcast → Append)

**Versão:** 0.1  
**Status:** Operacional (executável no papel)  
**Compatível com:** `SPEC.md` v0.1  
**Autor:** Armando José Freire de Melo  

---

## 1. Objetivo deste Documento

Este documento define o **fluxo operacional** do PoE v0.1: como um evento
produzido por um verificador/oráculo entra no FIFO, é ordenado globalmente,
é distribuído aos armazenadores e é persistido no ledger *append-only*.

Este documento **não altera** o `SPEC.md`. Ele apenas remove ambiguidade
de implementação (ordem de chamadas, estados, erros e comportamento sob falhas).

---

## 2. Entidades

- **Client (Verificador/Oráculo):** prepara e submete `PoE_Event`.
- **FIFO Gateway (Ordenador Global):** porta de entrada soberana. Serializa.
- **Storage Node (Armazenador):** replica o ledger e pode emitir prova.
- **Ledger:** sequência *append-only* encadeada por hash.

---

## 3. Artefatos

### 3.1 Evento Canônico
Conforme `SPEC.md` v0.1:

<img width="391" height="243" alt="image" src="https://github.com/user-attachments/assets/b0c69a5d-7c49-43f4-8c92-28fc2df577e8" />


### 3.2 Receipt (Recibo de Enfileiramento)

O FIFO DEVE retornar um recibo quando aceitar um evento no FIFO:

<img width="420" height="201" alt="image" src="https://github.com/user-attachments/assets/aab0b73b-441d-4e05-8994-7459c956e6b4" />

-fifo_sequence é o número lógico global da fila (ordem canônica).

-observed_prev_hash é o hash do último evento aceito no momento do enqueue.

### 3.3 Commitment Proof (Opcional)

Conforme SPEC.md:

<img width="357" height="201" alt="image" src="https://github.com/user-attachments/assets/0c6760b2-110f-4773-a871-696316771b93" />

## 4. Estados do FIFO Gateway

O FIFO Gateway opera em um dos seguintes estados:

- **IDLE**: fila vazia, pronto para receber.
- **ACCEPTING**: recebendo e enfileirando eventos.
- **THROTTLED**: recebendo com limitação (rate limit ativo).
- **DEGRADED**: aceitando, porém com degradação de broadcast (ex.: poucos nós disponíveis).
- **DOWN**: indisponível (nenhuma submissão aceita).

**Observação:** o estado **NÃO** altera a regra de ordem. Ele controla apenas disponibilidade.

---

## 5. Códigos de Erro (Normativo)

O FIFO Gateway **DEVE** responder com códigos de erro determinísticos:

- `ERR_BAD_FORMAT` — evento não segue o formato canônico.
- `ERR_BAD_VERSION` — versão do evento não suportada.
- `ERR_BAD_PREV_HASH` — `previous_event_hash` não corresponde ao estado esperado.
- `ERR_NO_TOKEN` — pagamento insuficiente.
- `ERR_RATE_LIMIT` — limite de submissão excedido.
- `ERR_DUPLICATE_EVENT` — `event_id` já enfileirado/confirmado (deduplicação).
- `ERR_FIFO_DOWN` — FIFO indisponível.
- `ERR_INTERNAL` — erro interno (usar minimamente; preferir erros explícitos).

---

## 6. Fluxo Operacional  
### Submit → Enqueue → Broadcast → Append

### 6.1 Preparação do Evento (Client)

O Client (verificador/oráculo):

- verifica o evento fora do PoE (Camada 1);
- gera `payload_hash`;
- monta `PoE_Event`;
- define `previous_event_hash` como o hash do último evento conhecido.

**Nota:** se `previous_event_hash` estiver desatualizado, a submissão será rejeitada
com `ERR_BAD_PREV_HASH` (ver Seção 6.4).

---

### 6.2 Submissão ao FIFO (Client → FIFO)

O Client envia:

- `PoE_Event`;
- prova/transferência de pagamento em Token PoE (mecanismo definido na implementação).

O FIFO valida:

- formato canônico;
- versão suportada;
- pagamento suficiente;
- deduplicação mínima (ver Seção 7);
- coerência do `previous_event_hash` com o estado observado.

Se aprovado, o FIFO retorna `Enqueue_Receipt` com `accepted = true`.

---

### 6.3 Ordenação (FIFO)

O FIFO impõe ordem global única:

- incrementa `fifo_sequence` estritamente (`1, 2, 3, ...`);
- define a sequência canônica de eventos;
- prepara o evento para distribuição aos armazenadores.

O FIFO **NÃO** avalia significado do conteúdo.

---

### 6.4 Rejeição por `previous_event_hash`

Se o Client submeter um evento com `previous_event_hash` divergente:

- o FIFO responde `ERR_BAD_PREV_HASH`;
- o FIFO **NÃO** enfileira o evento.

O Client deve:

1. obter o último `previous_event_hash` canônico atual; e
2. reenviar o evento com o hash correto.

**Regra prática:** quem “caiu” e reenviou entra como submissão nova (vai ao final da fila),
exceto se deduplicação por `event_id` estiver ativa (ver Seção 7).

---

### 6.5 Broadcast (FIFO → Storage Nodes)

Ao aceitar um evento, o FIFO distribui aos armazenadores ativos:

- evento completo (`PoE_Event`);
- metadado de ordem (`fifo_sequence`);
- hash anterior observado (`observed_prev_hash`).

O broadcast pode ser:

- **push** (FIFO empurra), ou
- **pull** (armazenadores puxam o stream).

A ordem lógica **DEVE** ser idêntica em todos os armazenadores.

---

### 6.6 Append (Storage Node)

Cada armazenador:

- recebe o próximo evento canônico;
- verifica:
  - formato;
  - se `previous_event_hash` corresponde ao último evento local;
- grava o evento no ledger local (*append-only*);
- calcula o `event_hash`;
- opcionalmente emite `Commitment_Proof`.

Se a verificação falhar (ex.: evento pulado), o nó está fora de sincronia e **DEVE**
executar sincronização.

---

## 7. Regras de Deduplicação e Retry (Normativo)

Para evitar cobrança dupla por queda de conexão, o FIFO **DEVE** implementar:

### 7.1 Deduplicação por `event_id`

Se um `event_id` já foi aceito no FIFO e o Client reenviar:

- o FIFO **DEVE** responder `ERR_DUPLICATE_EVENT` **OU**
- retornar o mesmo `Enqueue_Receipt`.

A implementação **DEVE** documentar o comportamento escolhido.

---

### 7.2 Retry Seguro

- O Client deve considerar a submissão confirmada apenas após receber `Enqueue_Receipt`.
- Se a conexão cair sem receipt, o Client **PODE** reenviar o mesmo `event_id`.
- A deduplicação impede dupla cobrança lógica.

---

## 8. Nós Offline e Sincronização

### 8.1 Armazenador Offline

Se um armazenador ficar offline:

- ele perde eventos enquanto não acompanha o stream;
- ao retornar, **DEVE** sincronizar a lacuna.

---

### 8.2 Estratégia de Sincronização (Pull)

O armazenador deve:

1. consultar um endpoint de stream (`from_sequence = last_local_sequence + 1`);
2. baixar os eventos faltantes;
3. aplicar os eventos em ordem até alcançar o estado canônico.

Se houver divergência de hash durante o replay, o nó **DEVE**:

- descartar a cópia inconsistente; e
- refazer a sincronização a partir de um checkpoint confiável.

---

## 9. Falhas do FIFO (Disponibilidade)

### 9.1 FIFO Indisponível

Se o FIFO estiver `DOWN`:

- nenhuma submissão nova é aceita;
- o ledger permanece válido (apenas não cresce);
- quando o FIFO retorna, a ordem continua a partir do último `fifo_sequence`.

---

### 9.2 Garantia do Protocolo

O PoE v0.1 define o FIFO como **ordenador soberano**.

Alta disponibilidade (HA), failover e descoberta de endpoint são detalhes de implantação
e **DEVEM** ser documentados em `docs/ARCHITECTURE.md`.

---

## 10. Regras de “Evento Errado” (Normativo)

O PoE aceita eventos de teste ou conteúdo semanticamente incorreto se:

- o formato canônico estiver correto;
- o pagamento tiver sido efetuado; e
- a ordem FIFO for respeitada.

Eventos aceitos **NUNCA** são apagados.
Correções ou invalidações ocorrem por **novo evento** referenciando o anterior.

---

## 11. Interface Mínima (Sugerida)

- `POST /submit` — submete `PoE_Event` e retorna `Enqueue_Receipt` ou erro.
- `GET /head` — retorna `last_event_hash` e `last_fifo_sequence`.
- `GET /stream?from_sequence=N` — retorna eventos em ordem a partir de `N`.

---

## 12. Encerramento

Este protocolo v0.1 define um fluxo simples e determinístico:

**Oráculos verificam. FIFO ordena. Armazenadores replicam.**

**A blockchain não decide. Ela testemunha.**



