package repository

import (
	"gym/cmd/domain/admin/entity"
)

type AdminRepository interface {
	FindAll() (*entity.AdminList, error)
	Find(adminId uint) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Insert(admin *entity.Admin) (*entity.Admin, error)
}
