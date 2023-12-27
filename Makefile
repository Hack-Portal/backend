rundb:
	docker run --name hackportal-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=hack_portal -d postgres:16

dbstart:
	docker start hackportal-postgres

postgresStop:
	docker stop hackportal-postgres

serverRun:
	go run ./cmd/app/main.go

initSwag:
	swag init -g ./cmd/app/main.go
	
test:
		go test ./... --cover --short
	
.PHONY: rundb dbstart postgresStop serverRun initSwag test