postgresRun:
	docker run --name hackhack-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=hackhack -d postgres:12-alpine

postgresStart:
	docker start hackhack-postgres

postgresStop:
	docker stop hackhack-postgres

resetdb:
	docker stop hackhack-postgres
	docker start hackhack-postgres
	docker exec -it hackhack-postgres dropdb hackhack
	docker exec -it hackhack-postgres createdb --username=root --owner=root hackhack

migrateup:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose up

migrateup1:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose up 1

migratedown:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose down

migratedown1:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

serverRun:
	go run ./cmd/app/main.go

makeSwagger:
	swag fmt
	swag init -g ./cmd/app/main.go
	
test:
	go test -coverpkg=./...  ./...
	
.PHONY: postgresRun postgresStart postgresStop resetDB installmigrate migrateup migrateup1 migratedown migratedown1 sqlc serverRun test makeSwagger