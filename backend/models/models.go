package models

// Model представляет модель автомобиля.
type Model struct {
	ModelID     int64  `json:"model_id"`    // ID модели автомобиля
	ModelName   string `json:"model_name"`  // Название модели автомобиля
	Description string `json:"description"` // Описание модели автомобиля
}
