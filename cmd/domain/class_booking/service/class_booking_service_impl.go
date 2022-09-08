package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/dto"
	"gym/cmd/domain/class_booking/repository"
)

type ClassBookingImpl struct {
	Repo repository.ClassBookingRepository
}

func (s *ClassBookingImpl) GetAll(ctx echo.Context) (*dto.ClassBookingListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassBookingImpl) GetByInvoice(ctx echo.Context, invoice string) (*dto.ClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassBookingImpl) Create(ctx echo.Context, request *dto.ClassBookingCreateRequest) (*dto.ClassBookingResponse, error) {
	//TODO implement me
	panic("implement me")
}
