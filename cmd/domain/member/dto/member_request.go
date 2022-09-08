package dto

type MemberCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type MemberTypeCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Duration    uint    `json:"duration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type MemberLoginRequest struct {
	Email    string `json:"email"   validate:"required,email"`
	Password string `json:"password"   validate:"required"`
}

type MemberJoinRequest struct {
	MemberType uint   `json:"member_type_id" validate:"required"`
	StartAt    string `json:"start_at" validate:"required"`
	Nik        string `json:"nik" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Address    string `json:"address" validate:"required"`
	City       string `json:"city" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
}
