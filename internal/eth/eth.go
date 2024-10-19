package eth

import (
    "context"
    "log"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/SashaZezulinsky/block-scout/internal/models"
)

var transferSignatureHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

type EthClient struct {
    client *ethclient.Client
}

func NewEthClient(rpcURL string) (*EthClient, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, err
    }
    return &EthClient{client: client}, nil
}

func (e *EthClient) GetUSDCTransfers(blockNumber int64) ([]models.Transfer, error) {
    block, err := e.client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
    if err != nil {
        return nil, err
    }

    var transfers []models.Transfer
    for _, tx := range block.Transactions() {
        // Check if the transaction type is supported
        if tx.Type() != types.LegacyTxType && tx.Type() != types.DynamicFeeTxType {
            continue
        }

        receipt, err := e.client.TransactionReceipt(context.Background(), tx.Hash())
        if err != nil {
            log.Printf("Failed to get receipt for tx %s: %v", tx.Hash().Hex(), err)
            continue
        }
 
        for _, logEntry := range receipt.Logs {
            if len(logEntry.Topics) == 3 && logEntry.Address == common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48") { // USDC contract
                if logEntry.Topics[0] == transferSignatureHash {
                    transfers = append(transfers, models.Transfer{
                        Sender:    common.HexToAddress(logEntry.Topics[1].Hex()).Hex(),
                        Recipient: common.HexToAddress(logEntry.Topics[2].Hex()).Hex(),
                        Value:     new(big.Int).SetBytes(logEntry.Data).String(),
                    })
                }
            }
        }
    }
    return transfers, nil
}
