package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_category/entity"
	"gym/pkg/database"
)

type ClassCategoryRepository interface {
	FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassCategoryList, error)
	Insert(ctx echo.Context, trainer *entity.ClassCategory) (*entity.ClassCategory, error)
	Update(ctx echo.Context, trainer *entity.ClassCategory) (*entity.ClassCategory, error)
	Delete(ctx echo.Context, id uint) (*entity.ClassCategory, error)
}
