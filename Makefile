all: init_db test build

run: build start

init_db:
	@echo " >> initialize database"
	@go run cmd/migration/main.go

test:
	@echo " >> running tests"
	@go test -v -cover -race ./...

build:
	@echo " >> building binaries"
	@go build -o bin/main main.go
	
start:
	@echo " >> starting binaries"
	@./bin/main
