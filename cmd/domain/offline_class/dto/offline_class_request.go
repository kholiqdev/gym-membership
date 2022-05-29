package dto

type OfflineClassCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quota       uint    `json:"quota" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Date        string  `json:"date" validate:"required"`
	Time        uint    `json:"time" validate:"required"`
	Duration    string  `json:"duration" validate:"required"`
	TrainerID   uint    `json:"trainer_id" validate:"required"`
}

type OfflineClassUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quota       uint    `json:"quota"`
	Image       string  `json:"image"`
	Date        string  `json:"date"`
	Time        uint    `json:"time"`
	Duration    string  `json:"duration"`
	TrainerID   uint    `json:"trainer_id"`
}
