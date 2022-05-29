package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/admin/entity"
)

type AdminRepositoryMock struct {
	Mock mock.Mock
}

func (r *AdminRepositoryMock) FindAll() (*entity.AdminList, error) {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("not found arguments")
	} else {
		admins := arguments.Get(0).(entity.AdminList)
		return &admins, nil
	}
}

func (r *AdminRepositoryMock) Find(adminId uint) (*entity.Admin, error) {
	arguments := r.Mock.Called(adminId)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found adminId")
	}
	admin := arguments.Get(0).(entity.Admin)
	return &admin, nil
}

func (r *AdminRepositoryMock) FindByAdminname(adminname string) (*entity.Admin, error) {
	arguments := r.Mock.Called(adminname)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		admin := arguments.Get(0).(entity.Admin)
		return &admin, nil
	}
}

func (r *AdminRepositoryMock) FindByEmail(email string) (*entity.Admin, error) {
	arguments := r.Mock.Called(email)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		admin := arguments.Get(0).(entity.Admin)
		return &admin, nil
	}
}

func (r *AdminRepositoryMock) Insert(admin *entity.Admin) (*entity.Admin, error) {
	arguments := r.Mock.Called(admin)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		admin := arguments.Get(0).(entity.Admin)
		return &admin, nil
	}
}
