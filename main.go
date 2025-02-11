package main

import (
	"log"
	"net/http"
	"todoactual/config"
	"todoactual/database"
	"todoactual/handlers"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer database.DB.Close()

	// Создание таблицы
	err = database.CreateTasksTable()
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	// Настройка маршрутов
	http.HandleFunc("/tasks", handlers.GetTasks)

	// Запуск сервера
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
