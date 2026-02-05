# Economia Operacional v0.1 — Proof of Event (PoE)
## Contabilidade Determinística de Consumo (Fuel)

**Versão:** 0.1  
**Status:** Normativo (regras mecânicas executáveis)  
**Compatibilidade:** SPEC.md v0.1 + protocol/v0.1.md  
**Autor:** Armando Freire  
**Licença:** Apache License 2.0  

---

## 0. Princípio

O Proof of Event (PoE) não decide conteúdo, não interpreta eventos e não resolve disputas.

O PoE testemunha eventos por meio de:
- timestamp canônico;
- prova criptográfica determinística;
- ledger append-only.

A economia do PoE **não é tokenizada**.

Ela é uma **contabilidade interna determinística de consumo**, usada exclusivamente para:
- custear a operação do Certificador;
- limitar abuso;
- garantir custo real por evento certificado.

---

## 1. Unidade Econômica: Fuel (Crédito de Uso)

O PoE utiliza uma unidade interna chamada **Fuel**.

O Fuel:
- **não é token on-chain**;
- **não é transferível**;
- **não possui decimais**;
- **não é especulável**;
- **não representa ativo financeiro**;
- **não participa da prova criptográfica**.

O Fuel existe exclusivamente para **autorizar a certificação de eventos**.

---

## 2. Regra Central de Consumo

### 2.1 Regra Única


1 evento aceito = 1 Fuel consumido


Não existem:
- descontos;
- bônus;
- prioridades;
- múltiplas taxas;
- consumo variável.

---

## 3. Origem do Fuel

O Fuel é creditado **exclusivamente** por:

- pagamentos externos confirmados;
- registrados de forma idempotente;
- auditados em ledger administrativo.

Formas típicas de pagamento:
- moeda fiduciária;
- meios de pagamento tradicionais;
- acordos institucionais.

Pagamentos:
- ocorrem fora do protocolo PoE;
- não criam tokens;
- não concedem direitos econômicos;
- não afetam a prova criptográfica.

---

## 4. Atomicidade do Consumo (Normativo)

O consumo de Fuel ocorre **atomicamente** com o registro do evento.

Regra fundamental:

> Ou o evento foi registrado **e** o Fuel foi consumido,  
> ou o evento não foi registrado **e** nenhum Fuel foi consumido.

Não existe:
- cobrança antecipada;
- cobrança posterior;
- estado intermediário.

---

## 5. Falhas que NÃO Consomem Fuel

Exemplos de falhas antes do registro:

- `BAD_EVENT`
- `INVALID_EVENT_ID`
- `INVALID_HASH_512`
- `NO_COMMUNITY_FUEL`
- `DUPLICATE_EVENT`
- `RATE_LIMIT`
- `CERTIFIER_OFFLINE`

Nestes casos:
- ❌ não há append no ledger
- ❌ não há consumo de Fuel

---

## 6. Registro Bem-Sucedido

Um evento é considerado **registrado** quando:

- formato validado;
- requisitos operacionais atendidos;
- timestamp canônico atribuído;
- prova criptográfica calculada;
- append confirmado no ledger.

Somente após isso:
- **1 Fuel é consumido**.

---

## 7. Congestionamento e Prioridade

O PoE **não vende prioridade**.

Normativo:
- congestionamento não altera ordem;
- congestionamento não compra precedência;
- congestionamento não consome Fuel adicional.

Rate limit e filas são **medidas operacionais**, não parte da economia do protocolo.

---

## 8. Auditoria Econômica

O Certificador mantém:
- estado interno de Fuel;
- ledger administrativo append-only de pagamentos;
- relatório reexecutável de consumo.

A auditoria:
- é determinística;
- é verificável;
- não depende de blockchain.

---

## 9. Separação entre Pagamento e Prova

Pagamentos:
- não entram no `poe_hash`;
- não alteram o ledger de eventos;
- não criam direitos sobre eventos passados.

A prova PoE permanece:
- semanticamente neutra;
- economicamente independente;
- reexecutável sem contexto financeiro.

---

## 10. Limitações Explícitas

A economia do PoE **não promete**:
- valorização;
- governança;
- rendimento;
- proteção contra uso indevido;
- garantias econômicas.

Ela existe apenas para **viabilizar operação sustentável**.

---

## 11. Encerramento

No PoE v0.1:

- quem registra evento consome 1 Fuel;
- quem não registra, não consome;
- o custo é explícito;
- o modelo é simples e auditável.

Se está no ledger, foi consumido.  
Se não está no ledger, não foi consumido.

Sem token.  
Sem governança.  
Sem promessa.

**A blockchain não decide. Ela testemunha.**






