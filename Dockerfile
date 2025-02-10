FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник
RUN go build -o main .

#To lighten the golang image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .env

# Запускаем приложение
CMD ["./main"]
