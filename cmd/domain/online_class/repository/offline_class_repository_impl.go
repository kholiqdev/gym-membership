package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/online_class/entity"
)

type OnlineClassRepositoryImpl struct {
	Db *gorm.DB
}

func (r *OnlineClassRepositoryImpl) FindAll(ctx echo.Context) (*entity.OnlineClassList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassRepositoryImpl) Find(ctx echo.Context, id uint) (*entity.OnlineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassRepositoryImpl) Insert(ctx echo.Context, entity *entity.OnlineClass) (*entity.OnlineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassRepositoryImpl) Update(ctx echo.Context, entity *entity.OnlineClass) (*entity.OnlineClass, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassRepositoryImpl) Delete(ctx echo.Context, id uint) (*entity.OnlineClass, error) {
	//TODO implement me
	panic("implement me")
}
