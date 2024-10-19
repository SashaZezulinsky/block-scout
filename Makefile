.PHONY: lint test build

lint:
	go vet ./...

test:
	go test ./...

build:
	go build -o block-scout ./cmd/block-scout
