.PHONY: lint test build

clean:
	rm -rf build *.db

lint:
	go vet ./...

test:
	go test ./...

build: clean
	go build -o build/block-scout ./cmd/

