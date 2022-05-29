package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/trainer/entity"
)

type TrainerRepositoryImpl struct {
	Db *gorm.DB
}

func (r *TrainerRepositoryImpl) FindAll(ctx echo.Context) (*entity.TrainerList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TrainerRepositoryImpl) Find(ctx echo.Context, TrainerID uint) (*entity.Trainer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TrainerRepositoryImpl) FindByEmail(ctx echo.Context, email string) (*entity.Trainer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TrainerRepositoryImpl) Insert(ctx echo.Context, user *entity.Trainer) (*entity.Trainer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TrainerRepositoryImpl) Update(ctx echo.Context, user *entity.Trainer) (*entity.Trainer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TrainerRepositoryImpl) Delete(ctx echo.Context, TrainerID uint) (*entity.Trainer, error) {
	//TODO implement me
	panic("implement me")
}
