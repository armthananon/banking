postgres:
	docker compose up -d

createdb:
	docker exec -it postgres createdb --username=myuser --owner=myuser banking

dropdb:
	docker exec -it postgres dropdb banking

migrateup:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/armthananon/banking/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock