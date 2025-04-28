package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeHandler struct {
	DB *sql.DB
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT employeeid, firstname, lastname, email, phone, position, hiredate FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(
			&emp.ID,
			&emp.Firstname,
			&emp.Lastname,
			&emp.Email,
			&emp.Phone,
			&emp.Position,
			&emp.Hiredate,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var emp models.Employee
	row := h.DB.QueryRow("SELECT employeeid, firstname, lastname, email, phone, position, hiredate FROM employees WHERE employeeid = $1", id)
	err = row.Scan(
		&emp.ID,
		&emp.Firstname,
		&emp.Lastname,
		&emp.Email,
		&emp.Phone,
		&emp.Position,
		&emp.Hiredate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req models.EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Хеширование пароля
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	var emp models.Employee
	err = h.DB.QueryRow(
		`INSERT INTO employees 
		(firstname, lastname, email, phone, position, hiredate, passwordhash) 
		VALUES ($1, $2, $3, $4, $5, CURRENT_DATE, $6) 
		RETURNING employeeid, firstname, lastname, email, phone, position, hiredate`,
		req.Firstname,
		req.Lastname,
		req.Email,
		req.Phone,
		req.Position,
		hashedPassword,
	).Scan(
		&emp.ID,
		&emp.Firstname,
		&emp.Lastname,
		&emp.Email,
		&emp.Phone,
		&emp.Position,
		&emp.Hiredate,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var req models.EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновление данных
	_, err = h.DB.Exec(
		`UPDATE employees SET 
		firstname = $1,
		lastname = $2,
		email = $3,
		phone = $4,
		position = $5
		WHERE employeeid = $6`,
		req.Firstname,
		req.Lastname,
		req.Email,
		req.Phone,
		req.Position,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем обновленные данные
	var emp models.Employee
	err = h.DB.QueryRow(
		"SELECT employeeid, firstname, lastname, email, phone, position, hiredate FROM employees WHERE employeeid = $1",
		id,
	).Scan(
		&emp.ID,
		&emp.Firstname,
		&emp.Lastname,
		&emp.Email,
		&emp.Phone,
		&emp.Position,
		&emp.Hiredate,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	_, err = h.DB.Exec("DELETE FROM employees WHERE employeeid = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
