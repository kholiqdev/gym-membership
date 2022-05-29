package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class_booking/entity"
)

type OfflineClassBookingRepository interface {
	FindAll(ctx echo.Context) (*entity.OfflineClassBooking, error)
	FindByInvoice(ctx echo.Context, trainerID uint) (*entity.OfflineClassBooking, error)
	Insert(ctx echo.Context, entity *entity.OfflineClassBooking) (*entity.OfflineClassBooking, error)
}
