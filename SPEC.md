# Proof of Event (PoE)
## Especificação Técnica Oficial (Camada 2 — Ledger Determinístico)

**Versão:** 0.1  
**Status:** Fundação Técnica  
**Autor:** Armando José Freire de Melo  
**Licença:** Apache License 2.0  

---

## 1. Escopo do Protocolo

O Proof of Event (PoE) é um protocolo determinístico para registro de eventos externos, cujo objetivo é produzir uma prova criptográfica **imutável, auditável e reexecutável** de que um evento foi registrado.

O PoE **não valida significado**, **não decide verdade** e **não resolve disputas**. Ele registra eventos cuja ocorrência e validação **já foram tratadas fora do protocolo** (Camada 1).

Esta especificação define **exclusivamente** a **Camada 2 — Ledger Determinístico PoE**.

---

## 2. Princípios Fundamentais

1. **Determinismo Absoluto**  
   Dada a mesma sequência de entradas válidas (na mesma ordem), qualquer implementação PoE **DEVE** produzir exatamente o mesmo ledger e os mesmos hashes.

2. **Ausência de Consenso**  
   Não existe votação, mineração, staking, slashing, fork choice, nem resolução de conflitos on-chain.

3. **Ordem Global Única (FIFO Soberano)**  
   Existe uma única linha do tempo canônica. A ordem dos eventos é determinada exclusivamente pela ordem de submissão ao FIFO global.

4. **Imutabilidade Append-Only**  
   Eventos aceitos **NUNCA** são alterados ou removidos.

5. **Neutralidade do Protocolo**  
   O protocolo é cego a reputação, saldo, status social ou poder econômico. O protocolo aplica regras mecânicas: **formato + pagamento + ordem**.

---

## 3. Definições

- **Evento Externo (Camada 1):** Fato ocorrido fora do PoE.
- **Verificador / Oráculo:** Entidade que valida o evento off-chain e prepara a submissão ao PoE.
- **Evento Canônico:** Representação determinística e estruturada do evento externo.
- **Ledger PoE:** Registro sequencial e encadeado de eventos aceitos.
- **FIFO Global:** Porta de entrada soberana que serializa a escrita e define a ordem canônica.
- **Armazenador (Storage Node):** Nó que replica o ledger PoE e pode emitir provas de armazenamento.
- **Token PoE:** Criptomoeda nativa usada para pagar o uso do FIFO e redistribuir remuneração operacional.

---

## 4. Arquitetura do Sistema

### 4.1 Camada 1 — Evento Externo (Fora do Escopo)
A validação, auditoria e responsabilidade legal do evento ocorrem **antes** da submissão ao PoE.

O protocolo PoE assume que:
- houve um processo de verificação off-chain; e
- o evento foi convertido para um formato canônico.

### 4.2 Camada 2 — Ledger Determinístico PoE (Este Protocolo)
A Camada 2 é um ledger *append-only* com:
- **ordem global única por FIFO**, e
- **encadeamento criptográfico** entre eventos.

Não existem forks. Se um nó diverge, ele está errado e deve sincronizar.

---

## 5. Evento Canônico (Formato)

Todo evento submetido ao PoE **DEVE** seguir um formato rígido e versionado.

<img width="472" height="389" alt="image" src="https://github.com/user-attachments/assets/e359b95b-47e8-49d0-bd6c-6b9a3725231a" />

## 5.1 Regras Normativas do Formato

- Todos os campos são obrigatórios.
- A ordem dos campos é fixa.
- `previous_event_hash` **DEVE** referenciar o hash do último evento aceito no ledger canônico no momento da submissão ao FIFO.
- Eventos malformados **DEVEM** ser rejeitados pelo FIFO e **NUNCA** entram no ledger.

## 5.2 Observação sobre `local_timestamp`

- `local_timestamp` é informativo (auditoria/correlação).
- `local_timestamp` **NÃO** define ordenação.
- A ordenação é definida **exclusivamente** pelo FIFO global.

---

## 6. Ordem Global — FIFO Soberano

### 6.1 Propriedade

- Eventos são processados **exclusivamente** pela ordem de chegada ao FIFO global.
- Não existe prioridade comprável.
- Não existe reordenação.
- Não existe paralelismo lógico de entrada.

### 6.2 Serialização de Escrita

- Participantes submetem eventos ao FIFO.
- Ninguém “escreve direto” no ledger.
- O FIFO libera eventos **um por vez**, impondo uma ordem única.
- Todos os armazenadores ativos recebem a mesma sequência liberada pelo FIFO.

---

## 7. Pagamento para Submissão (Uso do FIFO)

### 7.1 Regra Geral

- Para submeter um evento ao FIFO, o verificador/oráculo **DEVE** pagar uma taxa em Token PoE.
- O pagamento é condição necessária para entrada no FIFO.
- O FIFO **não cria**, **não emite** e **não destrói** tokens.
- Pagar o FIFO significa **transferir tokens** para redistribuição operacional.

### 7.2 Falta de Token

Se o verificador não possui token suficiente:

- o evento **NÃO** é aceito no FIFO;
- não há fila alternativa; e
- a submissão pode ser reexecutada posteriormente, quando houver token.

---

## 8. Eventos Arbitrários, de Teste e Conteúdo “Errado”

O PoE não diferencia eventos de produção, teste, experimento, erro ou conteúdo inválido do ponto de vista semântico.

Qualquer evento que:

- respeite o formato canônico,
- pague a taxa exigida, e
- siga a ordem FIFO,

**DEVE** ser aceito pelo ledger, independentemente do valor semântico do `payload_hash`.

O custo de submissão funciona como contenção natural de spam e uso abusivo.

### 8.1 Correção e Invalidação

Eventos aceitos **NUNCA** são removidos. Correções ou invalidações **DEVEM** ocorrer por meio de novos eventos que referenciem o evento anterior (via `event_id` no `payload_hash` ou por payload estruturado), preservando a trilha de auditoria.

---

## 9. Replicação e Função dos Armazenadores

### 9.1 Comportamento Esperado

Armazenadores:

- mantêm uma cópia completa e imutável do ledger;
- aplicam os eventos na ordem do FIFO global;
- não validam significado;
- não votam;
- não criam forks.

### 9.2 Nós Offline

Se um armazenador ficar offline:

- ele simplesmente para de acompanhar o ledger;
- perde remuneração por eventos que não armazenou; e
- ao retornar, **DEVE** sincronizar e reexecutar os eventos ausentes em ordem para voltar ao hash canônico.

---

## 10. Redistribuição do Token (Liquidação Operacional)

O Token PoE pago pelo uso do FIFO é redistribuído como remuneração operacional:

- **Armazenadores:** armazenamento, replicação e disponibilidade do ledger.
- **Verificadores/Oráculos:** trabalho off-chain (validação, tokenização, preparação).
- **Plataforma:** operação do ecossistema, manutenção e infraestrutura.

O token circula. Não há emissão adicional dentro do protocolo.

---

## 11. Prova de Compromisso (Opcional)

Ao registrar um evento, um armazenador **PODE** emitir uma prova de compromisso:

```plaintext
Commitment_Proof {
    event_hash:       bytes32
    store_node_id:    bytes32
    store_timestamp:  uint64
    signature:        bytes
}
<img width="304" height="208" alt="image" src="https://github.com/user-attachments/assets/8d15fce5-d31d-498e-b459-cdb090c02883" />
Essa prova atesta que um nó específico testemunhou e registrou aquele evento.

---

## 12. Condições de Aceitação (Normativo)

Um evento é aceito pelo PoE se, e somente se:

- segue o formato canônico;
- referencia corretamente o `previous_event_hash` esperado (conforme estado no FIFO);
- o pagamento em Token PoE foi efetuado; e
- respeita a ordem FIFO global.

---

## 13. O que o Protocolo NÃO Faz

O PoE:

- não interpreta eventos;
- não valida significado;
- não resolve disputas;
- não implementa governança;
- não cria consenso;
- não garante retorno financeiro, valorização ou rendimento do Token PoE.

---

## 14. Considerações de Segurança

- A segurança deriva de reexecução determinística e encadeamento por hash.
- A confiança é substituída por verificação.
- Qualquer divergência de hash entre nós indica divergência de estado; o nó divergente deve sincronizar.

---

## 15. Versionamento

Mudanças no protocolo:

- **DEVEM** incrementar `version`;
- **DEVEM** declarar compatibilidade ou quebra de compatibilidade; e
- **NUNCA** alteram eventos já registrados no ledger.

---

## 16. Encerramento

O Proof of Event existe para registrar eventos como fatos criptográficos, não como decisões sociais.

**A blockchain não decide. Ela testemunha.**

### Regra de Retenção Completa do Ledger (Normativa)

Um Armazenador PoE DEVE manter uma cópia completa e contínua
de TODOS os segmentos históricos do ledger determinístico,
desde o GENESIS até o evento mais recente.

A retenção parcial do ledger é PROIBIDA.

Um nó que:
- possua apenas parte dos ledgers diários;
- descarte ledgers históricos;
- inicie operação sem sincronizar todo o histórico;

É considerado FORA DE SINCRONIA
e NÃO é um Armazenador PoE válido.



