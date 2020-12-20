ENV_FILE := .env
ENV := $(shell cat $(ENV_FILE))

ENV_TEST_FILE := ./.env.test
ENV_TEST := $(shell cat $(ENV_TEST_FILE))

.PHONY:build
build:
	go build -o ./server

.PHONY:run
run:
	export `$(ENV) | xargs` && ./server

.PHONY:test
test:
	$(ENV_TEST) go test -v -count=1 ./gameapi/...

.PHONY:docker-up
docker-up:
	$(ENV) docker-compose -f docker-compose.yml up --build -d

.PHONY:docker-down
docker-down:
	docker-compose down

.PHONY:docker-up-for-test
docker-up-for-test:
	export `$(ENV_TEST) | xargs` && docker-compose -f docker-compose.test.yml up --build -d
