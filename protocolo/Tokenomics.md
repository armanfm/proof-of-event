# Proof of Event (PoE)
## Tokenomics — Modelo Econômico do Protocolo

**Versão:** 0.1  
**Status:** Definição Econômica Inicial (Normativa)  
**Compatível com:** SPEC.md v0.1 / protocol v0.1  
**Autor:** Armando José Freire de Melo  

---

## 1. Objetivo do Token PoE

O **Token PoE** é o ativo econômico nativo do protocolo Proof of Event.

Sua função é **liquidar o uso da infraestrutura**, coordenar incentivos operacionais
e permitir a remuneração de participantes que executam trabalho real no sistema.

O token **não** representa:

- participação societária;
- direito a dividendos;
- governança política do protocolo;
- promessa de retorno financeiro.

---

## 2. Natureza do Token

- Tipo: **Ativo digital / criptomoeda**
- Função principal: **token de uso e liquidação operacional**
- Papel econômico: **meio de pagamento do protocolo**
- Governança: **nenhuma** (governança fora do escopo do token)

O valor do token emerge do uso do protocolo e da demanda de mercado, sob risco total
dos participantes.

---

## 3. Oferta (Supply)

### 3.1 Oferta Total

- **Supply fixo e finito**
- Criado **uma única vez** (genesis)
- Não inflacionário
- Não há mint contínuo

> O protocolo **NUNCA** cria novos tokens após o genesis.

---

## 4. Entrada em Circulação

O Token PoE entra em circulação por **trabalho executado** e por aquisição em mercado secundário.

### 4.1 Canais de Distribuição Inicial

- **Armazenadores:** recebem tokens por armazenamento e replicação do ledger.
- **Verificadores / Oráculos:** recebem tokens por verificação off-chain e submissão válida.
- **Plataforma:** recebe tokens por operação do FIFO e manutenção do ecossistema.

Não existe:

- airdrop promocional;
- recompensa por holding;
- yield automático.

---

## 5. Uso do Token (Demanda)

### 5.1 Pagamento do FIFO

Para submeter um evento ao PoE, o verificador/oráculo **DEVE** pagar uma taxa em Token PoE.

Esse pagamento:

- é obrigatório;
- é feito **antes** do evento entrar no FIFO;
- não concede prioridade;
- não altera a ordem.

### 5.2 Quem paga com Token PoE

- verificadores;
- oráculos;
- tokenizadores;
- operadores de sistemas integrados.

> O **cliente final** paga em moeda fiduciária fora do protocolo.

---

## 6. Fluxo Econômico (Loop com Queima Protocolar)

1. O **cliente final** paga o serviço em moeda fiduciária (USD, BRL, EUR).
2. O **verificador/oráculo** executa a verificação do evento fora do PoE (Camada 1).
3. O **verificador** adquire Token PoE por meio de mercado secundário.
4. O **verificador** paga a taxa de uso do FIFO em Token PoE.
5. Ao receber o pagamento, o protocolo aplica a liquidação econômica:
   - **10% dos tokens pagos são queimados permanentemente**;
   - **90% dos tokens são redistribuídos operacionalmente**.
6. Os tokens redistribuídos retornam ao mercado por venda ou reutilização.
7. O ciclo se repete conforme o uso do protocolo.

> A queima é **obrigatória, automática e irrevogável**.  
> Não existe configuração, exceção ou adiamento da queima.

---

## 7. Redistribuição do Token

A redistribuição ocorre como **pagamento por trabalho efetivo**, não por especulação.

### 7.1 Regra de Redistribuição (Normativa)

Dos tokens pagos ao FIFO:

- **10%** são queimados permanentemente;
- **90%** são redistribuídos da seguinte forma:
  - **40%** para Armazenadores (armazenamento e replicação do ledger);
  - **30%** para Verificadores/Oráculos (trabalho off-chain);
  - **20%** para Plataforma (operação do FIFO e infraestrutura).

Esses percentuais são parte integrante do protocolo PoE v0.1 e **NÃO podem ser alterados**
sem versionamento formal do protocolo.

### 7.2 Critérios Operacionais (Implementação)

A implementação define critérios **mecânicos e auditáveis** para cálculo de distribuição,
por exemplo:

- eventos efetivamente armazenados por nó;
- disponibilidade (uptime) medida por janela;
- participação em replicação/sincronização.

A implementação **NÃO** pode introduzir discriminação subjetiva entre participantes
para a mesma classe de trabalho.

---

## 8. Ausência de Incentivos Financeiros Promissórios

O Token PoE:

- não garante valorização;
- não garante liquidez;
- não garante retorno;
- não garante proteção contra perda.

Não há:

- staking;
- slashing;
- lock obrigatório;
- dividendos;
- buyback protocolar.

---

## 9. Ataques Econômicos e Mitigações (Resumo)

### 9.1 Spam

- Spam custa token.
- Uso abusivo torna-se economicamente inviável.

### 9.2 Sybil

- Criar identidades não reduz custo de uso.
- Cada evento exige pagamento em Token PoE.

### 9.3 Centralização

- Não há poder político associado ao token.
- Acumular token **não dá controle do protocolo**.

### 9.4 Acumulação Excessiva (Hoarding)

A **queima protocolar** reduz o supply ao longo do tempo e desencoraja a acumulação
infinita, forçando circulação econômica para uso contínuo do FIFO.

---

## 10. Mercado Secundário (Requisito Operacional)

A existência de **mercado secundário funcional** para o Token PoE é condição necessária
para a operação contínua do protocolo.

O protocolo:

- não opera exchange;
- não garante liquidez;
- mas assume a existência de meios de aquisição do token por participantes.

Sem acesso ao Token PoE, não há submissão ao FIFO.

---

## 11. Parâmetros Fora do Escopo

Este documento **não define**:

- preço do token;
- market making;
- listagem em exchanges;
- estratégias de estabilização de preço.

Esses aspectos pertencem ao mercado, não ao núcleo do protocolo.

---

## 12. Transparência e Auditoria

Todos os fluxos de uso do Token PoE são:

- verificáveis;
- rastreáveis;
- auditáveis externamente.

O protocolo não oculta fluxos econômicos.

---

## 13. Declarações e Avisos de Risco

1. **Não é participação societária**  
   O Token PoE não representa propriedade, controle administrativo ou direitos sobre receitas.

2. **Sem promessa de lucro**  
   Não há promessa, garantia ou expectativa contratual de valorização.

3. **Risco total de mercado**  
   O valor do token pode cair significativamente, inclusive até **zero**.

4. **Responsabilidade fiscal e regulatória**  
   Participantes são responsáveis por obrigações fiscais e conformidade aplicáveis
   em sua jurisdição.

---

## 14. Encerramento

O Token PoE existe para **pagar infraestrutura**, remunerar trabalho operacional e impor
custo econômico real ao uso do FIFO.

**Quem usa, paga.  
Quem trabalha, recebe.  
Quem segura, assume risco de mercado.**
