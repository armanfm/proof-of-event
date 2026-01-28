# Proof of Event (PoE) — Modelo de Responsabilidade Institucional

## 1. Premissa Fundamental

O Proof of Event é uma **infraestrutura técnica neutra**, não um prestador de serviços de validação, certificação ou custódia.

## 2. Papéis e Responsabilidades (Normativo)

### 2.1 Cliente Final (Consumidor de Prova)
- **Responsabilidade:** escolher e contratar um **Operador** idôneo (verificador e/ou armazenador).
- **Direitos:** receber prova criptográfica do registro.
- **Limites:** não precisa compreender o funcionamento interno do PoE para verificar a prova.

### 2.2 Verificador / Oráculo (Produtor do Evento)
- **Responsabilidade primária por:**  
  - validação semântica do evento (**off-chain**);  
  - geração de `payload_hash` e metadados do evento;  
  - escolha de algoritmos criptográficos usados fora do ledger (clássicos ou PQC);  
  - conformidade legal do procedimento e do conteúdo validado (quando aplicável);  
  - relacionamento contratual com o Cliente Final (se houver).
- **Direitos:** cobrar por serviço de verificação **fora do PoE**.
- **Normativo:** o Verificador **não recebe POE on-chain** pelo protocolo.

### 2.3 Armazenador (Storage Node / Replicador do Ledger)
- **Responsabilidade primária por:**  
  - manter infraestrutura de replicação do ledger;  
  - executar append do ledger conforme ordem canônica;  
  - emitir `Commitment_Proof` (quando aplicável).
- **Direitos:** receber remuneração **exclusivamente** pelo trabalho de armazenamento, conforme TOKENOMICS v0.1.
- **Observação:** custódia de blobs/dados originais é **opcional e contratual**, não inerente ao PoE.

### 2.4 Plataforma PoE (Infraestrutura Técnica)
- **Responsabilidade técnica por:**  
  - manter o FIFO soberano e ordenado;  
  - preservar o ledger append-only;  
  - garantir reexecução determinística do encadeamento e regras do protocolo.
- **Não assume responsabilidade por:**  
  - veracidade, significado, legalidade ou utilidade do conteúdo;  
  - custódia de dados originais;  
  - seleção de criptografia usada fora do ledger;  
  - continuidade operacional de operadores específicos.
- **Normativo:** a Plataforma **não recebe POE on-chain** por padrão.


- **Responsabilidade NÃO assume:**
  - Validação de conteúdo ou significado
  - Custódia de dados originais
  - Conformidade legal
  - Relacionamento com cliente final
  - Decisões sobre algoritmos criptográficos

## 3. Analogias Jurídicas

O PoE é análogo a:
- **Cartório de Registro de Imóveis** (mas para hashes)
- **Sistema de Protocolo Digital** (como e-SAJ, e-CAC)
- **Infraestrutura Pública** (como DNS, NTP)

**NÃO é análogo a:**
- Certificadora Digital (não emite certificados)
- Provedor de Armazenamento (não guarda dados originais)
- Serviço de Validação (não valida conteúdo)

## 4. Fluxo de Responsabilidade

Cliente tem documento → Contrata Armazenador → Armazenador valida
↓ ↓
(Responsabilidade (Responsabilidade
contratual) técnica total)
↓ ↓
Armazenador → FIFO PoE → Ledger Público
↓
(Prova entregue ao cliente)


## 5. Declaração de Limites (Normativo)

> "O Proof of Event garante que um hash foi registrado em uma ordem temporal específica.
> Não garante, valida ou assume responsabilidade sobre o significado, veracidade,
> legalidade ou utilidade do conteúdo representado pelo hash.
> A responsabilidade pela verificação do conteúdo é inteiramente do Armazenador/Verificador
> e de eventuais acordos contratuais entre este e o Cliente Final."

## 6. Modelo Comercial Típico

### Contrato Cliente ↔ Armazenador:

SERVIÇOS DO ARMAZENADOR:

Validação documental (off-chain)

Assinatura digital (algoritmo à escolha)

Submissão ao PoE

Custódia de documentos originais (opcional)

Emissão de relatórios de auditoria

PREÇO: R$ X por documento/evento

SERVIÇOS DO POE (via Armazenador):

Registro imutável no ledger público

Prova criptográfica de timestamp

Ordenação canônica garantida

CUSTO: Incluído no preço do Armazenador


## 7. Conclusão

Este modelo:
- ✅ Protege legalmente a Plataforma PoE
- ✅ Define responsabilidades claras
- ✅ É institucionalmente aceitável
- ✅ Permite ecossistema comercial
- ✅ Alinha com regulação existente

**A blockchain não decide. Ela testemunha.**
**O Armazenador valida. O PoE registra.**


### Não-Exclusividade (Normativo)

Nenhum Cliente Final é obrigado a utilizar um Armazenador específico,
nem a Plataforma PoE impõe exclusividade, dependência técnica ou
vínculo contratual obrigatório entre participantes.

Qualquer Cliente pode migrar entre Armazenadores, desde que possua
os dados necessários para verificação independente.


### Portabilidade de Prova

A prova gerada pelo PoE é pública e verificável de forma independente.
O Cliente Final pode verificar a prova sem depender do Armazenador
original, desde que possua o conteúdo correspondente ao hash registrado.

### Neutralidade Criptográfica

O Proof of Event não impõe, recomenda ou restringe algoritmos
criptográficos utilizados fora do ledger.

A escolha de esquemas clássicos ou pós-quânticos é responsabilidade
exclusiva do Armazenador/Verificador.

## Continuidade Operacional e Limite de Responsabilidade

Embora o Proof of Event exija que Armazenadores mantenham
infraestrutura ativa e ledger completo enquanto operam como nós PoE,
a Plataforma **não garante a continuidade individual** de nenhum
Armazenador ou Verificador específico.

É responsabilidade do Cliente Final:

- escolher um Armazenador/Verificador idôneo;
- avaliar sua capacidade técnica, jurídica e operacional;
- manter, quando necessário, cópias próprias dos dados originais.

O Proof of Event **não é responsável** por falhas, interrupções,
encerramento de atividades ou desaparecimento de Armazenadores,
Verificadores ou operadores individuais.

Enquanto houver **Armazenadores PoE ativos**, o ledger público e
imutável continuará disponível, auditável e reexecutável,
independentemente da permanência de prestadores específicos.


## 2. Papéis e Responsabilidades (Normativo)

### 2.1 Cliente Final (Consumidor de Prova)
- **Responsabilidade:** escolher e contratar um **Operador** idôneo (verificador e/ou armazenador).
- **Direitos:** receber prova criptográfica do registro.
- **Limites:** não precisa compreender o funcionamento interno do PoE para verificar a prova.

### 2.2 Verificador / Oráculo (Produtor do Evento)
- **Responsabilidade primária por:**  
  - validação semântica do evento (**off-chain**);  
  - geração de `payload_hash` e metadados do evento;  
  - escolha de algoritmos criptográficos usados fora do ledger (clássicos ou PQC);  
  - conformidade legal do procedimento e do conteúdo validado (quando aplicável);  
  - relacionamento contratual com o Cliente Final (se houver).
- **Direitos:** cobrar por serviço de verificação **fora do PoE**.
- **Normativo:** o Verificador **não recebe POE on-chain** pelo protocolo.

### 2.3 Armazenador (Storage Node / Replicador do Ledger)
- **Responsabilidade primária por:**  
  - manter infraestrutura de replicação do ledger;  
  - executar append do ledger conforme ordem canônica;  
  - emitir `Commitment_Proof` (quando aplicável).
- **Direitos:** receber remuneração **exclusivamente** pelo trabalho de armazenamento, conforme TOKENOMICS v0.1.
- **Observação:** custódia de blobs/dados originais é **opcional e contratual**, não inerente ao PoE.

### 2.4 Plataforma PoE (Infraestrutura Técnica)
- **Responsabilidade técnica por:**  
  - manter o FIFO soberano e ordenado;  
  - preservar o ledger append-only;  
  - garantir reexecução determinística do encadeamento e regras do protocolo.
- **Não assume responsabilidade por:**  
  - veracidade, significado, legalidade ou utilidade do conteúdo;  
  - custódia de dados originais;  
  - seleção de criptografia usada fora do ledger;  
  - continuidade operacional de operadores específicos.
- **Normativo:** a Plataforma **não recebe POE on-chain** por padrão.

