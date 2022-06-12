package dto

type ClassCreateRequest struct {
	ClassCategory uint    `json:"class_category_id" validate:"required"`
	TrainerName   string  `json:"trainer_name" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	MeetLink      string  `json:"meet_link" validate:"required"`
	Image         string  `json:"image" validate:"required"`
	Date          string  `json:"date" validate:"required"`
	StartTime     string  `json:"start_time" validate:"required"`
	EndTime       string  `json:"end_time" validate:"required"`
	IsOffline     bool    `json:"is_offline" validate:"required"`
}

type ClassUpdateRequest struct {
	ClassCategory string  `json:"class_category_id"`
	TrainerName   string  `json:"trainer_name"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	MeetLink      string  `json:"meet_link"`
	Category      string  `json:"category"`
	Image         string  `json:"image"`
	Date          string  `json:"date"`
	StartTime     uint    `json:"start_time"`
	EndTime       string  `json:"end_time"`
	IsOffline     bool    `json:"is_offline"`
}
