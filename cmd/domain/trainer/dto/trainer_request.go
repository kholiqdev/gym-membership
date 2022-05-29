package dto

type TrainerCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Avatar   string `json:"avatar" validate:"required"`
}

type TrainerUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}
