include .env

create_migration:
	migrate create -ext=sql -dir=db/migrations -seq $(name)

migrate_up:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: create_migration migrate_up migrate_down