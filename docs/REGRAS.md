# Proof of Event (PoE) — Modelo de Responsabilidade Institucional  
**Versão 0.1 (Corrigida — SEM Armazenadores)**

---

## 1. Premissa Fundamental

O Proof of Event (PoE) é uma infraestrutura técnica neutra para registro
determinístico de eventos externos.

O PoE:

- **NÃO** valida conteúdo  
- **NÃO** interpreta significado  
- **NÃO** julga veracidade  
- **NÃO** executa custódia de dados originais  

O PoE testemunha fatos criptográficos já verificados fora do protocolo.

---

## 2. Papéis e Responsabilidades (Normativo)

### 2.1 Cliente Final (Consumidor de Prova)

**Responsabilidade:**

- Escolher e contratar um Certificador PoE idôneo.

**Direitos:**

- Receber uma prova criptográfica verificável do registro.

**Limites:**

- Não precisa compreender o funcionamento interno do PoE para verificar a prova.

---

### 2.2 Verificador / Oráculo (Produtor do Evento)

Responsabilidade primária por:

- Validação semântica do evento (off-chain);
- Geração do `payload_hash`;
- Definição e uso de algoritmos criptográficos off-chain (clássicos ou pós-quânticos);
- Conformidade legal do procedimento e do conteúdo validado;
- Relacionamento contratual com o Cliente Final (se houver).

**Direitos:**

- Cobrar pelo serviço de verificação fora do PoE.

**Normativo:**

- O Verificador **NÃO** recebe tokens POE on-chain pelo protocolo.

---

### 2.3 Certificador PoE (Executor do Protocolo)

O Certificador PoE é a **única entidade técnica** que:

- recebe eventos já verificados;
- atribui timestamp canônico;
- calcula a `PoE_Proof`;
- executa o append do ledger determinístico;
- mantém o ledger completo e imutável;
- emite recibos verificáveis.

**Responsabilidade técnica total:**

- Integridade do ledger sob sua custódia;
- Determinismo do encadeamento;
- Disponibilidade enquanto operar como Certificador.

**Direitos:**

- Receber remuneração técnica conforme o modelo econômico definido
  (TOKENOMICS v0.1).

**Observação normativa:**

O Certificador **PODE**, mas **NÃO É OBRIGADO**, a:

- armazenar blobs;
- utilizar IPFS, Pinata ou sistemas equivalentes;
- assinar provas com algoritmos pós-quânticos (PQC).

Esses serviços são **opcionais, contratuais e off-chain**.

---

### 2.4 Plataforma PoE (Infraestrutura do Protocolo)

Responsabilidade técnica por:

- Manter a implementação de referência do protocolo PoE;
- Preservar a definição formal das regras de aceitação e do
  **append determinístico do ledger local ao Certificador**;
- Garantir a estabilidade e consistência da especificação pública do protocolo.

A Plataforma **NÃO** assume responsabilidade por:

- Veracidade, significado ou legalidade do conteúdo registrado;
- Custódia de dados originais;
- Algoritmos criptográficos usados fora do ledger PoE;
- Continuidade operacional de Certificadores específicos.

**Normativo:**

- A Plataforma **NÃO** recebe tokens POE on-chain por padrão.

---

## 3. Fluxo de Responsabilidade

Cliente Final
↓ (contrato privado)
Verificador / Oráculo
↓ (evento verificado)
Certificador PoE
↓
Ledger PoE Público e Imutável
↓
Prova entregue ao Cliente



---

## 4. Declaração de Limites (Normativo)

> “O Proof of Event garante exclusivamente que um hash foi registrado por
> um Certificador PoE em um ponto temporal específico.
>
> O protocolo **NÃO** garante, valida ou assume responsabilidade sobre o
> significado, veracidade, legalidade ou utilidade do conteúdo representado
> pelo hash.
>
> Toda responsabilidade pela verificação do conteúdo é externa ao PoE.”

---

## 5. Não-Exclusividade e Portabilidade

- Nenhum Cliente é obrigado a utilizar um Certificador específico;
- Nenhum Certificador possui exclusividade técnica ou institucional;
- A prova PoE é pública e verificável de forma independente;
- O Cliente pode migrar entre Certificadores livremente, desde que possua
  o conteúdo correspondente ao hash registrado.

---

## 6. Neutralidade Criptográfica

O Proof of Event:

- **NÃO** impõe algoritmos criptográficos off-chain;
- **NÃO** restringe o uso de esquemas clássicos ou pós-quânticos;
- **NÃO** depende de assinaturas para validade da prova.

A escolha criptográfica é exclusiva do Verificador e/ou do Certificador,
fora do núcleo do protocolo.

---

## 7. Continuidade Operacional e Limite de Responsabilidade

O PoE **NÃO** garante a continuidade individual de Certificadores específicos.

É responsabilidade do Cliente Final:

- escolher Certificadores idôneos;
- avaliar riscos operacionais;
- manter cópias próprias dos dados originais, quando necessário.

Enquanto houver Certificadores PoE ativos, os ledgers públicos permanecem
auditáveis, verificáveis e reexecutáveis.

---

## 8. Encerramento

O Proof of Event é infraestrutura.

O Certificador executa.  
O Verificador valida.  
O Cliente interpreta.

**A blockchain não decide. Ela testemunha.**
