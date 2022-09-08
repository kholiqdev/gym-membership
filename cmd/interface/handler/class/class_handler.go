package class

import "github.com/labstack/echo/v4"

type ClassHandler interface {
	Get(ctx echo.Context) error
	Detail(ctx echo.Context) error
}
