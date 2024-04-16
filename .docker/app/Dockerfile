FROM golang:1.22.0 AS builder
WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go

FROM scratch as app
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/casbin_model.conf .
COPY --from=builder /build/cmd/migrations ./cmd/migrations
COPY --from=builder /build/app .

ENV PORT=8080

EXPOSE ${PORT}

CMD ["/app/app"]