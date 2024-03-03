include .env

# DATABASE
db-migrateup:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} up

db-migratedown:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} down

db-migrateforce:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} force 1

sqlc:
	@sqlc generate

# Test
test:
	@go test -v -cover ./...
