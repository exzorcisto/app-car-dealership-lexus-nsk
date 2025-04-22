package models

type Car struct {
	CarID     int     `json:"car_id"`
	ModelID   int     `json:"model_id"`
	TrimLevel string  `json:"trim_level"`
	Year      int     `json:"year"`
	VIN       string  `json:"vin"`
	Price     float64 `json:"price"`
	Color     string  `json:"color"`
	Mileage   int     `json:"mileage"`
	IsNew     bool    `json:"is_new"`
}
