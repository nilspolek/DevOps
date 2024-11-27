generate:
	go generate

build: generate
	go build -o ./bin/chat main.go

run: build
	./bin/chat

mongo-up:
	docker compose up -d

mongo-down:
	docker compose down

fmt:
	go fmt ./...
	swag fmt
