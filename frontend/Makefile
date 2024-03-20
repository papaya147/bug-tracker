build:
	npm i
	npm run build
	docker image rm papaya147/buggy-frontend || true
	docker build -t papaya147/buggy-frontend .

.PHONY: build