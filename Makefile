
.PHONY:build
build:
	go build -o ./server

.PHONY:server
run:
	./server

.PHONY:docker-up
docker-up:
	docker-compose up --build
