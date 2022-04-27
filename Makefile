sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: sqlc test