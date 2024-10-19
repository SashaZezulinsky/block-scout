.PHONY: lint test build

lint:
	go vet ./...

test:
	go test ./...

build:
	go build -o build/block-scout ./cmd/

clean:
	rm -rf build
