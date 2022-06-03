package dto

type AdminCreateRequest struct {
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

type AdminLoginRequest struct {
	Email    string `json:"email"   validate:"required,email"`
	Password string `json:"password"   validate:"required"`
}
