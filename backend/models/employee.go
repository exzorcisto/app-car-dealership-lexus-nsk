package models

import (
	"time"
)

// Employee - основная модель сотрудника
type Employee struct {
	ID           int       `json:"id"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Position     string    `json:"position"`
	Hiredate     time.Time `json:"hiredate"`
	PasswordHash string    `json:"-"` // Исключаем из JSON
}

// EmployeeRequest - для запросов создания/обновления сотрудника
type EmployeeRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	Position  string `json:"position" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"` // Только для создания
}
