-include .env

init:
	cp .env.dist .env
	createdb
	sqlc
	migrateup

start:
	docker compose up

createdb:
	docker-compose exec postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker-compose exec postgres dropdb ${POSTGRES_DB}
migrateup1:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down 1
migrateup:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/tr0b/movies-code-challenge/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
