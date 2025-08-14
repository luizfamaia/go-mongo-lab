# =========================
# Build & Run locais
# =========================

# build binário local
build:
	go build -o bin/user-api ./cmd/user-api

# rodar local (sem docker)
run:
	PORT=8088 MONGO_URI=mongodb://localhost:27017 go run ./cmd/user-api

# docker compose (stack completa: mongo + user-api + mongo-express)
up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f user-api


# =========================
# Testes
# =========================

.PHONY: unit it it-up it-test it-stop test-all ci quick mongo-logs mongo-shell

# Unit tests (sem tocar no Mongo)
unit:
	go test ./...

# Sobe SOMENTE o serviço 'mongo' via docker compose e espera ficar pronto
it-up:
	@echo ">> Subindo Mongo (docker compose: serviço 'mongo')..."
	docker compose up -d mongo
	@echo ">> Aguardando Mongo responder ping..."
	@until docker compose exec -T mongo mongosh --eval 'db.adminCommand("ping")' >/dev/null 2>&1; do \
	  sleep 1; \
	done
	@echo ">> Mongo pronto em localhost:27017"

# Roda SOMENTE os testes de integração (build tag 'integration', pasta integration/)
it-test:
	@echo ">> Rodando testes de integração contra mongodb://localhost:27017 ..."
	MONGO_URI="mongodb://localhost:27017" go test -tags=integration ./integrations -v

# Para o Mongo (sem derrubar o resto da stack)
it-stop:
	@echo ">> Parando serviço 'mongo' do compose..."
	docker compose stop mongo || true

# Pipeline completo de integração: sobe -> testa -> para
it: it-up it-test it-stop
	@echo "✅ Integração OK"

# Atalhos
test-all: unit it
	@echo "✅ Todos os testes passaram (unit + integração)"

ci: test-all

quick: unit

# Úteis para depurar o Mongo do compose
mongo-logs:
	docker compose logs -f mongo

mongo-shell:
	docker compose exec -it mongo mongosh
