include app.env

# DATABASE
db-migrateup:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} up

db-migratedown:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} down

db-migrateup1:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} up 1

db-migratedown1:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} down 1

db-migrateforce:
	@migrate -path /Users/theaveasso/Dev/go/bab/migrations -database ${POSTGRES_DSN} force 2

sqlc:
	@sqlc generate

server:
	@go run ./cmd/bab

mock:
	@mockgen -destination internal/db/mock/store.go theaveasso.bab/internal/db Store

# Test
test:
	@go test -v -cover ./...
