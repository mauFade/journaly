include .env

MIGRATION_NAME ?= unnamed_migration

.PHONY: create_migration migrate_up migrate_down help

create_migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq $(MIGRATION_NAME)

migrate_up:
	migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

help:
	@echo "Uso das migrations:"
	@echo "  make create_migration MIGRATION_NAME=nome_da_migration"
	@echo "  make create_migration  # usa 'unnamed_migration' como padrão"
	@echo "  make migrate_up        # executa todas as migrations pendentes"
	@echo "  make migrate_down      # desfaz a última migration"