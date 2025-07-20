FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN GO111MODULE=on go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o PostGo ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/PostGo .

COPY config/ ./config/

EXPOSE 8080

CMD ["./PostGo"]