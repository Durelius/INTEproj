.PHONY: run test

run: 
	go run ./cmd/app/main.go


test:
	@echo "Running tests..."
	go test ./... -v

stan:
	@echo "Running static analysis..."
	staticcheck ./...

# Run all tests, and check coverage for all files in ./internal except for the ones in internal/assets.
# There is no reason to test methods that just return an ascii string.
test-cover: 	
	@echo "Running tests..."
	@go test -covermode=atomic \
		-coverpkg=$$(go list ./... | tr '\n' ',' | sed 's/,$$//') \
		-coverprofile=coverage_raw.out \
		./... -v > test.log 2>&1
	@grep -v '/internal/assets/' coverage_raw.out > coverage.out
	@echo "All tests passed."
	go tool cover -html=coverage.out
	rm coverage_raw.out | rm coverage.out | rm test.log

build: 
	go build -o ./bin/game ./cmd/app/main.go

all: test stan run

