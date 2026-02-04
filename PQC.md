# Proof of Event (PoE)
## Criptografia Pós-Quântica (PQC) — Diretriz Técnica v0.1

### Princípio Central

O Proof of Event (PoE) foi projetado para que a **validade da prova criptográfica
não dependa de chaves privadas, consenso distribuído ou coordenação social**.

O protocolo opera sobre três pilares:

- determinismo absoluto;
- ancoragem temporal canônica;
- ledger append-only imutável por encadeamento criptográfico.

A criptografia pós-quântica (PQC) é tratada no PoE como **camada opcional de reforço
institucional**, e **não como requisito para validade da prova**.

---

## 1. Núcleo Criptográfico do PoE

O núcleo do Proof of Event utiliza exclusivamente:

- **SHA-512** como função de hash;
- concatenação determinística de campos;
- encadeamento por hash chain com raiz GENESIS;
- ausência total de chaves privadas no mecanismo central.

A prova PoE é derivada de dados públicos e verificáveis, incluindo:

- `payload_hash_512`
- `timestamp_canônico` (microsegundos UTC em 16 dígitos)
- identificadores técnicos (cliente, verificador opcional)
- hash anterior da cadeia

Essa arquitetura garante que:

- qualquer prova é verificável independentemente;
- não existe ponto único de confiança criptográfica;
- não há dependência de identidade ou assinatura para validade.

---

## 2. Robustez Pós-Quântica do SHA-512

O PoE utiliza SHA-512 por razões técnicas explícitas:

- alta resistência a colisões;
- ausência de dependência de segredo;
- robustez significativa frente a ataques quânticos conhecidos.

Mesmo sob o modelo de Grover, a segurança efetiva permanece elevada,
tornando o SHA-512 adequado para cenários de longo prazo.

No PoE:

- a integridade temporal é garantida por hash;
- a prova não depende de autenticação criptográfica;
- a quebra de qualquer esquema de assinatura não invalida o ledger.

---

## 3. Assinaturas Pós-Quânticas como Camada Opcional

Embora não façam parte do núcleo do protocolo, **RECOMENDA-SE** que
Certificadores PoE utilizem **assinaturas pós-quânticas (PQC)** de forma opcional,
quando aplicável.

Exemplos de algoritmos adequados incluem:

- Dilithium
- Falcon
- outros esquemas padronizados ou aceitos institucionalmente

A assinatura pode ser aplicada sobre:

- o hash final da prova PoE; ou
- o recibo JSON emitido ao cliente.

---

## 4. Propriedades da Assinatura PQC no PoE

A assinatura pós-quântica opcional:

- **NÃO altera** a prova PoE;
- **NÃO interfere** no encadeamento do ledger;
- **NÃO impacta** a eficiência do sistema;
- **NÃO é requisito** para verificação independente;
- **NÃO cria dependência operacional**.

Seu uso tem como objetivos exclusivos:

- autenticar a identidade institucional do Certificador;
- atender exigências regulatórias ou contratuais;
- reforçar a longevidade criptográfica do recibo.

A ausência de assinatura PQC **não reduz a validade da prova**.

---

## 5. Ausência de Consenso e Segurança Criptográfica

O Proof of Event **não utiliza consenso distribuído**.

Não existem:

- múltiplos nós competindo;
- votação;
- mineração;
- staking;
- forks;
- coordenação econômica.

Cada Certificador opera de forma soberana e independente.

Como consequência direta:

- não existem ataques de consenso;
- não existem reorganizações de histórico;
- não há coordenação maliciosa possível no protocolo;
- camadas opcionais (como PQC) não introduzem risco sistêmico.

O processo de certificação é **estritamente determinístico**:
dada a mesma entrada válida e o mesmo estado interno,
o resultado criptográfico é sempre idêntico.

---

## 6. Separação Clara de Responsabilidades Criptográficas

No PoE:

- **SHA-512** garante integridade e imutabilidade temporal;
- **PQC opcional** pode garantir identidade institucional;
- nenhuma camada depende da outra para validade.

Essa separação garante:

- neutralidade do protocolo;
- simplicidade auditável;
- compatibilidade institucional;
- longevidade criptográfica.

---

## 7. Considerações sobre Armazenamento e Eficiência

O uso de assinaturas pós-quânticas:

- ocorre fora do núcleo do ledger;
- não impacta o processo de append;
- não afeta rate limits ou throughput;
- não altera o tamanho ou estrutura do hash chain.

Armazenamento adicional (ex.: recibos assinados) é opcional
e **não interfere na eficiência do sistema PoE**.

---

## 8. Declaração Final

O Proof of Event não depende de promessas criptográficas futuras.

Ele foi projetado para:

- ser verificável com ferramentas mínimas;
- resistir à obsolescência de esquemas de assinatura;
- permanecer auditável mesmo sob mudanças tecnológicas profundas.

A criptografia pós-quântica é bem-vinda no PoE —
**não como obrigação**, mas como **reforço consciente e opcional**.

O protocolo permanece simples, honesto e determinístico.

**A blockchain não decide.  
Ela testemunha.**
