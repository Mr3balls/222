# Используем официальный образ Golang
FROM golang:1.22.4-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Компилируем приложение
RUN go build -o api-gateway .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./api-gateway"]
