package repository

import (
	"gorm.io/gorm"
	"gym/cmd/domain/admin/entity"
)

type AdminRepositoryImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	Db *gorm.DB
}

func (r AdminRepositoryImpl) FindAll() (*entity.AdminList, error) {
	var admins entity.AdminList

	if e := r.Db.Debug().Find(&admins).Error; e != nil {
		return nil, e
	}

	return &admins, nil
}

func (r AdminRepositoryImpl) Find(adminId uint) (*entity.Admin, error) {
	var admin entity.Admin

	if e := r.Db.Debug().First(&admin, adminId).Error; e != nil {
		return nil, e
	}

	return &entity.Admin{
		ID:       admin.ID,
		Name:     admin.Name,
		Phone:    admin.Phone,
		Email:    admin.Email,
		Password: admin.Password,
	}, nil
}

func (r AdminRepositoryImpl) FindByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin

	if e := r.Db.Debug().Where("email = ?", email).First(&admin).Error; e != nil {
		return nil, e
	}

	return &entity.Admin{
		ID:       admin.ID,
		Name:     admin.Name,
		Phone:    admin.Phone,
		Email:    admin.Email,
		Password: admin.Password,
	}, nil
}

func (r AdminRepositoryImpl) Insert(admin *entity.Admin) (*entity.Admin, error) {
	if e := r.Db.Debug().Create(&admin).Error; e != nil {
		return nil, e
	}
	return admin, nil
}
