package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gorilla/mux"
)

type CarHandler struct {
	DB *sql.DB
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT carid, model_name, trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2 FROM cars")
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
			&car.ModelName,
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
	}
}

func (h *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	row := h.DB.QueryRow("SELECT carid, model_name, trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2 FROM cars WHERE carid = $1", carID)

	var car models.Car
	if err := row.Scan(
		&car.CarID,
		&car.ModelName,
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
		if err == sql.ErrNoRows {
			http.Error(w, "Car not found", http.StatusNotFound)
		} else {
			log.Printf("Error scanning car row: %v", err)
			http.Error(w, "Failed to process car data", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("Error encoding car to JSON: %v", err)
		http.Error(w, "Failed to encode car", http.StatusInternalServerError)
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация обязательных полей
	if car.ModelName == "" || car.VIN == "" {
		http.Error(w, "Model name and VIN are required", http.StatusBadRequest)
		return
	}

	sqlStatement := `
	INSERT INTO cars (model_name, trimlevel, year, vin, price, color, bodywork, engine, engine_capacity, fuel, image, description_1, description_2)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING carid`

	err := h.DB.QueryRow(sqlStatement,
		car.ModelName,
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
	).Scan(&car.CarID)

	if err != nil {
		log.Printf("Error creating car: %v", err)
		http.Error(w, "Failed to create car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("Error encoding car to JSON: %v", err)
	}
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверка существования автомобиля
	var exists bool
	err = h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM cars WHERE carid = $1)", carID).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	_, err = h.DB.Exec(
		`UPDATE cars SET 
		model_name = $1,
		trimlevel = $2,
		year = $3,
		vin = $4,
		price = $5,
		color = $6,
		bodywork = $7,
		engine = $8,
		engine_capacity = $9,
		fuel = $10,
		image = $11,
		description_1 = $12,
		description_2 = $13
		WHERE carid = $14`,
		car.ModelName,
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
		carID,
	)
	if err != nil {
		log.Printf("Error updating car: %v", err)
		http.Error(w, "Failed to update car", http.StatusInternalServerError)
		return
	}

	car.CarID = carID
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("Error encoding car to JSON: %v", err)
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	// Проверка существования автомобиля
	var exists bool
	err = h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM cars WHERE carid = $1)", carID).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	_, err = h.DB.Exec("DELETE FROM cars WHERE carid = $1", carID)
	if err != nil {
		log.Printf("Error deleting car: %v", err)
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
