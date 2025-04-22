package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/db"       // Замените на ваш путь к модулю
	"backend/handlers" // Замените на ваш путь к модулю

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/cars", handlers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	router.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	// router.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	// router.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your React app's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router) // Wrap the router with CORS
	fmt.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler)) // Pass handler
}
