// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/**
 * @title PoEAnchor
 * @notice Ancora + liquida com multisig de validadores.
 *
 * Segurança:
 *   - Lista de validadores on-chain (uma chave por nó)
 *   - Quorum de assinaturas obrigatório antes de executar
 *   - Nenhum nó executa sozinho
 *   - Merkle root + assinaturas = prova imutável
 *
 * Fluxo:
 *   1. Go pré-valida + monta batch
 *   2. N nós assinam o Merkle root
 *   3. Qualquer nó envia batch + assinaturas
 *   4. Contrato verifica quorum de assinaturas
 *   5. Ancora + liquida atomicamente
 */

interface IDREX {
    function transferFrom(address from, address to, uint256 amount)
        external returns (bool);
    function balanceOf(address account)
        external view returns (uint256);
}

contract PoEAnchor {

    // --------------------------------------------------------
    // Eventos
    // --------------------------------------------------------

    event Anchored(
        bytes32 indexed merkleRoot,
        uint256         fromSeq,
        uint256         toSeq,
        uint256         entryCount,
        uint256         timestamp
    );

    event Transferido(
        uint256 indexed seq,
        address indexed de,
        address indexed para,
        uint256         valor,
        bytes32         payloadHash
    );

    event TransferenciaFalhou(
        uint256 indexed seq,
        address de,
        address para,
        uint256 valor,
        string  motivo
    );

    event ValidadorAdicionado(address indexed validador);
    event ValidadorRemovido(address indexed validador);

    // --------------------------------------------------------
    // Estado
    // --------------------------------------------------------

    address public immutable admin;      // só adiciona/remove validadores
    IDREX   public immutable drex;

    // Lista de validadores — um por nó Go
    mapping(address => bool)    public validadores;
    address[]                   public listaValidadores;
    uint256                     public totalValidadores;
    uint256                     public quorumMinimo;

    uint256 public lastAnchoredSeq;
    uint256 public totalBatches;
    uint256 public totalTransferido;

    // Anti-replay de batches já executados
    mapping(bytes32 => bool) public batchExecutado;

    struct AnchorRecord {
        bytes32 merkleRoot;
        uint256 fromSeq;
        uint256 toSeq;
        uint256 entryCount;
        uint256 timestamp;
        uint256 valorLiquidado;
        uint256 transferenciasOk;
        uint256 assinaturasUsadas;
    }

    struct Instrucao {
        uint256 seq;
        address de;
        address para;
        uint256 valor;
        bytes32 payloadHash;
        uint256 deadline;
    }

    mapping(uint256 => AnchorRecord) public anchors;

    // --------------------------------------------------------
    // Constructor
    // --------------------------------------------------------

    constructor(address _drex, uint256 _quorumMinimo) {
        require(_drex != address(0),   "DREX zero");
        require(_quorumMinimo > 0,     "quorum zero");
        admin         = msg.sender;
        drex          = IDREX(_drex);
        quorumMinimo  = _quorumMinimo;
    }

    // --------------------------------------------------------
    // Modifiers
    // --------------------------------------------------------

    modifier apenasAdmin() {
        require(msg.sender == admin, "nao e admin");
        _;
    }

    modifier apenasValidador() {
        require(validadores[msg.sender], "nao e validador");
        _;
    }

    // --------------------------------------------------------
    // Gestão de validadores
    // --------------------------------------------------------

    /**
     * @notice Adiciona um nó validador.
     * Cada nó Go tem sua própria chave.
     */
    function adicionarValidador(address validador)
        external
        apenasAdmin
    {
        require(validador != address(0),    "endereco zero");
        require(!validadores[validador],    "ja e validador");

        validadores[validador] = true;
        listaValidadores.push(validador);
        totalValidadores++;

        emit ValidadorAdicionado(validador);
    }

    /**
     * @notice Remove um nó validador comprometido.
     */
    function removerValidador(address validador)
        external
        apenasAdmin
    {
        require(validadores[validador], "nao e validador");

        validadores[validador] = false;
        totalValidadores--;

        // Remove da lista
        for (uint256 i = 0; i < listaValidadores.length; i++) {
            if (listaValidadores[i] == validador) {
                listaValidadores[i] = listaValidadores[listaValidadores.length - 1];
                listaValidadores.pop();
                break;
            }
        }

        // Ajusta quorum se necessário
        if (quorumMinimo > totalValidadores) {
            quorumMinimo = totalValidadores;
        }

        emit ValidadorRemovido(validador);
    }

    /**
     * @notice Atualiza o quorum mínimo.
     */
    function atualizarQuorum(uint256 novoQuorum)
        external
        apenasAdmin
    {
        require(novoQuorum > 0,                "quorum zero");
        require(novoQuorum <= totalValidadores, "quorum maior que validadores");
        quorumMinimo = novoQuorum;
    }

    // --------------------------------------------------------
    // Verificação de quorum de assinaturas
    // --------------------------------------------------------

    /**
     * @notice Verifica se o Merkle root foi assinado
     *         por validadores suficientes.
     *
     * Cada nó assina: keccak256(merkleRoot + fromSeq + toSeq + entryCount)
     * Contrato recupera o endereço de cada assinatura
     * e verifica se é validador registrado.
     */
    function _verificarQuorum(
        bytes32    merkleRoot,
        uint256    fromSeq,
        uint256    toSeq,
        uint256    entryCount,
        bytes[]    calldata assinaturas
    )
        internal
        view
        returns (uint256 assinaturasValidas)
    {
        // Mensagem que cada nó assinou
        bytes32 mensagem = keccak256(
            abi.encodePacked(
                merkleRoot,
                fromSeq,
                toSeq,
                entryCount
            )
        );

        // Prefixo Ethereum padrão
        bytes32 hash = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                mensagem
            )
        );

        // Verifica cada assinatura
        address[] memory jaContados = new address[](assinaturas.length);
        uint256 contagem = 0;

        for (uint256 i = 0; i < assinaturas.length; i++) {
            address signer = _recuperarSigner(hash, assinaturas[i]);

            // É validador registrado?
            if (!validadores[signer]) continue;

            // Não contar o mesmo validador duas vezes
            bool duplicado = false;
            for (uint256 j = 0; j < contagem; j++) {
                if (jaContados[j] == signer) {
                    duplicado = true;
                    break;
                }
            }
            if (duplicado) continue;

            jaContados[contagem] = signer;
            contagem++;
        }

        return contagem;
    }

    function _recuperarSigner(bytes32 hash, bytes calldata sig)
        internal
        pure
        returns (address)
    {
        require(sig.length == 65, "sig invalida");
        bytes32 r;
        bytes32 s;
        uint8   v;
        assembly {
            r := calldataload(sig.offset)
            s := calldataload(add(sig.offset, 32))
            v := byte(0, calldataload(add(sig.offset, 64)))
        }
        if (v < 27) v += 27;
        return ecrecover(hash, v, r, s);
    }

    // --------------------------------------------------------
    // Core — ancora + liquida com multisig
    // --------------------------------------------------------

    /**
     * @notice Recebe batch pré-validado com quorum de assinaturas.
     *
     * @param merkleRoot   Merkle root calculado pelo Go
     * @param fromSeq      Sequência inicial
     * @param toSeq        Sequência final
     * @param entryCount   Quantidade de eventos
     * @param instrucoes   Transferências pré-validadas
     * @param assinaturas  Uma assinatura por nó validador
     */
    function receberELiquidar(
        bytes32           merkleRoot,
        uint256           fromSeq,
        uint256           toSeq,
        uint256           entryCount,
        Instrucao[]       calldata instrucoes,
        bytes[]           calldata assinaturas
    )
        external
        apenasValidador
    {
        // Anti-replay — esse batch já foi executado?
        bytes32 batchId = keccak256(
            abi.encodePacked(merkleRoot, fromSeq, toSeq)
        );
        require(!batchExecutado[batchId], "batch ja executado");

        // Validações mínimas
        require(merkleRoot != bytes32(0),    "root invalido");
        require(toSeq >= fromSeq,            "intervalo invalido");
        require(entryCount > 0,              "batch vazio");
        require(
            fromSeq == lastAnchoredSeq + 1,
            "gap na sequencia"
        );

        // Verifica quorum de assinaturas
        uint256 assinaturasValidas = _verificarQuorum(
            merkleRoot, fromSeq, toSeq, entryCount, assinaturas
        );
        require(
            assinaturasValidas >= quorumMinimo,
            "quorum insuficiente"
        );

        // Marca como executado — anti-replay
        batchExecutado[batchId] = true;

        // ── Executa transferências ─────────────────────────
        uint256 valorLiquidado   = 0;
        uint256 transferenciasOk = 0;

        for (uint256 i = 0; i < instrucoes.length; i++) {
            Instrucao calldata inst = instrucoes[i];

            if (block.timestamp > inst.deadline) {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "expirado"
                );
                continue;
            }

            if (drex.balanceOf(inst.de) < inst.valor) {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "saldo_insuficiente"
                );
                continue;
            }

            bool ok = drex.transferFrom(inst.de, inst.para, inst.valor);

            if (ok) {
                valorLiquidado   += inst.valor;
                transferenciasOk += 1;
                emit Transferido(
                    inst.seq, inst.de, inst.para,
                    inst.valor, inst.payloadHash
                );
            } else {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "transferencia_falhou"
                );
            }
        }

        // ── Ancora o Merkle root ───────────────────────────
        uint256 batchIdx = totalBatches;

        anchors[batchIdx] = AnchorRecord({
            merkleRoot:        merkleRoot,
            fromSeq:           fromSeq,
            toSeq:             toSeq,
            entryCount:        entryCount,
            timestamp:         block.timestamp,
            valorLiquidado:    valorLiquidado,
            transferenciasOk:  transferenciasOk,
            assinaturasUsadas: assinaturasValidas
        });

        lastAnchoredSeq  = toSeq;
        totalBatches     = batchIdx + 1;
        totalTransferido += valorLiquidado;

        emit Anchored(
            merkleRoot, fromSeq, toSeq,
            entryCount, block.timestamp
        );
    }

    // --------------------------------------------------------
    // Views
    // --------------------------------------------------------

    function getAnchor(uint256 batchIdx)
        external view
        returns (AnchorRecord memory)
    {
        require(batchIdx < totalBatches, "not found");
        return anchors[batchIdx];
    }

    function status()
        external view
        returns (
            uint256 lastSeq,
            uint256 batches,
            uint256 volumeTotal,
            uint256 nValidadores,
            uint256 quorum
        )
    {
        return (
            lastAnchoredSeq,
            totalBatches,
            totalTransferido,
            totalValidadores,
            quorumMinimo
        );
    }

    function getValidadores()
        external view
        returns (address[] memory)
    {
        return listaValidadores;
    }
}


