ui:
	docker run -p 8080:8080 -v "./docs:/openapi" -e SWAGGER_JSON=/openapi/swagger.json swaggerapi/swagger-ui

ui-detach:
	docker run -d -p 8080:8080 -v "./docs:/openapi" -e SWAGGER_JSON=/openapi/swagger.json swaggerapi/swagger-ui

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
