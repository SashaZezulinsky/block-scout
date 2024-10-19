package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/SashaZezulinsky/block-scout/internal/db"
	"github.com/SashaZezulinsky/block-scout/internal/eth"
)

func main() {
	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		log.Fatal("Environment variable ETH_RPC_URL is required")
	}

	// Use command-line flag for block number
	blockNumberFlag := flag.String("block", "", "Block number with real USDC transfers")
	flag.Parse()

	if *blockNumberFlag == "" {
		log.Fatal("Block number is required. Use the -block flag to provide it.")
	}

	blockNumber, err := strconv.ParseInt(*blockNumberFlag, 10, 64)
	if err != nil {
		log.Fatalf("Invalid block number: %v", err)
	}

	client, err := eth.NewEthClient(rpcURL)
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

