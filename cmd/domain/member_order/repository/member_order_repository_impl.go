package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/member_order/entity"
)

type MemberOrderRepositoryImpl struct {
	Db *gorm.DB
}

func (r *MemberOrderRepositoryImpl) FindAll(ctx echo.Context) (*entity.MemberOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberOrderRepositoryImpl) FindByInvoice(ctx echo.Context, trainerID uint) (*entity.MemberOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberOrderRepositoryImpl) Insert(ctx echo.Context, entity *entity.MemberOrder) (*entity.MemberOrder, error) {
	//TODO implement me
	panic("implement me")
}
