# Load environment variables from .env file
include .env

export

# Run the app
run:
	docker compose up -d
	go run cmd/myapp/main.go

register:
	curl -X POST "http://localhost:8080/register" \
	  -H "Content-Type: application/json" \
	  -d '{"username": "user@example.com", "password": "mypassword"}'

login:
	curl -X POST http://localhost:8080/login \
	-H 'Content-Type: application/json' \
	-d '{ "username": "user@example.com","password": "mypassword"}'



