package dto

import "time"

type ClassResponse struct {
	ClassCategory uint    `json:"class_category_id,omitempty"`
	TrainerName   string  `json:"trainer_name,omitempty"`
	ID            uint    `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price,omitempty"`
	MeetLink      string  `json:"meet_link,omitempty"`
	Image         string  `json:"image,omitempty"`
	Date          string  `json:"date,omitempty"`
	StartTime     string  `json:"start_time,omitempty"`
	EndTime       string  `json:"end_time,omitempty"`
	IsOffline     bool    `json:"is_offline"`
	CategoryName  string  `json:"category_name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ClassResponseList []*ClassResponse
