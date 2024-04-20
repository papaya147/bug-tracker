build:
	cd ./frontend && make build
	cd ./backend && make build

deploy: build
	docker-compose up

.PHONY: run