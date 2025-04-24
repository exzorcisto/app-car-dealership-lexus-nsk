package main

import (
	"log"
	"net/http"

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
	err := db.DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to database")

	router := mux.NewRouter()

	// Маршруты для работы с автомобилями
	router.HandleFunc("/cars", handlers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	router.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	// router.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	// router.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Обертка роутера с CORS
	handler := c.Handler(router)

	log.Println("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
