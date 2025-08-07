# MongoAPI

API simples em Go usando MongoDB e o framework Gin.

## 🚀 Funcionalidades
- Criar usuário
- Atualizar usuário
- Deletar usuário

## 💾 Stack
- Go 1.21
- Gin
- MongoDB

## 🧱 Estrutura do Projeto
```
mongoapi/
├── main.go                # Entry point
├── config/
│   └── db.go              # Conexão com MongoDB
├── models/
│   └── user.go            # Modelo de usuário
├── handlers/
│   └── user_handler.go    # Handlers das rotas
└── routes/
    └── routes.go          # Definição de rotas
```

## ▶️ Como rodar local
```bash
git clone https://github.com/seu-usuario/mongoapi.git
cd mongoapi
go mod tidy
go run main.go
```

## 🐋 Com Docker
### docker-compose.yml
```yaml
version: '3.8'
services:
  mongo:
    image: mongo
    ports:
      - '27017:27017'

  api:
    build: .
    ports:
      - '8088:8088'
    depends_on:
      - mongo
    environment:
      - MONGO_URI=mongodb://mongo:27017
```

### Dockerfile
```dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main .

EXPOSE 8088

CMD ["./main"]
```

## ✅ Testes
Crie arquivos `*_test.go` com `httptest` e `testify` para validar rotas.

```bash
go test ./...
```

## 📦 CI/CD (opcional)
Usar GitHub Actions para rodar testes antes do deploy:
- `.github/workflows/ci.yml`

---

Feito com ❤️ em Go.