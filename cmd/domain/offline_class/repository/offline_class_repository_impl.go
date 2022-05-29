package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/offline_class/entity"
)

type OfflineClassRepositoryImpl struct {
	Db *gorm.DB
}

func (r *OfflineClassRepositoryImpl) FindAll(ctx echo.Context) (*entity.OfflineClassList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassRepositoryImpl) Find(ctx echo.Context, id uint) (*entity.OfflineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassRepositoryImpl) FindByTrainer(ctx echo.Context, trainerID uint) (*entity.OfflineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassRepositoryImpl) Insert(ctx echo.Context, entity *entity.OfflineClass) (*entity.OfflineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassRepositoryImpl) Update(ctx echo.Context, entity *entity.OfflineClass) (*entity.OfflineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassRepositoryImpl) Delete(ctx echo.Context, id uint) (*entity.OfflineClass, error) {
	//TODO implement me
	panic("implement me")
}
