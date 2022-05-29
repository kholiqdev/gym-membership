package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class/entity"
)

type OfflineClassRepository interface {
	FindAll(ctx echo.Context) (*entity.OfflineClassList, error)
	Find(ctx echo.Context, id uint) (*entity.OfflineClass, error)
	FindByTrainer(ctx echo.Context, trainerID uint) (*entity.OfflineClass, error)
	Insert(ctx echo.Context, entity *entity.OfflineClass) (*entity.OfflineClass, error)
	Update(ctx echo.Context, entity *entity.OfflineClass) (*entity.OfflineClass, error)
	Delete(ctx echo.Context, id uint) (*entity.OfflineClass, error)
}
