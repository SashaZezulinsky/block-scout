package eth

import (
    "testing"
    "log"
    "os"
)

func TestNewEthClient(t *testing.T) {
    rpcURL := os.Getenv("ETH_RPC_URL")
    if rpcURL == "" {
	    log.Fatal("Environment variable ETH_RPC_URL is required")
    }

    _, err := NewEthClient(rpcURL)
    if err != nil {
        t.Fatalf("Failed to create Ethereum client: %v", err)
    }
}

func TestGetUSDCTransfers(t *testing.T) {
    rpcURL := os.Getenv("ETH_RPC_URL")
    if rpcURL == "" {
	    log.Fatal("Environment variable ETH_RPC_URL is required")
    }

    client, err := NewEthClient(rpcURL)
    if err != nil {
        t.Fatalf("Failed to create Ethereum client: %v", err)
    }

    blockNumber := int64(21001969) // Block number with real USDC transfers
    transfers, err := client.GetUSDCTransfers(blockNumber)
    if err != nil {
        t.Fatalf("Failed to get USDC transfers: %v", err)
    }

    for _, transfer := range transfers {
        if transfer.Sender == "" || transfer.Recipient == "" || transfer.Value == "" {
            t.Errorf("Invalid transfer data: %+v", transfer)
        }
        log.Printf("Sender: %s", transfer.Sender)
        log.Printf("Recipient: %s", transfer.Recipient)
        log.Printf("Value: %s", transfer.Value)
    }
}
