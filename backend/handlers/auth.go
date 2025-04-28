package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	DB          *sql.DB
	JWTSecret   string
	TokenExpiry time.Duration
}

type AuthResponse struct {
	Token    string          `json:"token"`
	Employee models.Employee `json:"employee"`
}

// RegisterRequest - структура для запроса регистрации сотрудника
type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	Position  string `json:"position" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
}

// LoginRequest - структура для запроса входа
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверка существования сотрудника
	var exists bool
	err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE email = $1)", req.Email).Scan(&exists)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Employee with this email already exists", http.StatusConflict)
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Создание сотрудника
	var employee models.Employee
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
		string(hashedPassword),
	).Scan(
		&employee.ID,
		&employee.Firstname,
		&employee.Lastname,
		&employee.Email,
		&employee.Phone,
		&employee.Position,
		&employee.Hiredate,
	)

	if err != nil {
		log.Printf("Employee creation error: %v", err)
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	// Генерация JWT токена
	token, err := h.generateToken(employee)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Token:    token,
		Employee: employee,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("JSON encoding error: %v", err)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Поиск сотрудника
	var employee models.Employee
	var hashedPassword string
	err := h.DB.QueryRow(
		`SELECT employeeid, firstname, lastname, email, phone, position, hiredate, passwordhash 
		FROM employees WHERE email = $1`,
		req.Email,
	).Scan(
		&employee.ID,
		&employee.Firstname,
		&employee.Lastname,
		&employee.Email,
		&employee.Phone,
		&employee.Position,
		&employee.Hiredate,
		&hashedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Генерация JWT токена
	token, err := h.generateToken(employee)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Token:    token,
		Employee: employee,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("JSON encoding error: %v", err)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// В JWT logout обычно реализуется на клиенте
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully logged out. Please remove the token from client storage.",
	})
}

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(h.JWTSecret), nil
	})

	if err != nil {
		log.Printf("Token validation error: %v", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	employeeID, ok := claims["employee_id"].(float64)
	if !ok {
		http.Error(w, "Invalid employee ID in token", http.StatusUnauthorized)
		return
	}

	var employee models.Employee
	err = h.DB.QueryRow(
		`SELECT employeeid, firstname, lastname, email, phone, position, hiredate 
		FROM employees WHERE employeeid = $1`,
		int(employeeID),
	).Scan(
		&employee.ID,
		&employee.Firstname,
		&employee.Lastname,
		&employee.Email,
		&employee.Phone,
		&employee.Position,
		&employee.Hiredate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(employee); err != nil {
		log.Printf("JSON encoding error: %v", err)
	}
}

func (h *AuthHandler) generateToken(employee models.Employee) (string, error) {
	claims := jwt.MapClaims{
		"employee_id": employee.ID,
		"email":       employee.Email,
		"position":    employee.Position,
		"exp":         time.Now().Add(h.TokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.JWTSecret))
}
