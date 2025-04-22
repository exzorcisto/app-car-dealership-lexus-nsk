package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// GetDBConnection устанавливает соединение с базой данных, используя переменную среды DATABASE_URL.
func GetDBConnection() (*sql.DB, error) {
	dbConnectionString := os.Getenv("DATABASE_URL") // Получаем строку подключения из переменной среды
	if dbConnectionString == "" {
		dbConnectionString = "user=postgres password=1234 dbname=car_dealership_lexus_db sslmode=disable" // Fallback
		log.Println("DATABASE_URL не установлена, используем резервную строку подключения.  Это НЕ рекомендуется для production.")
	}

	db, err := sql.Open("postgres", dbConnectionString) // Открываем соединение с базой данных
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии базы данных: %w", err) // Возвращаем ошибку, если открытие не удалось
	}

	err = db.Ping() // Проверяем соединение с базой данных
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err) // Возвращаем ошибку, если подключение не удалось
	}

	log.Println("Подключено к базе данных!") // Логируем успех
	return db, nil                           // Возвращаем соединение с базой данных
}
