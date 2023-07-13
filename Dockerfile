# golang install
FROM golang:1.20-alpine

## timezone update
RUN apk add tzdata
ENV TZ=Asia/Tokyo
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY app.env .
COPY serviceAccount.json .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .

## container listen port
EXPOSE 8080

CMD ./out/dist