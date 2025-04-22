package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"app_car_dealership_lexus_nsk/db"       // Замените "your_module" на имя вашего модуля
	"app_car_dealership_lexus_nsk/handlers" // Замените "your_module" на имя вашего модуля

	"github.com/gorilla/mux"
)

func main() {
	// Загружаем переменные среды.  Рекомендуется использовать библиотеку, такую как "github.com/joho/godotenv" для более надежной обработки файла .env.
	err := loadEnv()
	if err != nil {
		log.Printf("Ошибка при загрузке файла .env: %v (Продолжаем без него)", err)
	}

	// Подключение к базе данных
	database, err := db.GetDBConnection()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer database.Close()

	// Устанавливаем соединение с базой данных для обработчиков
	handlers.SetDB(database)

	// Инициализируем Router
	router := mux.NewRouter()

	// Определяем API endpoints
	router.HandleFunc("/api/models", handlers.GetModelsHandler).Methods("GET")
	router.HandleFunc("/api/models/{id}", handlers.GetModelHandler).Methods("GET")
	router.HandleFunc("/api/models", handlers.CreateModelHandler).Methods("POST")
	router.HandleFunc("/api/models/{id}", handlers.UpdateModelHandler).Methods("PUT")
	router.HandleFunc("/api/models/{id}", handlers.DeleteModelHandler).Methods("DELETE")

	// Запускаем сервер
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	log.Printf("Сервер слушает на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, allowCORS(router)))
}

// loadEnv загружает переменные среды из файла .env (если он существует).
func loadEnv() error {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil // .env file does not exist, skip loading
	}

	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	// Very basic .env parsing.  Consider using a library like "github.com/joho/godotenv" for more robust parsing.
	decoder := json.NewDecoder(file)
	var env map[string]string
	err = decoder.Decode(&env)

	if err != nil {
		return fmt.Errorf("error decoding .env file: %w", err)
	}

	for key, value := range env {
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("error setting environment variable %s: %w", key, err)
		}

		log.Printf("Set Environment variable %s to %s", key, value)
	}

	return nil
}

// allowCORS - это middleware-функция, которая включает Cross-Origin Resource Sharing (CORS).
// Она устанавливает необходимые заголовки для разрешения запросов с любого источника (только для разработки!).
// В производственной среде вам следует ограничить разрешенные источники только теми, которым вы доверяете.
func allowCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем запросы с любого источника (только для разработки!)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
