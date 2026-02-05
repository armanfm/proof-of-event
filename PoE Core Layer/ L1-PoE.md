# Consenso, Ordem Temporal, Identidade Canônica e TimeSwap no Proof of Event (PoE)

## 0. Resumo executivo
O **Proof of Event (PoE)** é um protocolo de registro determinístico, append-only, que fixa eventos no tempo por encadeamento criptográfico.  
Ele **não implementa consenso interno**; em vez disso, fornece uma camada neutra de prova para ser consumida por consensos externos (jurídicos, institucionais, econômicos ou de outras blockchains).

No PoE:
- o passado é preservado por hash chain;
- a ordem é definida por tempo canônico selado;
- a identidade canônica do evento emerge do **TimeSwap**;
- a difusão pública de hashes fortalece a irreversibilidade histórica.

---

## 1. Princípio fundamental
O PoE parte do pressuposto de que **consenso é requisito do mundo real**, não necessariamente do protocolo de registro.  
Seu objetivo é:

1. registrar eventos de forma determinística;
2. impedir reordenação e retroinserção sem detecção;
3. permitir verificação independente/offline.

O papel do PoE é **provar existência e sequência**, não decidir validade social, jurídica ou semântica.

---

## 2. Diferença estrutural em relação ao PoH e blockchain tradicional

### 2.1 PoE vs PoH
O PoE não substitui consenso, não acelera consenso e não depende de votação entre nós.

| Aspecto | PoH | PoE |
|---|---|---|
| Função principal | Ordenar eventos para consenso | Fixar prova factual |
| Relação com consenso | Pré-consenso | Consenso externo |
| Tipo de tempo | Relógio criptográfico interno | Tempo canônico selado |
| Disputa de estado | Existe | Não existe no núcleo |

No PoE, a ordem temporal **não prepara consenso**; ela **encerra o evento**.

### 2.2 PoE vs blockchain tradicional
Blockchains tradicionais combinam:
- ordenação,
- consenso,
- execução/governança.

O PoE separa responsabilidades e mantém foco em:
- prova temporal determinística,
- custo menor,
- neutralidade da camada de registro.

---

## 3. Modelo de dados e semântica dos campos

Cada evento contém, no mínimo:

- `timestamp_canônico`
- `payload`
- `previous_hash`
- `current_hash = H(timestamp_canônico || payload || previous_hash)`
- `event_id` (**opcional**, identificador lógico de aplicação; **não canônico**)

> `event_id` não define identidade canônica no PoE.  
> Ele serve para integração, rastreamento, idempotência e referência de aplicação.  
> A identidade canônica emerge de tempo + contexto histórico + encadeamento criptográfico.

---

## 4. Tempo canônico selado
Tempo canônico selado significa:

- timestamp capturado no registro do evento;
- formato determinístico (ex.: RFC3339 / Unix);
- timestamp incluído no hash do próprio evento;
- vínculo obrigatório com `previous_hash`.

Consequências:
- eventos não podem ser reordenados sem quebrar a cadeia;
- retroinserções tornam-se detectáveis;
- alteração de conteúdo/tempo gera novo hash incompatível.

---

## 5. Identidade canônica derivada do tempo (TimeSwap)

No PoE, a identidade canônica não é endereço, chave, nó ou autoridade.  
Ela emerge da função **TimeSwap**.

### 5.1 Definição
Para um evento `E`:

TimeSwap(E) = H(T_c || P || H_prev)


#### Identidade Intrínseca
O **TimeSwap** é uma propriedade emergente do evento específico; **não é um token transferível**.  
Se **T_c**, **P** ou **H_prev** mudam, a identidade muda.

#### Irrepetibilidade Temporal
Cada **TimeSwap** corresponde a um **momento único** na história da cadeia.

---

## 6. Diagrama semântico do evento

┌─────────────────────────────────────────────────────┐
│ EVENTO E │
│ │
│ TIMESTAMP (T_c) ─────┐ │
│ "Quando" │ │
│ │ │
│ PAYLOAD (P) ─────────┼─▶ H(T_c || P || H_prev) │
│ "O quê" │ │ │
│ │ ▼ │
│ PREVIOUS HASH ───────┘ TimeSwap(E) │
│ "Onde (na história)" │ │
│ │ │
│ event_id (opcional) │ Identidade Canônica │
│ "Para aplicação" │ (Imutável, Única) │
└────────────────────────────┴────────────────────────┘


---

## 7. Modelo de ameaça

O PoE **não assume disputa por maioria** sobre estado global.  
A ameaça principal é a **adulteração retroativa de registros**.

### Mitigações

- hash chain *append-only*;
- tempo canônico selado no hash;
- verificação independente e *offline*;
- difusão externa de hashes (*testemunhas públicas independentes*).

Ataques do tipo **“51%”** só são aplicáveis a sistemas com votação ou consenso por maioria, o que **não é o mecanismo do núcleo PoE**.

---

## 8. Possíveis objeções e respostas

### 8.1 “Sem consenso interno, como evita 51%?”
O PoE não é um protocolo de votação.  
A segurança decorre de **integridade criptográfica + continuidade histórica + difusão de prova**, não de maioria.

### 8.2 “O que impede registrar evento falso?”
Nada no núcleo — **por design**.  
O PoE prova: *“alguém alegou X em T, na posição Y da sequência”*.  
Validade semântica ou jurídica é responsabilidade dos **consumidores externos** (assinaturas, credenciais, auditoria).

### 8.3 “Por que não usar blockchain tradicional?”
Porque o problema aqui é **prova temporal neutra**, não governança ou execução compartilhada.  
Separar camadas reduz custo e aumenta interoperabilidade.

### 8.4 “E se o relógio externo for manipulado?”
Utilizam-se múltiplas fontes de tempo e políticas explícitas de aceitação/auditoria.  
Uma vez selado e difundido, qualquer tentativa de reescrita torna-se detectável.

---

## 9. Difusão do hash e irreversibilidade histórica

Após registrar um evento, seu hash deve ser difundido por **múltiplas vias públicas e descentralizadas** (ex.: ancoragens, registros públicos verificáveis, canais independentes).

Isso:

- **não cria consenso**;
- **não exige P2P**;
- **não é votação**.

Trata-se de **fortalecimento da prova por testemunho distribuído**:

- cada via atesta: *“este hash existia nesse momento”*;
- quanto mais vias independentes, maior o custo de negação histórica;
- adulterar o ledger depois cria uma linha alternativa incompatível com testemunhos já publicados.

---

## 10. Onde está o consenso

O consenso ocorre **sobre as provas fornecidas pelo PoE**, em camadas externas, como:

- blockchains públicas;
- ledgers institucionais;
- DAOs ou consórcios;
- auditorias independentes;
- sistemas jurídicos ou regulatórios.

Em síntese:

- **PoE** = núcleo de prova  
- **Mundo externo** = reconhecimento e decisão  

---

## 11. Casos de uso concretos com TimeSwap

### Certificação de documentos
A identidade do certificado é vinculada ao **instante de emissão** e à cadeia histórica.

### IoT / Telemetria
Cada leitura possui identidade temporal única e sequência auditável.

### Registros institucionais
Atas, logs de auditoria e trilhas de decisão com prova temporal verificável.

### Votação digital (camada de registro)
O PoE fixa **quando** o voto foi registrado; apuração e legitimidade permanecem externas.

---

## 12. Comparação com outros modelos de identidade

| Sistema              | Base da identidade      | Transferível | Dependente de autoridade |
|----------------------|-------------------------|--------------|--------------------------|
| PKI                  | Chave pública           | Sim          | Sim (CA)                 |
| Endereço blockchain  | Hash de chave           | Sim          | Não                      |
| UUID                 | Aleatoriedade           | Sim          | Não                      |
| **PoE TimeSwap**     | T_c + P + H_prev        | Não          | Não (no núcleo)          |

---

## 13. Diretrizes de implementação

### 13.1 Política de aceitação de `timestamp_canônico` (crítica)

#### Política recomendada

**Monotonicidade local da cadeia**

T_c(E_i) ≥ T_c(E_{i-1})


**Janela anti-futuro**
- rejeitar ou quarentenar eventos com  
  `T_c > now_trusted + Δ_future_max`

**Desvio máximo de fontes confiáveis**
- aceitar apenas se  
  `|T_c - now_trusted| ≤ Δ_skew_max`  
- exceção: políticas explícitas de *backfill* ou recuperação

**Registro da fonte temporal (opcional)**
- fonte de tempo utilizada;
- instante de recepção local;
- política aplicada (*normal*, *backfill*, *recuperação*).

**Tratamento de eventos temporalmente inválidos**
- não inserir na cadeia principal;
- armazenar em fila de quarentena para auditoria.

---

### 13.2 Serialização canônica do payload

- formato único e determinístico;
- ordem fixa de campos;
- normalização textual, timezone e precisão;
- nenhuma serialização ambígua entre implementações.

---

### 13.3 Verificação de encadeamento

- `previous_hash` deve apontar exatamente para o hash do último evento válido;
- divergência implica rejeição ou abertura explícita de *fork*.

---

## 14. Protocolo mínimo de verificação (padronizável)

Para cada evento **E_i**:

1. validar *schema* e serialização canônica do payload;  
2. validar política temporal (`T_c`) da Seção 13.1;  
3. verificar `previous_hash == current_hash(E_{i-1})`;  
4. recomputar `current_hash`;  
5. comparar hash recomputado com hash registrado;  
6. registrar resultado:  
   `VALID`, `INVALID`, `QUARANTINE`, `FORK_DETECTED`.

Resultado: **qualquer implementação compatível obtém a mesma conclusão para a mesma cadeia**.

---

## 15. Recuperação e detecção de forks

Como o PoE não possui consenso interno, *forks* são possíveis, porém **detectáveis**.

### 15.1 Fork malicioso (reescrita de história)
- gera cadeias com hashes conflitantes;
- detectável por comparação com hashes já difundidos;
- ramo conflitante perde credibilidade ao divergir de testemunhos públicos.

### 15.2 Fork acidental (instâncias independentes)
- pode ocorrer por execução paralela;
- pode coexistir com *namespaces* distintos;
- geralmente identificado por *genesis* diferente ou divergência precoce.

### 15.3 Recuperação operacional
- consumidores definem política de aceitação;
- critério recomendado: cadeia com **maior lastro de testemunhos públicos verificáveis**;
- manter trilha de auditoria das decisões.

---

## 16. Performance e otimização

### Hashing incremental
Reduz recomputações em verificações contínuas.

### Compactação de payload e snapshots auditáveis
Preserva verificabilidade com menor custo de armazenamento.

### Indexação por `event_id` (opcional) + faixa temporal
Acelera consultas sem alterar identidade canônica.

### Batching de difusão externa de hashes
Otimiza custo mantendo rastreabilidade de ancoragem.

### SDKs multi-linguagem
Padroniza serialização e verificação entre implementações.

---

## 17. Conclusão

O **Proof of Event** utiliza **tempo canônico selado + encadeamento criptográfico + identidade temporal (TimeSwap)** para registrar eventos de forma determinística e auditável.

Ele não substitui o consenso do mundo; **ele o habilita** com provas neutras e interoperáveis.

> **No PoE, identidade não é posse. É tempo.**  
> **Neutralidade da prova é mais importante do que decisão por maioria.**


