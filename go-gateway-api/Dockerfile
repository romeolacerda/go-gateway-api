# Estágio de construção
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]
