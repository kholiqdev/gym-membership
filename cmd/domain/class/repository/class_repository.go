package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class/entity"
	"gym/pkg/database"
)

type ClassRepository interface {
	FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassList, error)
	Find(ctx echo.Context, id uint) (*entity.Class, error)
	FindByTrainer(ctx echo.Context, trainerID uint) (*entity.Class, error)
	Insert(ctx echo.Context, entity *entity.Class) (*entity.Class, error)
	Update(ctx echo.Context, entity *entity.Class) (*entity.Class, error)
	Delete(ctx echo.Context, id uint) (*entity.Class, error)
}
