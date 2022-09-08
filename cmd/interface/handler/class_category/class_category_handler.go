package class_category

import "github.com/labstack/echo/v4"

type ClassCategoryHandler interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
}
