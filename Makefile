# Subir com Docker
up:
	docker compose up -d

# Rodar a API local (sem docker)
run:
	go run main.go

# Rodar testes
test:
	go test ./...

# Instalar dependências
tidy:
	go mod tidy

# Build manual (opcional)
build:
	go build -o main .

# Limpar binários
clean:
	rm -f main
