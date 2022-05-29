package dto

type OnlineClassCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quota       uint    `json:"quota" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Duration    string  `json:"duration" validate:"required"`
}

type OnlineClassUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quota       uint    `json:"quota"`
	Image       string  `json:"image"`
	Duration    string  `json:"duration"`
}
