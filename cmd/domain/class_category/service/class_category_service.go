package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_category/dto"
	"gym/pkg/database"
)

type ClassCategoryService interface {
	GetAll(ctx echo.Context, pagination *database.Pagination) (*dto.ClassCategoryListResponse, error)
	Create(ctx echo.Context, request *dto.ClassCategoryCreateRequest) (*dto.ClassCategoryResponse, error)
	Update(ctx echo.Context, trainer *dto.ClassCategoryUpdateRequest) (*dto.ClassCategoryResponse, error)
	Delete(ctx echo.Context, id uint) (*dto.ClassCategoryResponse, error)
}
