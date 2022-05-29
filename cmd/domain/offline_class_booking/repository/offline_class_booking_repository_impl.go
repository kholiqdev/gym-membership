package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/offline_class_booking/entity"
)

type OfflineClassBookingRepositoryImpl struct {
	Db *gorm.DB
}

func (r *OfflineClassBookingRepositoryImpl) FindAll(ctx echo.Context) (*entity.OfflineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassBookingRepositoryImpl) FindByInvoice(ctx echo.Context, trainerID uint) (*entity.OfflineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *OfflineClassBookingRepositoryImpl) Insert(ctx echo.Context, entity *entity.OfflineClassBooking) (*entity.OfflineClassBooking, error) {
	//TODO implement me
	panic("implement me")
}
