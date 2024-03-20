postgresup:
	docker run --name postgres -p 5432:5432 -e PGUSER=postgres -e POSTGRES_PASSWORD=password -v ./postgres/:/var/lib/postgresql/data/ -d postgres
	@echo "Waiting for PostgreSQL to be ready..."
	@docker run --rm \
		--link postgres:postgres \
		--entrypoint sh \
		postgres \
		-c 'while ! pg_isready -h "$$POSTGRES_PORT_5432_TCP_ADDR" -p "$$POSTGRES_PORT_5432_TCP_PORT" > /dev/null; do echo "Waiting for PostgreSQL to be ready..."; sleep 1; done'

postgresdown:
	docker stop postgres
	docker rm postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres buggy || true

dropdb:
	docker exec -it postgres dropdb buggy

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/buggy?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/buggy?sslmode=disable" -verbose down

setup_postgres: postgresup createdb migrateup

test:
	go test -v -cover ./...

clean_run:
	make migratedown
	make migrateup
	make test
	go run main.go

build:
	docker image rm papaya147/buggy-backend || true
	env GOOS=linux CGO_ENABLED=0 go build -o server-app main.go
	docker build -t papaya147/buggy-backend:latest .

.PHONY: test postgres createdb dropdb