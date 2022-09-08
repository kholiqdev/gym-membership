package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/entity"
)

type ClassBookingRepository interface {
	FindAll(ctx echo.Context) (*entity.ClassBooking, error)
	FindByInvoice(ctx echo.Context, trainerID uint) (*entity.ClassBooking, error)
	Insert(ctx echo.Context, entity *entity.ClassBooking) (*entity.ClassBooking, error)
}
