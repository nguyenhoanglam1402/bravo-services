# --- Builder stage ---
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# --- Runtime stage ---
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
RUN chmod +x ./main

EXPOSE 8080
CMD ["./main"]