package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// InitDB загружает конфиг и подключается к БД
func InitDB() (*sql.DB, error) {
	err := godotenv.Load() // Загружаем .env
	if err != nil {
		log.Println("Файл .env не найден, использую переменные окружения")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("БД недоступна: %w", err)
	}

	log.Println("Успешное подключение к БД")
	return db, nil
}
