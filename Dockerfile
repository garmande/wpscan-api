FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o wpscan-api ./cmd/server

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/wpscan-api .

EXPOSE 8080

CMD ["./wpscan-api"]
