package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class/dto"
)

type OnlineClassServiceImpl struct{}

func (s *OnlineClassServiceImpl) GetAll(ctx echo.Context) (*dto.OnlineClassResponseList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassServiceImpl) GetById(ctx echo.Context, id uint) (*dto.OnlineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassServiceImpl) Create(ctx echo.Context, request *dto.OnlineClassCreateRequest) (*dto.OnlineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassServiceImpl) Update(ctx echo.Context, onlineClass *dto.OnlineClassUpdateRequest) (*dto.OnlineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OnlineClassServiceImpl) Delete(ctx echo.Context, id uint) (*dto.OnlineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}
