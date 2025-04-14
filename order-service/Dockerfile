# Используем официальный образ Go
FROM golang:1.22.4-alpine


# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go mod и sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]
