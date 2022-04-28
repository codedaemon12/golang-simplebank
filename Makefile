
   
DB_URL=postgresql://root:password@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: sqlc test migrateup