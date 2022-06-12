package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/class_booking/entity"
)

type ClassBookingRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassBookingRepositoryImpl) FindAll(ctx echo.Context) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassBookingRepositoryImpl) FindByInvoice(ctx echo.Context, trainerID uint) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassBookingRepositoryImpl) Insert(ctx echo.Context, entity *entity.ClassBooking) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}
