package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/online_class/dto"
)

type OnlineClassService interface {
	GetAll(ctx echo.Context) (*dto.OnlineClassResponseList, error)
	GetById(ctx echo.Context, id uint) (*dto.OnlineClassResponse, error)
	Create(ctx echo.Context, request *dto.OnlineClassCreateRequest) (*dto.OnlineClassResponse, error)
	Update(ctx echo.Context, onlineClass *dto.OnlineClassUpdateRequest) (*dto.OnlineClassResponse, error)
	Delete(ctx echo.Context, id uint) (*dto.OnlineClassResponse, error)
}
