package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class/dto"
	"gym/pkg/database"
)

type ClassService interface {
	GetAll(ctx echo.Context, pagination *database.Pagination) (*dto.ClassResponseList, error)
	GetById(ctx echo.Context, id uint) (*dto.ClassResponse, error)
	GetByTrainer(ctx echo.Context, trainerID uint) (*dto.ClassResponseList, error)
	Create(ctx echo.Context, request *dto.ClassCreateRequest) (*dto.ClassResponse, error)
	Update(ctx echo.Context, class *dto.ClassUpdateRequest) (*dto.ClassResponse, error)
	Delete(ctx echo.Context, id uint) (*dto.ClassResponse, error)
}
