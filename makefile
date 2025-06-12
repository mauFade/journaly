include .env

MIGRATION_NAME ?= unnamed_migration

.PHONY: create_migration migrate_up migrate_down help

create_migration:
	migrate create -ext=sql -dir=internal/infrastucture/database/migrations -seq $(MIGRATION_NAME)

migrate_up:
	migrate -path=internal/infrastucture/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose up

migrate_down:
	migrate -path=internal/infrastucture/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down 1

help:
	@echo "Uso das migrations:"
	@echo "  make create_migration MIGRATION_NAME=nome_da_migration"
	@echo "  make create_migration  # usa 'unnamed_migration' como padrão"
	@echo "  make migrate_up        # executa todas as migrations pendentes"
	@echo "  make migrate_down      # desfaz a última migration"