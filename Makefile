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
test: lint docker-up-for-test
	$(ENV_TEST) go test -v -count=1 ./gameapi/...

.PHONY:test-coverage
test-coverage: lint docker-up-for-test
	$(ENV_TEST) go test -v -count=1 -coverprofile=profile ./gameapi/...
	go tool cover -html=profile -o ./cover.html

.PHONY:build
build: lint
	go build -o ./server

.PHONY:run
run: build
	$(ENV_LOCAL) ./server

.PHONY:docker-up-for-test
docker-up-for-test: down
	$(ENV_TEST) docker-compose -f docker-compose.test.yml up --build -d

.PHONY:lint
lint:
	sh ./tools/runLinters.sh
