ui:
	docker run -p 8080:8080 -v "./open_API:/openapi" -e SWAGGER_JSON=/openapi/openAPI.yml swaggerapi/swagger-ui

ui-detach:
	docker run -d -p 8080:8080 -v "./open_API:/openapi" -e SWAGGER_JSON=/openapi/openAPI.yml swaggerapi/swagger-ui

build:
	go build -o ./bin/chat main.go

run: build
	./bin/chat

mongo-up:
	docker compose up -d


mongo-down:
	docker compose down
