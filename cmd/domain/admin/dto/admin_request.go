package dto

type AdminCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"   validate:"required,email"`
	Password string `json:"password"   validate:"required"`
}
