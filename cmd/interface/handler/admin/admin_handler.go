package admin

import (
	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	Get(ctx echo.Context) error
	GetMember(ctx echo.Context) error
	GetClassCategory(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Create(ctx echo.Context) error
	CreateMemberType(ctx echo.Context) error
	CreateClass(ctx echo.Context) error
	CreateClassCategory(ctx echo.Context) error
	Login(ctx echo.Context) error
	Refresh(ctx echo.Context) error
}
