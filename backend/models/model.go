package models

type Model struct {
	ModelID     int    `json:"model_id"`
	ModelName   string `json:"model_name"`
	Description string `json:"description"`
}
