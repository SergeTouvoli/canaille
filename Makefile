.PHONY: test test-verbose run build fmt vet check

test:
	go test ./...

test-verbose:
	go test -v ./...

run:
	go run ./cmd/canaille ./testdata/simple-compose.yaml

build:
	go build -o bin/canaille ./cmd/canaille

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet vet 