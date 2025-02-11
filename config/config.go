package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBName     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func LoadConfig() *Config {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	fmt.Println("DBUser:", os.Getenv("DB_USER"))
	fmt.Println("DBPassword:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DBHost:", os.Getenv("DB_HOST"))
	fmt.Println("DBPort:", os.Getenv("DB_PORT"))
	fmt.Println("DBName:", os.Getenv("DB_NAME"))

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

func (c *Config) ConnectionString() string {
	//fmt.Println("Connection String:", c.ConnectionString())
	//return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	//	c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPassword)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	//return "host=127.0.0.1 port=5432 user=youruser dbname=yourdb sslmode=disable password=yourpassword"
}
