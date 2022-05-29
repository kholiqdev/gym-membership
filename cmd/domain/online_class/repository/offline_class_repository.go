package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class/entity"
)

type OnlineClassRepository interface {
	FindAll(ctx echo.Context) (*entity.OnlineClassList, error)
	Find(ctx echo.Context, id uint) (*entity.OnlineClass, error)
	Insert(ctx echo.Context, entity *entity.OnlineClass) (*entity.OnlineClass, error)
	Update(ctx echo.Context, entity *entity.OnlineClass) (*entity.OnlineClass, error)
	Delete(ctx echo.Context, id uint) (*entity.OnlineClass, error)
}
