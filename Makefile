# build binário local
build:
	go build -o bin/user-api ./cmd/user-api

# rodar local (sem docker)
run:
	PORT=8088 MONGO_URI=mongodb://localhost:27017 go run ./cmd/user-api

# docker compose
up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f user-api