package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/trainer/dto"
)

type TrainerService interface {
	GetAll(ctx echo.Context) (*dto.TrainerListResponse, error)
	GetById(ctx echo.Context, id uint) (*dto.TrainerResponse, error)
	GetByEmail(ctx echo.Context, email string) (*dto.TrainerResponse, error)
	Create(ctx echo.Context, request *dto.TrainerCreateRequest) (*dto.TrainerResponse, error)
	Update(ctx echo.Context, trainer *dto.TrainerUpdateRequest) (*dto.TrainerResponse, error)
	Delete(ctx echo.Context, id uint) (*dto.TrainerResponse, error)
}
