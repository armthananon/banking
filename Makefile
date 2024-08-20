postgres:
	docker compose up -d

createdb:
	docker exec -it postgres createdb --username=myuser --owner=myuser banking

dropdb:
	docker exec -it postgres dropdb banking

migrateup:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/banking?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/armthananon/banking/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock