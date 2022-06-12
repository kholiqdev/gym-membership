package dto

type ClassCategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type ClassCategoryUpdateRequest struct {
	Name string `json:"name"`
}
