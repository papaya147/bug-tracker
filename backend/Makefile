postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres bug

dropdb:
	docker exec -if postgres dropdb bug

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/bug?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/bug?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

clean_run:
	make migratedown
	make migrateup
	go run main.go

.PHONY: test postgres createdb dropdb