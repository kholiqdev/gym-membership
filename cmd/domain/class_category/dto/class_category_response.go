package dto

type ClassCategoryResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type ClassCategoryListResponse []*ClassCategoryResponse
