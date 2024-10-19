package main

import (
	"flag"
	"fmt"
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

	// Use command-line flags for block number and database file
	blockNumberFlag := flag.String("block", "", "Block number with real USDC transfers")
	dbFileFlag := flag.String("db", "./transfers.db", "Path to the database file")
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

	dbConn, err := db.NewDatabase(*dbFileFlag)
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

	dumpTransfers(dbConn)
}

func dumpTransfers(dbConn *db.Database) {
	transfers, err := dbConn.GetTransfers()
	if err != nil {
		log.Fatalf("Failed to fetch transfers: %v", err)
	}
	fmt.Println("USDC Transfers:")
	for _, transfer := range transfers {
		log.Printf("Sender: %s, Recipient: %s, Value: %s", transfer.Sender, transfer.Recipient, transfer.Value)
	}
}
