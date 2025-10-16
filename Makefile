.PHONY: run test

run: 
	go run ./cmd/app/main.go

test: 
	go test ./test/player