postgres:
	docker run --name dev-postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=bank-app -d -p 5432:5432 postgres:16-alpine

createdb:
	docker exec -it dev-postgres16 createdb --username=root --owner=root bank-app

dropdb:
	docker exec -it dev-postgres16 dropdb bank-app

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/bank-app?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/bank-app?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc