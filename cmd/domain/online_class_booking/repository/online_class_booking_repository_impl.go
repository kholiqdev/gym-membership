package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/online_class_booking/entity"
)

type OnlineClassBookingRepositoryImpl struct {
	Db *gorm.DB
}

func (r *OnlineClassBookingRepositoryImpl) FindAll(ctx echo.Context) (*entity.OnlineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassBookingRepositoryImpl) FindByInvoice(ctx echo.Context, invoice string) (*entity.OnlineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OnlineClassBookingRepositoryImpl) Insert(ctx echo.Context, entity *entity.OnlineClassBooking) (*entity.OnlineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}
