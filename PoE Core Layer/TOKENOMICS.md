# Tokenomics do Proof of Event (PoE)
## Modelo Econômico Baseado em Testemunho Temporal

---

## 1. Objetivo da Tokenomics no PoE

A tokenomics do **Proof of Event (PoE)** não tem como objetivo:

- eleger vencedores probabilísticos;
- criar segurança por desperdício energético;
- substituir consenso social por computação.

Seu objetivo é **remunerar serviços reais de testemunho temporal**, garantindo:

- disponibilidade contínua do sistema;
- integridade da sequência histórica;
- observação correta do tempo;
- difusão pública da prova.

No PoE, o token **não compra poder** — ele remunera **prestação de serviço correto**.

---

## 2. Diferença Estrutural em Relação ao Proof of Work (PoW)

### 2.1 Proof of Work (PoW)

No **PoW**:

- múltiplos agentes competem computacionalmente;
- apenas um vence;
- todos os outros desperdiçam energia;
- a dificuldade ajusta artificialmente a escassez;
- o “tempo” emerge da corrida de hash.

> **No PoW, o custo vem da competição.  
> A segurança nasce do desperdício.**

---

### 2.2 Proof of Event (PoE)

No **PoE**:

- o tempo é observado, não criado;
- a ordem é selada por tempo canônico;
- eventos são aceitos por regras determinísticas;
- não há votação, maioria ou dificuldade;
- a identidade emerge do **TimeSwap**.

> **No PoE, o custo vem da observação do tempo.  
> A segurança nasce da história imutável.**

---

## 3. A Corrida no PoE (Natureza Distinta)

Existe no PoE uma **corrida logística de tempo**, mas ela não é computacional.

| Aspecto | PoW | PoE |
|------|----|----|
| Tipo de corrida | Computacional | Operacional / logística |
| Base da vitória | Hashpower | Registro correto |
| Custo marginal | Crescente | Constante |
| Desperdício | Alto | Mínimo |
| Criação do tempo | Pela corrida | Externa |
| Incentivo | Gastar mais | Respeitar o tempo |

No PoE, **ninguém ganha por gastar mais energia**.  
Ganha quem **registra corretamente dentro da janela válida**.

---

## 4. Modelo de Recompensa com Janela Top-N

Para evitar spam, desperdício e centralização, o PoE adota um modelo **Top-N de elegibilidade econômica**.

- Para cada evento válido, apenas os **N primeiros registros válidos** dentro da janela temporal recebem recompensa.
- `N` é um parâmetro econômico (valor base: **50**).
- Após atingir `N`, novos registros permanecem válidos como prova, mas **não geram tokens**.

O Top-N **não define validade**, apenas **direito econômico**.

---

## 4.2 Determinação do Top-N

A elegibilidade é determinada de forma **determinística**, sem consenso interno.

### Processo

1. Evento `E` ocorre em tempo canônico `T`.
2. Define-se a janela de elegibilidade:

[T, T + Δt_max]

3. Registros válidos são ordenados por:
- validade do timestamp canônico;
- momento de recepção segundo o relógio de referência.
4. Os **primeiros N registros válidos** tornam-se elegíveis à recompensa.

### Relógio de Referência Canônico

Pode ser implementado por:
- múltiplas fontes NTS confiáveis;
- blockchain pública (timestamp de bloco);
- consórcio de servidores de tempo independentes.

O relógio **não decide validade semântica**, apenas ordenação auditável.

---

## 4.3 Rotatividade e Inclusividade

Para evitar captura recorrente do Top-N:

- limite de recompensa por participante por período;
- prioridade para novos participantes válidos;
- reputação baseada em:
- precisão temporal;
- ausência de eventos inválidos;
- disponibilidade contínua.

Essas regras **afetam apenas a economia**, não a prova.

---

## 4.4 Após o Limite N

Após o Top-N ser preenchido:

1. registros continuam válidos como prova;
2. não recebem tokens;
3. fortalecem a irreversibilidade histórica;
4. podem ser usados para auditoria e compliance.

---

## 4.5 Governança do Parâmetro N

O parâmetro `N` é ajustado externamente (Layer 2), com base em:

- número de participantes ativos;
- latência média observada;
- objetivo de descentralização;
- custo econômico por evento.

O **PoE Core (L1)** permanece neutro e determinístico.

---

## 5. Serviços Recompensados

1. **Registro válido**
- validação do timestamp;
- verificação do `previous_hash`;
- cálculo correto do TimeSwap.

2. **Testemunho temporal**
- manutenção de fontes de tempo confiáveis;
- mitigação de timestamps inválidos.

3. **Difusão da prova**
- publicação de hashes;
- redução do custo de negação histórica.

---

## 6. Modelo de Emissão

- Emissão **linear ao uso real**.
- 1 evento válido → emissão de `E` tokens.
- Emissão limitada pela janela Top-N.

Não existe:
- jackpot;
- dificuldade variável;
- inflação por corrida.

---

## 6.1 Parâmetros Econômicos de Referência

| Parâmetro | Valor Base | Variação Permitida |
|-----------|------------|-------------------|
| N (Top-N) | 50 | 10–200 (L2) |
| Δt_max | 10 segundos | 1–60 segundos |
| E (emissão/evento) | 100 tokens | Ajustável (L2) |
| Recompensa base | 50 tokens | — |
| Janela de reputação | 30 dias | — |
| Stake mínimo (opcional) | 1.000 tokens | — |

---

## 7. Distribuição da Recompensa

Distribuição base por evento:

- Registradores Top-N: 50%
- Operadores de tempo: 30%
- Difusores / ancoragens: 20%

---

## 7.1 Modelo Híbrido de Distribuição (Opcional)

A recompensa pode ser suavizada:

R(i) = R_base × f(posição_i) × f(precisão_i)



Exemplo (`f(posição) = 1 / √posição`):

| Posição | Recompensa |
|------|-----------|
| 1º | 100% |
| 10º | 31.6% |
| 50º | 14.1% |

---

## 8. Eficiência Energética

| Sistema | Energia por evento |
|------|------------------|
| PoW | Alta, desperdiçada |
| PoS | Média |
| **PoE** | Baixa, constante |

---

## 9. Segurança Econômica e Mitigações

### Ataque Sybil
Mitigações:
- custo mínimo por participação;
- limite por endereço/IP;
- identidade mínima para grandes volumes;
- reputação histórica.

### Corrida por Latência Extrema
Mitigações:
- `Δt_max` baseado em percentis de latência (P95);
- múltiplas fontes de tempo;
- bônus por diversidade geográfica.

### Spam de Eventos
Mitigações:
- custo mínimo de registro;
- validação semântica em L2/L3;
- reputação negativa para eventos inválidos.

---

## 10. Staking para Qualidade (Opcional)

Stake serve para:
- sinalizar compromisso operacional;
- receber bônus;
- sofrer slashing por fraude comprovada.

Stake **não confere poder político**.

---

## 11. Tokenomics em Camadas

### Layer 1 – PoE Core
- prova temporal;
- recompensa por registro válido;
- neutralidade absoluta.

### Layer 2 – Aplicações
- taxas específicas;
- governança econômica;
- ajuste de parâmetros.

### Layer 3 – Consenso Externo
- decisão semântica;
- auditoria;
- reconhecimento jurídico.

---

## 12. Conformidade Regulatória

### AML / KYC
- identificação exigida para grandes recompensas;
- trilha auditável completa.

### Tributação
- tokens como receita por serviço;
- relatórios automáticos.

### ESG
- consumo energético baixo e mensurável;
- incentivo à descentralização;
- governança transparente.

---

## 13. Fases de Adoção

### Fase 1 — Bootstrap (0–6 meses)
- N elevado;
- emissão maior;
- stake opcional.

### Fase 2 — Crescimento (6–24 meses)
- ajuste de N;
- reputação ativa;
- stake recomendado.

### Fase 3 — Maturidade (>24 meses)
- economia sustentável;
- múltiplas L2;
- integração institucional.

---

## 14. Síntese Final

O Proof of Event inaugura uma nova classe econômica:

- não baseada em computação;
- não baseada em posse;
- não baseada em desperdício.

> **No PoE, ninguém vence o tempo.  
> Quem ganha é quem o respeita.**
