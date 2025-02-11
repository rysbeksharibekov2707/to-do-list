# Используем официальный образ Go для сборки
FROM golang:1.18 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости и собираем проект
RUN go mod tidy
RUN go build -o main .

# Используем минимальный образ для запуска
FROM debian:buster

# Устанавливаем PostgreSQL клиент
RUN apt-get update && apt-get install -y postgresql-client --fix-missing

# Копируем собранное приложение из предыдущего этапа сборки
COPY --from=builder /app/main /app/main

# Открываем порт для приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["/app/main"]
