## 5. Criptografia e Resiliência Pós-Quântica

O ledger canônico do PoE:

- utiliza exclusivamente **hash criptográfico** (ex.: SHA-256);
- não depende de criptografia assimétrica;
- não contém assinaturas no histórico imutável.

Consequências:

- ataques quânticos não permitem forjar eventos passados;
- a ordem FIFO não depende de chaves privadas;
- a integridade do ledger permanece verificável.

Fluxos de identidade e assinatura (inclusive pós-quânticos) são executados
**off-chain** e podem ser atualizados sem impacto no histórico.

---

## 6. Modelo de Ameaças (Resumo)

### 6.1 Spam e Abuso
Mitigado por custo econômico de submissão (Token PoE).

### 6.2 Nós Maliciosos
Armazenadores não possuem poder de decisão.
Qualquer divergência de hash é detectável.

### 6.3 Sybil
Criar identidades adicionais não reduz custo nem altera ordem.

### 6.4 Falha de Armazenadores
Nós offline perdem remuneração e devem sincronizar ao retornar.

### 6.5 Falha do FIFO
O ledger permanece válido; apenas a entrada de novos eventos é interrompida.

---

## 7. Inteligência Artificial e Ferramentas Off-chain

Ferramentas como `mind.bin` operam **fora do protocolo** e não afetam:

- validade do ledger;
- ordem FIFO;
- aceitação de eventos.

Falhas, vieses ou erros em ferramentas off-chain **não comprometem**
a segurança do PoE.

---

## 8. Divulgação Responsável de Vulnerabilidades

Vulnerabilidades relacionadas ao protocolo PoE podem ser reportadas de forma
responsável por meio de:

- abertura de issue privada (quando disponível); ou
- contato direto com o mantenedor do repositório.

Relatórios devem incluir:

- descrição clara do problema;
- impacto potencial;
- passos mínimos para reprodução.

Não há programa de recompensa (bug bounty) neste estágio.

---

## 9. Aviso de Limitação

O Proof of Event não promete:

- inviolabilidade absoluta;
- resistência total a todos os modelos de ataque;
- proteção contra uso indevido fora do escopo definido.

A segurança do protocolo decorre da **simplicidade, transparência e verificabilidade**,
não de garantias implícitas ou obscuras.

---

## 10. Encerramento

O PoE foi projetado para ser **difícil de enganar**, não impossível de usar mal.

A segurança do sistema reside na clareza de seus limites,
na auditabilidade independente e na separação rigorosa de responsabilidades.
