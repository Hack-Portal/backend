FROM golang:1.21.4 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/app/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/casbin_model.conf .
COPY --from=builder /app/cmd/migrations ./cmd/migrations
COPY --from=builder /app/main .

# Server Settings
ENV SERVER_ADDR=8080
ENV SERVER_SHUTDOWN_TIMEOUT=5
ENV SERVER_CONTEXT_TIMEOUT=5
ENV SERVER_VERSION=
ENV SERVER_ADMIN_INIT_PASSWORD=

# Database Connection Settings
ENV DB_HOST=
ENV DB_PORT=5432
ENV DB_USER=
ENV DB_PASSWORD=
ENV DB_NAME=
ENV DB_SSLMODE=disable
ENV DB_TIMEZONE=Asia/Tokyo

# Redis Connection Settings
ENV REDIS_HOST=
ENV REDIS_PORT=
ENV REDIS_PASSWORD=
ENV REDIS_DB=

# NewRelic Settings
ENV NEWRELIC_APPLICATION_NAME=
ENV NEWRELIC_LICENSE_KEY=

# Cloudflare Settings
ENV BUCKETS_ENDPOINT=
ENV BUCKETS_ACCOUNT_ID=
ENV BUCKETS_ACCESS_KEY_ID=
ENV BUCKETS_ACCESS_KEY_SECRET=
ENV CLOUDFLARE_BUCKET=
ENV BUCKETS_EXPIRED=

EXPOSE 8080
CMD [ "/app/main"]
