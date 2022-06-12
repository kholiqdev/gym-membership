package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gym/cmd/domain/class/dto"
	"gym/cmd/domain/class/entity"
	"gym/cmd/domain/class/repository"
	"gym/pkg/database"
	"gym/pkg/nullstring"
)

type ClassServiceImpl struct {
	RepoClass repository.ClassRepository
}

func (s *ClassServiceImpl) GetAll(ctx echo.Context, pagination *database.Pagination) (*dto.ClassResponseList, error) {
	classes, err := s.RepoClass.FindAll(ctx, pagination)
	if err != nil {
		log.Err(err).Msg("Error fetch classes from DB")
		return nil, err
	}
	classesResp := dto.CreateClassListResponse(classes)
	return classesResp, nil
}

func (s *ClassServiceImpl) GetById(ctx echo.Context, id uint) (*dto.ClassResponse, error) {
	class, err := s.RepoClass.Find(ctx, id)
	if err != nil {
		log.Err(err).Msg("Error fetch class from DB")
		return nil, err
	}
	classResp := dto.CreateClassResponse(class)
	return classResp, nil
}

func (s *ClassServiceImpl) GetByTrainer(ctx echo.Context, trainerID uint) (*dto.ClassResponseList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassServiceImpl) Create(ctx echo.Context, request *dto.ClassCreateRequest) (*dto.ClassResponse, error) {

	class, err := s.RepoClass.Insert(ctx, &entity.Class{
		ClassCategoryID: request.ClassCategory,
		TrainerName:     request.TrainerName,
		Name:            request.Name,
		Description:     request.Description,
		Price:           request.Price,
		MeetLink:        nullstring.NewNullString(request.MeetLink),
		Image:           request.Image,
		Date:            request.Date,
		StartTime:       request.StartTime,
		EndTime:         request.EndTime,
		IsOffline:       request.IsOffline,
	})
	if err != nil {
		log.Err(err).Msg("Error cant insert class to DB")
		return nil, err
	}
	classResp := dto.CreateClassResponse(class)
	return classResp, nil
}

func (s *ClassServiceImpl) Update(ctx echo.Context, class *dto.ClassUpdateRequest) (*dto.ClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassServiceImpl) Delete(ctx echo.Context, id uint) (*dto.ClassResponse, error) {
	//TODO implement me
	panic("implement me")
}
