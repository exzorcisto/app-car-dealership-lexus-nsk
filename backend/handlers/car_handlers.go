package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT carid, trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2 FROM cars")
	if err != nil {
		log.Printf("Error querying cars: %v", err)
		http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(
			&car.CarID,
			&car.TrimLevel,
			&car.Year,
			&car.VIN,
			&car.Price,
			&car.Color,
			&car.Bodywork,
			&car.Engine,
			&car.EngineCapacity,
			&car.Fuel,
			&car.Image,
			&car.Description1,
			&car.Description2,
		); err != nil {
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

	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	row := db.DB.QueryRow("SELECT carid, trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2 FROM cars WHERE carid = $1", carID)

	var car models.Car
	err = row.Scan(
		&car.CarID,
		&car.TrimLevel,
		&car.Year,
		&car.VIN,
		&car.Price,
		&car.Color,
		&car.Bodywork,
		&car.Engine,
		&car.EngineCapacity,
		&car.Fuel,
		&car.Image,
		&car.Description1,
		&car.Description2,
	)

	if err != nil {
		log.Printf("Error scanning car row: %v", err)
		http.Error(w, "Failed to process car data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("Error encoding car to JSON: %v", err)
		http.Error(w, "Failed to encode car", http.StatusInternalServerError)
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
	INSERT INTO cars (trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING carid`
	id := 0
	err = db.DB.QueryRow(sqlStatement,
		car.TrimLevel,
		car.Year,
		car.VIN,
		car.Price,
		car.Color,
		car.Bodywork,
		car.Engine,
		car.EngineCapacity,
		car.Fuel,
		car.Image,
		car.Description1,
		car.Description2,
	).Scan(&id)
	if err != nil {
		log.Printf("Error creating car: %v", err)
		http.Error(w, "Failed to create car", http.StatusInternalServerError)
		return
	}

	car.CarID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}
