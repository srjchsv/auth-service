run:
	docker compose up db -d
	go run cmd/myapp/main.go


container:
	docker compose up -d


register:
	curl -X POST "http://localhost:8080/register" \
	  -H "Content-Type: application/json" \
	  -d '{"username": "user@example.com", "password": "mypassword"}'

login:
	curl -X POST http://localhost:8080/login \
	-H 'Content-Type: application/json' \
	-d '{ "username": "user@example.com","password": "mypassword"}'



