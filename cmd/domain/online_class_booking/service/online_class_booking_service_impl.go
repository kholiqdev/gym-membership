package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class_booking/dto"
	"gym/cmd/domain/online_class_booking/repository"
)

type OnlineClassBookingImpl struct {
	Repo repository.OnlineClassBookingRepository
}

func (s *OnlineClassBookingImpl) GetAll(ctx echo.Context) (*dto.OnlineClassBookingListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassBookingImpl) GetByInvoice(ctx echo.Context, invoice string) (*dto.OnlineClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassBookingImpl) Create(ctx echo.Context, request *dto.OnlineClassBookingCreateRequest) (*dto.OnlineClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}
