include .env.local

.PHONY: up down version

MIGRATE_CMD=migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

create_migration: 
	migrate create -ext=sql -dir=internal/database/migrations -seq init
migrate_up: 
	$(MIGRATE_CMD) -verbose up

migrate_force: 
	$(MIGRATE_CMD) -verbose force

migrate_down: 
	$(MIGRATE_CMD) -verbose down 1

migrate_version:
	$(MIGRATE_CMD) version

commit:
	git add . && git-cz

push:
	git push -u origin HEAD