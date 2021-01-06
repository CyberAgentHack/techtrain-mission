

ENV_FILE := .env
ENV := $(shell cat $(ENV_FILE))

ENV_TEST_FILE := .env.test
ENV_TEST := $(shell cat $(ENV_TEST_FILE))

ENV_LOCAL_FILE := .env.local
ENV_LOCAL := $(shell cat $(ENV_LOCAL_FILE))

.PHONY:docker-up
docker-up: down
	$(ENV) docker-compose -f docker-compose.yml up --build -d

.PHONY:down
down:
	docker-compose down

.PHONY:test
test: lint
	$(ENV_TEST) go test -covermode=atomic -coverprofile=coverage.out ./gameapi/...

.PHONY:test-with-coverage
test-with-coverage: 
	$(ENV_TEST) go test -covermode=atomic -coverprofile=coverage.out ./gameapi/...
	go tool cover -html=coverage.out -o ./cover.html

.PHONY:build
build: lint
	go build -o ./server

.PHONY:run
run: build
	$(ENV_LOCAL) ./server

.PHONY:lint
lint:
	sh ./tools/runLinters.sh
