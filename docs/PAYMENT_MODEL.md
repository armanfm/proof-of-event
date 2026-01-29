# Proof of Event (PoE)
## Modelo de Pagamento e Consumo — v0.1

- **Status:** Normativo  
- **Escopo:** Institucional / Econômico  
- **Aplicável a:** PoE Core v0.1  

---

## 1. Objetivo

Este documento define como o uso do Proof of Event (PoE) é pago e como esse
pagamento é separado do consumo técnico de tokens POE no protocolo v0.1.

Este documento **NÃO** define:

- preços;
- valores comerciais;
- contratos;
- SLAs;
- condições de mercado.

Esses aspectos são externos ao protocolo.

---

## 2. Separação de Responsabilidades Econômicas

### 2.1 Pagamento Off-chain

Os pagamentos pelo uso do PoE ocorrem fora do protocolo.

O protocolo PoE:

- **NÃO** processa pagamentos;
- **NÃO** valida transações financeiras;
- **NÃO** armazena dados financeiros.

Os meios de pagamento podem incluir, por exemplo:

- contratos em moeda fiduciária (fiat);
- faturamento recorrente;
- notas fiscais;
- SLAs;
- sistemas internos de billing.

Esses mecanismos são inteiramente externos ao PoE.

---

### 2.2 Consumo de Token On-chain

O token POE é utilizado exclusivamente como **unidade técnica de consumo**
do protocolo.

O consumo de POE representa:

- uso do **mecanismo determinístico de submissão e registro de eventos
  local ao Certificador**;
- crescimento do ledger determinístico;
- carga de infraestrutura do Certificador.

O consumo técnico de POE:

- **NÃO** é um pagamento;
- **NÃO** é realizado pelo Cliente Final;
- **NÃO** representa contraprestação financeira.

O consumo pode resultar em:

- burn;
- redistribuição técnica;
- alocação interna definida pelo protocolo.

---

## 3. Papel da Plataforma PoE

A Plataforma PoE atua como ponte econômica entre o mundo off-chain e o protocolo.

**Responsabilidades da Plataforma:**

- receber pagamentos off-chain;
- autorizar o uso do PoE;
- contabilizar eventos aceitos;
- consumir tokens POE de forma correspondente.

A Plataforma:

- **PODE** adquirir tokens POE por qualquer meio;
- **NÃO** exige que Clientes finais possuam tokens;
- **NÃO** transfere tokens para Clientes.

Clientes finais **NÃO** interagem com tokens POE.

---

## 4. Natureza Não-Promissória do Token POE

O token POE:

- **NÃO** garante acesso contínuo ao serviço;
- **NÃO** representa participação societária;
- **NÃO** representa direito a receita;
- **NÃO** é resgatável;
- **NÃO** é instrumento financeiro.

O POE é uma **unidade técnica interna do protocolo**.

---

## 5. Auditabilidade

Relatórios econômicos gerados pela Plataforma **PODEM** ter seus hashes
registrados no ledger PoE.

Isso permite:

- auditoria independente;
- verificação de consumo;
- transparência operacional;

sem expor:

- valores financeiros;
- dados contratuais;
- informações sensíveis off-chain.

---

## 6. Versionamento

Qualquer alteração neste modelo:

- **DEVE** ser versionada explicitamente (v0.2+);
- **DEVE** ser documentada publicamente;
- **DEVE** ter seu hash registrado no PoE.

Mudanças silenciosas são proibidas.

---

## 7. Modelos de Pagamento (Informativo)

O protocolo Proof of Event **NÃO** processa pagamentos.

Pagamentos pelo uso do serviço PoE:

- ocorrem fora do protocolo;
- podem ser realizados em moeda fiduciária ou criptoativos;
- autorizam apenas a submissão de eventos via Plataforma.

O token POE:

- **NÃO** é meio de pagamento;
- **NÃO** é exigido do Cliente Final;
- é utilizado exclusivamente como unidade técnica interna de consumo.

---

## 8. Encerramento

No Proof of Event v0.1:

- pagamento é externo;
- consumo é técnico;
- tokens não são dinheiro;
- o protocolo permanece neutro.

O PoE não cobra.  
O PoE não vende.  
O PoE registra.
