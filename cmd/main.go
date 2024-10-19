package main

import (
  "github.com/username/block-scout/internal/db"
  "github.com/username/block-scout/internal/eth"
  "log"
)

func main() {
  blockNumber := int64(12345678)
  client, err := eth.NewEthClient("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
  if err != nil {
    log.Fatalf("Failed to create Ethereum client: %v", err)
  }

  dbConn, err := db.NewDatabase("./transfers.db")
  if err != nil {
    log.Fatalf("Failed to create database connection: %v", err)
  }

  transfers, err := client.GetUSDCTransfers(blockNumber)
  if err != nil {
    log.Fatalf("Failed to get USDC transfers: %v", err)
  }

  for _, transfer := range transfers {
    dbConn.SaveTransfer(transfer)
  }
}
