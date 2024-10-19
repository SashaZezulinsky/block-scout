
# Block Scout

Block Scout is a Go-based application that allows you to interact with the Ethereum blockchain to extract and store data related to USDC transfers in a SQLite database

## Project Structure

The project is organized as follows:

```
block-scout/
├── cmd/                     # Main application entry point
│   └── main.go              # Main executable code for running the application
├── internal/
│   ├── db/                  # DB package for interacting with SQLite
│   │   ├── db.go            # DB connection and operations
│   │   ├── db_test.go       # DB connection and operations tests
│   ├── eth/                 # Ethereum client package for blockchain interaction
│   │   ├── eth.go           # Ethereum client to connect and fetch tx data
│   │   ├── eth_test.go      # Ethereum client connection and data fetching tests
│   ├── models/              # Models representing the data structure
│   │   ├── transfer.go      # Transfer model representing a USDC transfer
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
├── LICENSE                  # License file
├── README.md                # Project documentation (this file)
```

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher
- Ethereum node RPC URL (Infura, Alchemy, or self-hosted)

## Installation

Clone the repository:

```bash
git clone https://github.com/SashaZezulinsky/block-scout.git
cd block-scout
```

Install the necessary Go modules:

```bash
go mod tidy
```

## Usage

To run the application, use the following command:

```bash
go run cmd/main.go -block <block_number> -db <path_to_database_file>
```

### Command Line Arguments

- `-block` (required): The block number to fetch real USDC transfers from.
- `-db` (optional): Path to the SQLite database file. Default is `./transfers.db`.

### Environment Variables

- `ETH_RPC_URL` (required): The URL of your Ethereum node RPC (e.g., from Infura or Alchemy).

### Example

Run the following command to get USDC transfers for a specific block number and save them to a database:

```bash
ETH_RPC_URL=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID go run cmd/main.go -block 12345678 -db transfers.db
```

This will connect to the Ethereum node, extract the USDC transfers from the specified block, store them in the specified database, and print out the stored data.

## Testing

To run the unit tests:

```bash
make test
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
