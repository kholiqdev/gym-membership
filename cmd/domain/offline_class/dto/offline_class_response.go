package dto

import "time"

type OfflineClassResponse struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quota       uint    `json:"quota,omitempty"`
	Image       string  `json:"image,omitempty"`
	Date        string  `json:"date,omitempty"`
	Time        string  `json:"time,omitempty"`
	Duration    uint    `json:"duration,omitempty"`
	TrainerID   uint    `json:"trainer_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type OfflineClassResponseList []*OfflineClassResponse
