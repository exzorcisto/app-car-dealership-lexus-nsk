package main

import (
	"log"
	"net/http"
	"time"

	"backend/db"
	"backend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Инициализация подключения к базе данных
	db.InitDB()
	defer db.DB.Close()

	// Проверка соединения с базой данных
	if err := db.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to database")

	router := mux.NewRouter()

	// Инициализация обработчиков
	carHandler := &handlers.CarHandler{DB: db.DB}
	empHandler := &handlers.EmployeeHandler{DB: db.DB}

	// Инициализация обработчиков
	authHandler := &handlers.AuthHandler{
		DB:          db.DB,
		JWTSecret:   "your-very-secure-secret-key", // Замените на реальный секрет
		TokenExpiry: 24 * time.Hour,
	}

	// Маршруты аутентификации
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	router.HandleFunc("/validate", authHandler.ValidateToken).Methods("GET")

	// Маршруты для работы с автомобилями
	router.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", carHandler.GetCar).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	// Маршруты для работы с сотрудниками
	router.HandleFunc("/employees", empHandler.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", empHandler.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", empHandler.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", empHandler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", empHandler.DeleteEmployee).Methods("DELETE")

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Обертка роутера с CORS
	handler := c.Handler(router)

	// Запуск сервера
	log.Println("Server starting on port 8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
