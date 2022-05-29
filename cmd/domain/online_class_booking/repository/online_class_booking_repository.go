package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class_booking/entity"
)

type OnlineClassBookingRepository interface {
	FindAll(ctx echo.Context) (*entity.OnlineClassBooking, error)
	FindByInvoice(ctx echo.Context, invoice string) (*entity.OnlineClassBooking, error)
	Insert(ctx echo.Context, entity *entity.OnlineClassBooking) (*entity.OnlineClassBooking, error)
}
