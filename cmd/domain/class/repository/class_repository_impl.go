package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym/cmd/domain/class/entity"
	"gym/pkg/database"
)

type ClassRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassRepositoryImpl) FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassList, error) {
	var classes entity.ClassList

	if e := r.Db.Debug().Preload("ClassCategory").Scopes(database.Paginate(classes, pagination, r.Db)).Preload(clause.Associations).Find(&classes).Error; e != nil {
		return nil, e
	}

	return &classes, nil
}

func (r *ClassRepositoryImpl) Find(ctx echo.Context, id uint) (*entity.Class, error) {
	var class entity.Class
	if e := r.Db.Debug().First(&class, id).Error; e != nil {
		return nil, e
	}

	return &class, nil
}

func (r *ClassRepositoryImpl) FindByTrainer(ctx echo.Context, trainerID uint) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepositoryImpl) Insert(ctx echo.Context, class *entity.Class) (*entity.Class, error) {
	if e := r.Db.Debug().Create(&class).Error; e != nil {
		return nil, e
	}
	return class, nil
}

func (r *ClassRepositoryImpl) Update(ctx echo.Context, entity *entity.Class) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepositoryImpl) Delete(ctx echo.Context, id uint) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}
