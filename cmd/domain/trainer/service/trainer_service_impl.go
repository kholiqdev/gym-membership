package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/trainer/dto"
	"gym/cmd/domain/trainer/repository"
)

type TrainerServiceImpl struct {
	Repo repository.TrainerRepository
}

func (s *TrainerServiceImpl) GetAll(ctx echo.Context) (*dto.TrainerListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TrainerServiceImpl) GetById(ctx echo.Context, id uint) (*dto.TrainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TrainerServiceImpl) GetByEmail(ctx echo.Context, email string) (*dto.TrainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TrainerServiceImpl) Create(ctx echo.Context, request *dto.TrainerCreateRequest) (*dto.TrainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TrainerServiceImpl) Update(ctx echo.Context, trainer *dto.TrainerUpdateRequest) (*dto.TrainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TrainerServiceImpl) Delete(ctx echo.Context, id uint) (*dto.TrainerResponse, error) {
	//TODO implement me
	panic("implement me")
}
