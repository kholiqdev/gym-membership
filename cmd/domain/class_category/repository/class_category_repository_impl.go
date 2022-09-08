package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym/cmd/domain/class_category/entity"
	"gym/pkg/database"
)

type ClassCategoryRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassCategoryRepositoryImpl) FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassCategoryList, error) {
	var classCategoryList entity.ClassCategoryList

	if e := r.Db.Debug().Scopes(database.Paginate(classCategoryList, pagination, r.Db)).Preload(clause.Associations).Find(&classCategoryList).Error; e != nil {
		return nil, e
	}

	return &classCategoryList, nil
}

func (r *ClassCategoryRepositoryImpl) Insert(ctx echo.Context, classCategory *entity.ClassCategory) (*entity.ClassCategory, error) {
	if e := r.Db.Debug().Create(&classCategory).Error; e != nil {
		return nil, e
	}
	return classCategory, nil
}

func (r *ClassCategoryRepositoryImpl) Update(ctx echo.Context, classCategory *entity.ClassCategory) (*entity.ClassCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassCategoryRepositoryImpl) Delete(ctx echo.Context, id uint) (*entity.ClassCategory, error) {
	//TODO implement me
	panic("implement me")
}
