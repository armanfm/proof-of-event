# Verificação de Eventos no Proof of Event (PoE)

## 1. Princípio Fundamental

No Proof of Event (PoE), **a verificação de eventos NÃO é executada pelo protocolo**.

O PoE registra fatos criptográficos (hashes) de eventos previamente verificados
fora do sistema. A responsabilidade pela verificação pertence ao **cliente**
ou a ferramentas off-chain por ele utilizadas.

> O PoE **testemunha** eventos.  
> Ele **não valida**, **não interpreta** e **não julga**.

---

## 2. Ledger PoE (Fonte Factual)

O ledger PoE é publicado em formato simples e acessível (ex.: arquivos TXT),
contendo exclusivamente hashes criptográficos em ordem FIFO canônica.

Características do ledger:

- público e acessível;
- imutável (append-only);
- leitura gratuita;
- independente de tokens;
- independente de semântica ou inteligência artificial.

Qualquer pessoa pode baixar o ledger e verificar localmente a existência
de um hash específico.

---

## 3. Verificação pelo Cliente (Off-chain)

A verificação de um evento ocorre **fora do protocolo PoE**, por meio de
ferramentas locais ou serviços delegados pelo cliente.

Essa verificação pode envolver, por exemplo:

- recomputação de hashes;
- validação de documentos;
- checagem de assinaturas criptográficas (incluindo pós-quânticas);
- análise de integridade e consistência;
- comparação ou similaridade entre registros.

O protocolo PoE **não impõe nem executa** esses processos.

---

## 4. mind.bin como Ferramenta de Verificação

O `mind.bin` é um artefato off-chain utilizado para facilitar a verificação,
indexação e análise de eventos registrados no PoE.

O `mind.bin`:

- consome o ledger PoE público (TXT);
- opera localmente ou em ambiente controlado pelo cliente;
- pode armazenar dados completos, metadados e assinaturas;
- não altera o ledger;
- não interfere na ordem FIFO;
- não é requisito para validade do PoE.

O uso do `mind.bin` é **opcional** e específico por aplicação.

---

## 5. Separação de Responsabilidades

| Componente | Responsabilidade |
|-----------|------------------|
| PoE (Ledger/FIFO) | Registro imutável e ordenado de hashes |
| Cliente | Verificação do evento |
| Ferramentas off-chain (ex.: mind.bin) | Auxílio à verificação e análise |

Essa separação garante:

- neutralidade do protocolo;
- auditabilidade independente;
- liberdade tecnológica;
- compatibilidade institucional.

---

## 6. Garantia de Verificabilidade

Mesmo na ausência de qualquer ferramenta adicional:

- o ledger continua válido;
- os hashes continuam verificáveis;
- a prova continua existindo.

Nenhuma dependência de IA, token ou serviço externo é necessária para
verificar um evento registrado no PoE.

---

## 7. Encerramento

No Proof of Event:

**A verdade factual é pública.  
A verificação é livre.  
A inteligência é opcional.**

O protocolo existe para registrar fatos criptográficos,
não para decidir como eles devem ser interpretados.
