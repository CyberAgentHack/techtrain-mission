ENV_FILE := .env
ENV := $(export `cat ${ENV_FILE} | xargs`)

.PHONY:build
build:
	go build -o ./server

.PHONY:run
run:
	$(ENV) && ./server

.PHONY:test
test:
	$(ENV_TEST) go test -v -count=1 ./gameapi/...

.PHONY:docker-up
docker-up:
	$(ENV) docker-compose up --build -d

.PHONY:docker-down
docker-down:
	docker-compose down
