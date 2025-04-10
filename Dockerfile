# Используем официальный образ Go
FROM golang:1.23

# Устанавливаем рабочую директорию для приложения
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Открываем порт
EXPOSE 8080

# Запуск приложения с go run
CMD ["go", "run", "cmd/main.go"]
