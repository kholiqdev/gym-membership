package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/trainer/entity"
)

type TrainerRepository interface {
	FindAll(ctx echo.Context) (*entity.TrainerList, error)
	Find(ctx echo.Context, TrainerID uint) (*entity.Trainer, error)
	FindByEmail(ctx echo.Context, email string) (*entity.Trainer, error)
	Insert(ctx echo.Context, user *entity.Trainer) (*entity.Trainer, error)
	Update(ctx echo.Context, user *entity.Trainer) (*entity.Trainer, error)
	Delete(ctx echo.Context, TrainerID uint) (*entity.Trainer, error)
}
