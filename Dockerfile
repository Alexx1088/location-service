FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# Install curl and download migrate CLI
RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz \
    | tar xvz \
    && mv migrate /usr/local/bin/

# Copy and download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the Go binary
COPY . .
RUN go build -o location-service ./cmd/location-service/main.go

# Final stage (runtime)
FROM alpine:latest
WORKDIR /app

# Copy compiled app and migrate CLI from builder stage
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/
COPY --from=builder /app/location-service .

EXPOSE 8080

CMD ["./location-service"]
