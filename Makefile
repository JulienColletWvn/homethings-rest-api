include .env
export

run:
	POSTGRES_USER=${POSTGRES_USER} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} POSTGRES_URL=${POSTGRES_URL} POSTGRES_DB=${POSTGRES_DB} go run main.go

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_URL}:5432/${POSTGRES_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_URL}:5432/${POSTGRES_DB}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc run
