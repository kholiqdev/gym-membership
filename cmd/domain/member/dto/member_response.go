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

type MemberTypeResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Duration    uint      `json:"image,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type MemberAuthResponse struct {
	Type         string `json:"type,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type MemberJoinResponse struct {
	ID                 uint    `json:"id,omitempty"`
	InvoiceNo          string  `json:"invoice_no,omitempty"`
	StartAt            string  `json:"start_at,omitempty"`
	MemberName         string  `json:"member_name,omitempty"`
	MemberNik          string  `json:"member_nik,omitempty"`
	MemberGender       string  `json:"member_gender,omitempty"`
	MemberEmail        string  `json:"member_email,omitempty"`
	MemberPhone        string  `json:"member_phone,omitempty"`
	MemberAddress      string  `json:"member_address,omitempty"`
	MemberCity         string  `json:"member_city,omitempty"`
	MemberPostalCode   string  `json:"member_postal_code,omitempty"`
	MemberTypeName     string  `json:"member_type_name,omitempty"`
	MemberTypeImage    string  `json:"member_type_image,omitempty"`
	MemberTypeDuration uint    `json:"member_type_duration,omitempty"`
	MemberTypePrice    float64 `json:"member_type_price,omitempty"`
	PaymentMethod      string  `json:"payment_method,omitempty"`
	Status             uint    `json:"status,omitempty"`
	Total              float64 `json:"total,omitempty"`
}

type MemberListResponse []*MemberResponse
