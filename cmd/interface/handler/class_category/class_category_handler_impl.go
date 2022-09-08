package class_category

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_category/service"
)

type ClassCategoryHandlerImpl struct {
	SvcClassCategory service.ClassCategoryService
}

func (h *ClassCategoryHandlerImpl) Get(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassCategoryHandlerImpl) Create(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
