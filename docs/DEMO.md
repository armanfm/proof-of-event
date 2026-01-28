# Demonstração de Verificação Off-chain

A imagem abaixo demonstra a verificação de um evento registrado no Proof of Event (PoE)
utilizando exclusivamente:

- o ledger PoE público (hashes em TXT);
- uma ferramenta off-chain de verificação (mind.bin).

O protocolo PoE não executa verificação, não interpreta conteúdo e não participa
do processo de matching.

A verificação ocorre integralmente no ambiente do cliente, garantindo:
- leitura gratuita do ledger;
- independência de infraestrutura central;
- reexecução determinística.

O resultado exibido demonstra:
- correspondência exata (100%) quando o hash existe no ledger;
- scores de similaridade apenas como diagnóstico opcional;
- separação clara entre prova factual e análise semântica.

<img width="1523" height="887" alt="image" src="https://github.com/user-attachments/assets/d6f0bd0f-460d-4b7b-b54e-63db8f61e249" />

---

---

## Observação Técnica — Uso de SHA-512 no PoE

A imagem abaixo ilustra um exemplo de verificação utilizando **hash SHA-512** como identificador criptográfico do evento.

No contexto do Proof of Event (PoE), o uso de SHA-512 **não impacta negativamente a performance do sistema**, pois:

- o hash é calculado **uma única vez por evento**;
- não existe mineração, competição ou repetição de hash;
- não há consenso probabilístico nem prova de trabalho;
- o custo dominante do sistema é **I/O, rede e armazenamento**, não hashing.

O hash criptográfico no PoE tem função **exclusivamente estrutural**:
- encadeamento determinístico do ledger;
- prova de integridade do evento;
- reexecução verificável por qualquer parte.

A escolha por SHA-512 aumenta a margem criptográfica de longo prazo, sem introduzir custo operacional relevante no modelo FIFO + append-only adotado pelo protocolo.

> No PoE, segurança criptográfica não compete com performance, pois o sistema não utiliza hashing competitivo.

A verificação continua sendo:
- **exata** (FOUND / NOT FOUND);
- **determinística**;
- **independente de infraestrutura central**;
- **executável integralmente no ambiente do cliente**.

<img width="1494" height="730" alt="image" src="https://github.com/user-attachments/assets/c86da68d-dadc-4ec4-b483-c40c64bf03d0" />


