# Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Package
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY migrations /root/migrations

EXPOSE 8080
CMD ["./app"]