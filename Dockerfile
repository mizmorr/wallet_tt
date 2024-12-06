FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY ./cmd/bin ./cmd/bin

# Устанавливаем рабочую директорию для сборки
WORKDIR /app/cmd/bin

# Сборка бинарного файла
RUN go build -o main .

# Устанавливаем порт, который будет использовать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
