include .env

up:
	@echo "starting containers"
	docker-compose up --build -d --remove-orphans

down:
	@echo "stopping containers"
	docker-compose down 

build:
	go build -o $(BINARY) ./cmd/api/

start:
	./$(BINARY)

restart: build start
