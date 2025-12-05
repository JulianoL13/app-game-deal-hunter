.PHONY: fmt lint test build run clean

MODULE := github.com/julianoL13/app-game-deal-hunter

# Formata o código
fmt:
	gofmt -w .
	goimports -w -local $(MODULE) .

# Roda o linter
lint:
	golangci-lint run ./...

# Formata e depois roda o linter
check: fmt lint

# Roda os testes
test:
	go test -v ./...

# Roda os testes com cobertura
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Build da aplicação
build:
	go build -o bin/app ./cmd/main.go

# Roda a aplicação
run:
	go run ./cmd/main.go

# Limpa artefatos
clean:
	rm -rf bin/ coverage.out coverage.html

# Instala ferramentas de desenvolvimento
tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Gera mocks (limpa e regenera)
mocks:
	find . -type d -name 'mocks' -exec rm -rf {} + 2>/dev/null || true
	mockery
