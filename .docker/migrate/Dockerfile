FROM golang:1.22.0 AS builder
WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go","run","cmd/migrate/main.go","-f","cmd/migrate/data"]