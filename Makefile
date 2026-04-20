DB_MIGRATE_URL = postgres://login:pass@localhost:5432/postgres?sslmode=disable
MIGRATE_PATH = ./migration

run:
	go run ./cmd/app

up:
	docker compose up --build --force-recreate

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "./migrations" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "./migrations" down -all

drop-db:
	docker compose exec -T pgdb psql -U login -d postgres -c "DROP DATABASE IF EXISTS postgres;"

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

