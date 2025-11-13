.PHONY: up down build logs migrate

up:
\tdocker compose -f deploy/docker-compose.yml up -d --build

down:
\tdocker compose -f deploy/docker-compose.yml down -v

build:
\tdocker compose -f deploy/docker-compose.yml build

logs:
\tdocker compose -f deploy/docker-compose.yml logs -f

# local DB migrations (adjust per service)
migrate-auth:
\tmigrate -path services/auth/migrations -database "postgres://postgres:postgres@localhost:5432/social?sslmode=disable" up
