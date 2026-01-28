# TOKENOMICS v0.1 — Proof of Event (PoE)
Implementação Econômica Normativa (Contabilidade Interna Determinística)

**Versão:** 0.1  
**Status:** Normativo (regras mecânicas executáveis)  
**Compatível com:** SPEC.md v0.1 + protocol/v0.1.md  
**Autor:** Armando Freire  
**Licença:** Apache 2.0  

---

## 0. Princípio

O PoE **não decide conteúdo** e **não interpreta eventos**.  
O PoE **testemunha**: ordem FIFO + ledger append-only.

A economia do PoE é **contabilidade interna determinística**:  
o consumo de POE ocorre por evento aceito, e a distribuição remunera  
**exclusivamente o trabalho de armazenamento do ledger**.

---

## 1. Ativo Nativo

- **Nome:** Token PoE  
- **Símbolo:** POE  
- **Decimais:** 18  
- **Governança:** nenhuma  

O POE é um **token técnico de consumo do protocolo**, não um instrumento financeiro.

---

## 2. Oferta (Supply)

- **Supply total (genesis):** `1.000.000.000.000 POE`  
- **Mint adicional:** proibido em v0.1  

Em v0.1, a emissão é **fechada**:  
não existe criação infinita de tokens por evento.

---

## 3. Identidades Mecânicas (Sem Identidade Civil)

Para contabilidade interna automática, o protocolo utiliza IDs mecânicos:

- **storer_id:** `bytes32`  
  Identifica o nó armazenador que realizou o append canônico do ledger.

- **payer_id:** `bytes32`  
  Identifica quem consome POE para submeter eventos  
  (verificador, cliente institucional ou Plataforma).

Esses IDs **não representam identidade civil**, apenas endereços técnicos  
para débito e crédito contábil.

---

## 4. Regra Central — Cobrança e Distribuição por Evento

### 4.1 Taxa fixa por evento aceito (v0.1)

- **FEE_STORER = 1 POE**


TOTAL_FEE_BASE = 1 POE

## Normativo

- O verificador **NÃO** recebe **POE**.  
- A Plataforma **NÃO** recebe **POE on-chain**.  
- O token **POE remunera exclusivamente** o trabalho de **armazenamento do ledger**.

---

## 4.2 Momento exato da liquidação

A liquidação ocorre **somente quando**:

1. o evento foi aceito pelo **FIFO**  
   *(formato válido + `previous_event_hash` correto + ordem FIFO)*, **e**
2. o evento foi efetivamente gravado no **ledger**  
   *(append confirmado no nó armazenador).*

Se a falha ocorrer **antes disso**:

- não há débito;
- não há crédito;
- não há consumo de POE.

---

## 4.3 Transferência determinística (contabilidade interna)

Ao aceitar e gravar um evento:

- debitar `payer_id` em **1 POE**
- creditar `storer_id` em **1 POE**

Essa operação é **determinística, reexecutável e auditável**.

### Normativo

Se `payer_id` não possuir saldo suficiente, a submissão **DEVE** ser recusada  
**antes do append**, com:

- `ERR_INSUFFICIENT_BALANCE`

---

## 5. Congestionamento (Sem Token)

O congestionamento **não utiliza POE**.

### Normativo

- congestionamento **não altera ordem**
- congestionamento **não compra prioridade**
- congestionamento **não consome nem distribui POE**

Mecanismos de rate limit, fila cheia ou cobrança adicional  
ocorrem **fora do protocolo PoE**, como política operacional da Plataforma.

---

## 6. Limites Técnicos (Hard Limits)

Mesmo pagando, existem limites rígidos para evitar abuso:

- `MAX_EVENT_JSON_BYTES = 8.192` (8 KiB)
- `MAX_BLOB_BYTES = 5.242.880` (5 MiB)

Acima desses limites:

- `ERR_PAYLOAD_TOO_LARGE`

---

## 7. Auditoria Pública (Obrigatória)

A Plataforma **DEVE** publicar relatório auditável periódico contendo:

- número de eventos aceitos
- total de POE debitado
- total de POE creditado aos armazenadores
- hash `SHA-256` do relatório

### Normativo

O hash do relatório **DEVE** ser registrado como evento no PoE  
(`payload_hash` do relatório).

---

## 8. Parâmetros e Mudanças de Regra

Em v0.1:

- `FEE_STORER = 1 POE`
- somente armazenadores recebem POE
- não existe remuneração on-chain para Plataforma ou verificadores

Qualquer mudança:

- exige versionamento (**v0.2+**)
- exige registro do hash do novo arquivo de parâmetros no PoE

Mudanças silenciosas são **proibidas**.

---

## 9. Erros Mínimos

- `ERR_BAD_FORMAT`
- `ERR_BAD_PREV_HASH`
- `ERR_PAYLOAD_TOO_LARGE`
- `ERR_INSUFFICIENT_BALANCE`
- `ERR_INTERNAL`

---

## 10. Avisos (Sem Promessa)

O token POE:

- não garante valorização
- não é governança
- não é dividendo
- não opera exchange
- não representa participação societária

---

## 11. Separação entre Pagamento e Consumo

O POE é exclusivamente uma **unidade técnica de consumo do protocolo**.

Pagamentos pelo uso do PoE:

- ocorrem fora do protocolo
- podem ser feitos em fiat ou cripto
- não criam tokens
- não concedem direitos on-chain

A Plataforma:

- converte pagamentos externos em consumo de POE
- assume risco operacional e cambial
- não recebe tokens on-chain por padrão

---

## 12. Encerramento

Em v0.1:

- quem usa, paga **1 POE** por evento
- quem mantém o ledger, recebe automaticamente
- o token remunera trabalho real
- o PoE permanece minimalista: ordena e testemunha

**Quem usa, paga.**  
**Quem mantém o ledger, recebe.**  
**Sem governança, sem promessa, sem privilégio.**






