.PHONY: run test

run: 
	go run ./cmd/app/main.go

test: 
	@echo "Running tests..."
	go test ./test/... -v

all: test run