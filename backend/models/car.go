package models

type Car struct {
	CarID          int     `json:"carid"`
	TrimLevel      string  `json:"trim_level"`
	Year           int     `json:"year"`
	VIN            string  `json:"vin"`
	Price          float64 `json:"price"`
	Color          string  `json:"color"`
	Bodywork       string  `json:"bodywork"`
	Engine         float64 `json:"engine"` // Changed from int to float64 to match frontend
	EngineCapacity string  `json:"engine_capacity"`
	Fuel           string  `json:"fuel"`
	Image          string  `json:"image"`
	Description1   string  `json:"description_1"`
	Description2   string  `json:"description_2"`
}
