FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .\cmd\app\main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY serviceAccount.json .


EXPOSE 8080
CMD [ "/app/main"]
