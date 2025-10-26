
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o location-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/location-service .

EXPOSE 8080

CMD ["./location-service"]
