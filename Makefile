run:
	docker-compose up --build

rund:
	docker-compose up -d --build

down:
	docker-compose down

initSwag:
	swag init -g ./cmd/app/main.go
	
test:
	go test ./... --cover --short

user="user"
file="cmd/seeder/seed.json"
seed:
	go run cmd/seeder/main.go -f $(file) -u $(user)

.PHONY: run rund down initSwag test