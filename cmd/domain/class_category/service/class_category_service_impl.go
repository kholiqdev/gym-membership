package service

import (
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gym/cmd/domain/class_category/dto"
	"gym/cmd/domain/class_category/entity"
	"gym/cmd/domain/class_category/repository"
	"gym/pkg/database"
)

type ClassCategoryServiceImpl struct {
	RepoClassCategory repository.ClassCategoryRepository
}

func (s *ClassCategoryServiceImpl) GetAll(ctx echo.Context, pagination *database.Pagination) (*dto.ClassCategoryListResponse, error) {
	classCategoryList, err := s.RepoClassCategory.FindAll(ctx, pagination)
	if err != nil {
		log.Err(err).Msg("Error fetch category list from DB")
		return nil, err
	}
	classCategoryListResp := dto.CreateClassCategoryListResponse(classCategoryList)
	return classCategoryListResp, nil
}

func (s *ClassCategoryServiceImpl) Create(ctx echo.Context, request *dto.ClassCategoryCreateRequest) (*dto.ClassCategoryResponse, error) {

	slug := slug.Make(request.Name)

	ClassCategory, err := s.RepoClassCategory.Insert(ctx, &entity.ClassCategory{
		Name: request.Name,
		Slug: slug,
	})

	if err != nil {
		log.Err(err).Msg("Error cant insert class_category to DB")
		return nil, err
	}
	ClassCategoryResp := dto.CreateClassCategoryResponse(ClassCategory)

	return ClassCategoryResp, nil
}

func (s *ClassCategoryServiceImpl) Update(ctx echo.Context, ClassCategory *dto.ClassCategoryUpdateRequest) (*dto.ClassCategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassCategoryServiceImpl) Delete(ctx echo.Context, id uint) (*dto.ClassCategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
