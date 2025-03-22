FROM golang:1.23

WORKDIR /app

# Сначала копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Теперь копируем весь проект
COPY . .

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]
