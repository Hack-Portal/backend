postgresRun:
	docker run --name hackhack-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=hackhack -d postgres:12-alpine

postgresStart:
	docker start hackhack-postgres

connectdb:
	docker exec -it hackhack-postgres psql hackhack

createdb:
	docker exec -it hackhack-postgres createdb --username=root --owner=root hackhack

dropdb:
	docker exec -it hackhack-postgres dropdb hackhack

installmigrate:
	scoop install migrate

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/hackhack?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
.PHONY: postgresRun postgresStart connectDB createdb dropdb installmigrate migrateup migratedown sqlc test