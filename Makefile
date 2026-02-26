.PHONY: help test fmt vet run-all clean

help:
	@echo "available commands:"
	@echo "  make test     - run all tests"
	@echo "  make fmt      - format all go files"
	@echo "  make vet      - run go vet"
	@echo "  make run-all  - run all modules"
	@echo "  make clean    - remove build artifacts"

test:
	@echo "running tests..."
	@go test -v ./...

fmt:
	@echo "formatting code..."
	@go fmt ./...

vet:
	@echo "running go vet..."
	@go vet ./...

run-all:
	@echo "running 01-introducao..."
	@cd 01-introducao && go run .
	@echo "\nrunning 03-escopo..."
	@cd 03-escopo && go run .
	@echo "\nrunning 04-funcoes..."
	@cd 04-funcoes && go run .
	@echo "\nrunning 05-ponteiros..."
	@cd 05-ponteiros && go run .
	@echo "\nrunning 06-structs..."
	@cd 06-structs && go run .

clean:
	@echo "cleaning build artifacts..."
	@find . -name "*.exe" -type f -delete
	@find . -name "*.test" -type f -delete
	@echo "done"
