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
