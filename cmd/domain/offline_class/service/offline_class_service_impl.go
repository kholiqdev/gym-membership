package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class/dto"
)

type OfflineClassServiceImpl struct{}

func (s *OfflineClassServiceImpl) GetAll(ctx echo.Context) (*dto.OfflineClassResponseList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassServiceImpl) GetById(ctx echo.Context, id uint) (*dto.OfflineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassServiceImpl) GetByTrainer(ctx echo.Context, trainerID uint) (*dto.OfflineClassResponseList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassServiceImpl) Create(ctx echo.Context, request *dto.OfflineClassCreateRequest) (*dto.OfflineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassServiceImpl) Update(ctx echo.Context, offlineClass *dto.OfflineClassUpdateRequest) (*dto.OfflineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OfflineClassServiceImpl) Delete(ctx echo.Context, id uint) (*dto.OfflineClassResponse, error) {
	//TODO implement me
	panic("implement me")
}
