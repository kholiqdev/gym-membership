package dto

import (
	"time"
)

type MemberResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MemberAuthResponse struct {
	Type         string `json:"type,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type MemberListResponse []*MemberResponse
