package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class_booking/dto"
)

type OfflineClassBookingService interface {
	GetAll(ctx echo.Context) (*dto.OfflineClassBookingListResponse, error)
	GetByInvoice(ctx echo.Context, invoice string) (*dto.OfflineClassBookingResponse, error)
	Create(ctx echo.Context, request *dto.OfflineClassBookingCreateRequest) (*dto.OfflineClassBookingResponse, error)
}
