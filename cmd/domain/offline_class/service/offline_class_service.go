package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/offline_class/dto"
)

type OfflineClassService interface {
	GetAll(ctx echo.Context) (*dto.OfflineClassResponseList, error)
	GetById(ctx echo.Context, id uint) (*dto.OfflineClassResponse, error)
	GetByTrainer(ctx echo.Context, trainerID uint) (*dto.OfflineClassResponseList, error)
	Create(ctx echo.Context, request *dto.OfflineClassCreateRequest) (*dto.OfflineClassResponse, error)
	Update(ctx echo.Context, offlineClass *dto.OfflineClassUpdateRequest) (*dto.OfflineClassResponse, error)
	Delete(ctx echo.Context, id uint) (*dto.OfflineClassResponse, error)
}
