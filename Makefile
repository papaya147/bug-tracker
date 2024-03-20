build:
	cd ./frontend && make build
	cd ./backend && make build

run: build
	docker-compose up

.PHONY: run