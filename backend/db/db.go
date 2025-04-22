package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load(".env") // Загрузка переменных окружения
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL") // Получение URL из окружения
	if dbURL == "" {
		dbURL = "user=postgres password=1234 dbname=car_dealership_lexus_db sslmode=disable" // Значение по умолчанию (ПЛОХАЯ практика)
	}

	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось проверить соединение с базой данных: %v", err)
	}

	fmt.Println("Подключено к базе данных!")
}
