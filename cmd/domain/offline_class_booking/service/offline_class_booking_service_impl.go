package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class_booking/dto"
	"gym/cmd/domain/offline_class_booking/repository"
)

type OfflineClassBookingImpl struct {
	Repo repository.OfflineClassBookingRepository
}

func (s *OfflineClassBookingImpl) GetAll(ctx echo.Context) (*dto.OfflineClassBookingListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassBookingImpl) GetByInvoice(ctx echo.Context, invoice string) (*dto.OfflineClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassBookingImpl) Create(ctx echo.Context, request *dto.OfflineClassBookingCreateRequest) (*dto.OfflineClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}
