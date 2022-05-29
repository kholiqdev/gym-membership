package dto

import "time"

type OnlineClassResponse struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quota       uint    `json:"quota,omitempty"`
	Image       string  `json:"image,omitempty"`
	Duration    uint    `json:"duration,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type OnlineClassResponseList []*OnlineClassResponse
