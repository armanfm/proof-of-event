# TOKENOMICS v0.1 — Proof of Event (PoE)
## Implementação Econômica Normativa  
### Contabilidade Interna Determinística

- **Versão:** 0.1  
- **Status:** Normativo (regras mecânicas executáveis)  
- **Compatível com:** SPEC.md v0.1 + protocol/v0.1.md  
- **Autor:** Armando Freire  
- **Licença:** Apache License 2.0  

---

## 0. Princípio

O Proof of Event (PoE) não decide conteúdo, não interpreta eventos
e não resolve disputas.

O PoE **testemunha** eventos por meio de:
- timestamp canônico;
- prova criptográfica;
- ledger append-only determinístico.

A economia do PoE é **contabilidade interna determinística**:

- o consumo de POE ocorre **exclusivamente por evento registrado**;
- não existe emissão dinâmica;
- o token remunera **trabalho real executado**;
- não há promessa econômica, governança ou privilégio.

---

## 1. Ativo Nativo

- **Nome:** Token PoE  
- **Símbolo:** POE  
- **Decimais:** 18  
- **Governança:** inexistente  

O Token POE é uma **unidade técnica de consumo do protocolo**.

Ele:
- não representa valor mobiliário;
- não confere direitos políticos;
- não promete valorização;
- não participa da prova criptográfica.

---

## 2. Oferta (Supply) e Distribuição

### 2.1 Supply Total

- **GENESIS:** 1.000.000.000.000 POE criados uma única vez
- **Emissão adicional:** proibida permanentemente

Não existe mint por evento, inflação programada ou expansão futura
em v0.1.

---

### 2.2 Distribuição Inicial

- **Reserva Técnica:** 1.000.000.000.000 POE (100%)
- Mantida pela Plataforma PoE como **reserva operacional**
- Utilizada exclusivamente para:
  - provisionar consumo do protocolo;
  - viabilizar operação inicial;
  - garantir disponibilidade de tokens para uso do sistema

Normativo:

A reserva técnica:
- **não implica governança**;
- **não garante valorização**;
- **não confere direitos econômicos** além do consumo do protocolo.

---

### 2.3 Aquisição de Tokens

Clientes podem adquirir POE por:

1. Compra direta da Plataforma (preço fixado em moeda fiduciária por evento)
2. Mercado secundário (preço livre de mercado)
3. Acordos privados entre participantes

Normativo:

- Não há airdrop
- Não há ICO
- Não há venda pública
- Não há distribuição gratuita

O Token POE existe exclusivamente para viabilizar
o uso do protocolo Proof of Event.

---

## 3. Identidades Mecânicas (Sem Identidade Civil)

O PoE utiliza identificadores técnicos para contabilidade automática.

### 3.1 `payer_id` (bytes32)

Identifica quem consome POE:
- cliente institucional;
- verificador;
- oráculo;
- ou a própria Plataforma.

### 3.2 `certifier_id` (bytes32)

Identifica o Certificador PoE responsável por:
- gerar timestamp canônico;
- calcular a PoE_Proof;
- realizar o append no ledger.

Esses IDs:
- não representam identidade civil;
- não implicam confiança;
- não conferem privilégios.

---

## 4. Regra Central — Cobrança por Evento

### 4.1 Taxa Fixa por Evento (v0.1)

- **FEE_EVENT = 1 POE**

Normativo:

- Verificadores **NÃO** recebem POE
- Plataforma **NÃO** recebe POE on-chain
- O POE remunera exclusivamente o serviço
  de certificação e registro do evento

---

## 5. Timing do Consumo (Normativo)

O consumo ocorre **atomicamente** com o append no ledger
do Certificador PoE.

Não existe estado intermediário.

### Regra Fundamental

> Ou o evento foi registrado **e** consumiu,  
> ou o evento não foi registrado **e não** consumiu.

Não há cobrança parcial, antecipada ou posterior.

---

### 5.1 Modelo Lógico de Atomicidade

1. Início da transação lógica  
2. Append do evento no ledger append-only  
3. Débito do saldo correspondente  
4. Commit da transação  

Se qualquer etapa falhar antes do commit:
- o append é abortado;
- nenhum débito ocorre;
- nenhum POE é consumido.

---

## 6. Falhas Antes do Registro (Não Consomem)

### Exemplos de falha ANTES do registro:

1. ❌ ERR_BAD_FORMAT  
2. ❌ ERR_BAD_VERSION  
3. ❌ ERR_NO_TOKEN  
4. ❌ ERR_DUPLICATE_EVENT  
5. ❌ ERR_RATE_LIMIT  
6. ❌ ERR_CERTIFIER_OFFLINE  

Nestes casos:

- ❌ não há append no ledger  
- ❌ não há débito  
- ❌ não há consumo de POE  

---

## 7. Registro Considerado Bem-Sucedido

Um evento é considerado **BEM-SUCEDIDO** quando:

1. ✅ Formato validado  
2. ✅ Pagamento validado (se aplicável)  
3. ✅ Timestamp canônico gerado  
4. ✅ PoE_Proof calculada  
5. ✅ Append confirmado no ledger  

Somente após **TODOS** esses passos ocorre consumo.

---

## 8. Congestionamento (Sem Token)

O congestionamento **não utiliza POE**.

Normativo:

- congestionamento não altera ordem;
- congestionamento não compra prioridade;
- congestionamento não consome nem distribui POE.

Rate limit, fila cheia ou políticas adicionais
são **operacionais** e fora do protocolo PoE.

---

## 9. Limites Técnicos (Hard Limits)

Mesmo pagando, existem limites rígidos:

- MAX_EVENT_JSON_BYTES = 8.192 (8 KiB)
- MAX_BLOB_BYTES = 5.242.880 (5 MiB)

Acima desses limites:
- ERR_PAYLOAD_TOO_LARGE

---

## 10. Auditoria Pública (Obrigatória)

A Plataforma **DEVE** publicar relatório periódico contendo:

- número de eventos aceitos;
- total de POE consumido;
- hash SHA-256 do relatório.

Normativo:

- o hash do relatório **DEVE** ser registrado como evento PoE;
- auditoria é pública, reexecutável e verificável.

---

## 11. Separação entre Pagamento e Consumo

O POE é exclusivamente uma **unidade técnica de consumo**.

Pagamentos:
- ocorrem fora do protocolo;
- podem ser feitos em fiat ou cripto;
- não criam tokens;
- não concedem direitos on-chain.

A Plataforma:
- converte pagamento externo em consumo interno;
- assume risco operacional e cambial;
- não recebe tokens on-chain por padrão.

---

## 12. Avisos (Sem Promessas)

O Token POE:

- não garante valorização;
- não é governança;
- não é dividendo;
- não representa participação societária;
- não opera exchange.

---

## 13. Encerramento

Em v0.1:

- quem usa, paga **1 POE por evento registrado**;
- quem certifica, recebe automaticamente;
- o token remunera trabalho real;
- o protocolo permanece minimalista.

**Se está no ledger, foi cobrado.**  
**Se não está no ledger, não foi cobrado.**

Sem governança.  
Sem promessa.  
Sem privilégio.

**A blockchain não decide. Ela testemunha.**








