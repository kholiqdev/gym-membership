package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class_booking/dto"
)

type OnlineClassBookingService interface {
	GetAll(ctx echo.Context) (*dto.OnlineClassBookingListResponse, error)
	GetByInvoice(ctx echo.Context, invoice string) (*dto.OnlineClassBookingResponse, error)
	Create(ctx echo.Context, request *dto.OnlineClassBookingCreateRequest) (*dto.OnlineClassBookingResponse, error)
}
