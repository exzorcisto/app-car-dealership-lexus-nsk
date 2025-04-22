package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/db"     // Замените на ваш путь к модулю
	"backend/models" // Замените на ваш путь к модулю

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT CarID, ModelID, TrimLevel, Year, VIN, Price, Color, Mileage, IsNew FROM Cars") // Адаптируйте запрос
	if err != nil {
		log.Printf("Error querying cars: %v", err)
		http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.CarID, &car.ModelID, &car.TrimLevel, &car.Year, &car.VIN, &car.Price, &car.Color, &car.Mileage, &car.IsNew); err != nil {
			log.Printf("Error scanning car row: %v", err)
			http.Error(w, "Failed to process car data", http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Printf("Error encoding cars to JSON: %v", err)
		http.Error(w, "Failed to encode cars", http.StatusInternalServerError)
		return
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carIDStr := vars["id"]

	carID, err := strconv.Atoi(carIDStr) // Convert string ID to integer
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}
	row := db.DB.QueryRow("SELECT CarID, ModelID, TrimLevel, Year, VIN, Price, Color, Mileage, IsNew FROM Cars WHERE CarID = $1", carID) // Adapt query

	var car models.Car
	err = row.Scan(&car.CarID, &car.ModelID, &car.TrimLevel, &car.Year, &car.VIN, &car.Price, &car.Color, &car.Mileage, &car.IsNew)
	if err != nil {
		log.Printf("Error scanning car row: %v", err)
		http.Error(w, "Failed to process car data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("Error encoding cars to JSON: %v", err)
		http.Error(w, "Failed to encode cars", http.StatusInternalServerError)
		return
	}
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
INSERT INTO Cars (ModelID, TrimLevel, Year, VIN, Price, Color, Mileage, IsNew)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING CarID`
	id := 0
	err = db.DB.QueryRow(sqlStatement, car.ModelID, car.TrimLevel, car.Year, car.VIN, car.Price, car.Color, car.Mileage, car.IsNew).Scan(&id)
	if err != nil {
		log.Printf("Error creating car: %v", err)
		http.Error(w, "Failed to create car", http.StatusInternalServerError)
		return
	}

	car.CarID = id // Set the CarID on the car object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}
