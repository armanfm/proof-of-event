// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/**
 * @title PoEAnchor
 * @notice Ancora batches pré-validados off-chain (Go) E liquida os tokens.
 *
 * Duas ações atômicas numa única transação:
 *   1. Ancora o Merkle root (prova imutável)
 *   2. Executa as transferências de tokens (liquidação real)
 *
 * O Go fez todo o trabalho:
 *   ✓ Ordenação FIFO
 *   ✓ Pré-validação (saldo, endereços, valor)
 *   ✓ Merkle root calculado
 *
 * O contrato só executa — não valida nada internamente.
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
        address indexed anchoredBy,
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

    // --------------------------------------------------------
    // Estado
    // --------------------------------------------------------

    address public immutable owner;
    IDREX   public immutable drex;       // token DREX

    uint256 public lastAnchoredSeq;
    uint256 public totalBatches;
    uint256 public totalTransferido;     // volume total liquidado

    struct AnchorRecord {
        bytes32 merkleRoot;
        uint256 fromSeq;
        uint256 toSeq;
        uint256 entryCount;
        uint256 timestamp;
        address anchoredBy;
        uint256 valorLiquidado;          // total liquidado neste batch
        uint256 transferenciasOk;        // quantas transferências ok
    }

    // Instrução de transferência — vem pré-validada do Go
    struct Instrucao {
        uint256 seq;
        address de;
        address para;
        uint256 valor;       // em wei
        bytes32 payloadHash;
        uint256 deadline;    // validade unix timestamp
    }

    mapping(uint256 => AnchorRecord) public anchors;

    // --------------------------------------------------------
    // Constructor
    // --------------------------------------------------------

    constructor(address _drex) {
        require(_drex != address(0), "DREX zero address");
        owner = msg.sender;
        drex  = IDREX(_drex);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "PoEAnchor: not owner");
        _;
    }

    // --------------------------------------------------------
    // Core — ancora + liquida atomicamente
    // --------------------------------------------------------

    /**
     * @notice Recebe batch pré-validado do Go e:
     *         1. Ancora o Merkle root (prova imutável)
     *         2. Executa todas as transferências (liquidação)
     *
     * @param merkleRoot  Calculado pelo Go — resume o batch inteiro
     * @param fromSeq     Sequência inicial do batch
     * @param toSeq       Sequência final do batch
     * @param entryCount  Quantidade de eventos
     * @param instrucoes  Transferências pré-validadas pelo Go
     */
    function receberELiquidar(
        bytes32      merkleRoot,
        uint256      fromSeq,
        uint256      toSeq,
        uint256      entryCount,
        Instrucao[]  calldata instrucoes
    )
        external
        onlyOwner
    {
        // Validações mínimas de sequência
        require(merkleRoot != bytes32(0),    "PoEAnchor: root invalido");
        require(toSeq >= fromSeq,            "PoEAnchor: intervalo invalido");
        require(entryCount > 0,              "PoEAnchor: batch vazio");
        require(
            fromSeq == lastAnchoredSeq + 1,
            "PoEAnchor: gap na sequencia"
        );

        // ── Executa as transferências ──────────────────────
        uint256 valorLiquidado   = 0;
        uint256 transferenciasOk = 0;

        for (uint256 i = 0; i < instrucoes.length; i++) {
            Instrucao calldata inst = instrucoes[i];

            // Deadline — rejeita instrução expirada
            if (block.timestamp > inst.deadline) {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "expirado"
                );
                continue;
            }

            // Saldo mínimo — última linha de defesa
            // (Go já verificou, mas contrato não confia cegamente)
            if (drex.balanceOf(inst.de) < inst.valor) {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "saldo_insuficiente"
                );
                continue;
            }

            // Executa a transferência
            bool ok = drex.transferFrom(inst.de, inst.para, inst.valor);

            if (ok) {
                valorLiquidado   += inst.valor;
                transferenciasOk += 1;
                emit Transferido(
                    inst.seq,
                    inst.de,
                    inst.para,
                    inst.valor,
                    inst.payloadHash
                );
            } else {
                emit TransferenciaFalhou(
                    inst.seq, inst.de, inst.para,
                    inst.valor, "transferencia_falhou"
                );
            }
        }

        // ── Ancora o Merkle root ───────────────────────────
        // Acontece APÓS as transferências
        // Se chegou aqui — batch foi processado
        uint256 batchIdx = totalBatches;

        anchors[batchIdx] = AnchorRecord({
            merkleRoot:        merkleRoot,
            fromSeq:           fromSeq,
            toSeq:             toSeq,
            entryCount:        entryCount,
            timestamp:         block.timestamp,
            anchoredBy:        msg.sender,
            valorLiquidado:    valorLiquidado,
            transferenciasOk:  transferenciasOk
        });

        lastAnchoredSeq  = toSeq;
        totalBatches     = batchIdx + 1;
        totalTransferido += valorLiquidado;

        emit Anchored(
            merkleRoot,
            fromSeq,
            toSeq,
            entryCount,
            msg.sender,
            block.timestamp
        );
    }

    // --------------------------------------------------------
    // Views — auditoria
    // --------------------------------------------------------

    function getAnchor(uint256 batchIdx)
        external view
        returns (AnchorRecord memory)
    {
        require(batchIdx < totalBatches, "PoEAnchor: not found");
        return anchors[batchIdx];
    }

    function status()
        external view
        returns (
            uint256 lastSeq,
            uint256 batches,
            uint256 volumeTotal
        )
    {
        return (lastAnchoredSeq, totalBatches, totalTransferido);
    }
}

