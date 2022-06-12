package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/admin/dto"
)

type AdminServiceMock struct {
	Mock mock.Mock
}

func (u AdminServiceMock) GetAdmins() (*dto.AdminListResponse, error) {
	arguments := u.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		admins := arguments.Get(0).(dto.AdminListResponse)
		return &admins, nil
	}
}

func (u AdminServiceMock) GetAdminById(adminId uint) (*dto.AdminResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u AdminServiceMock) Store(request *dto.AdminCreateRequest) (*dto.AdminResponse, error) {
	arguments := u.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		admin := arguments.Get(0).(dto.AdminResponse)
		return &admin, nil
	}
}

func (u AdminServiceMock) Login(request *dto.AdminLoginRequest) (*dto.AdminAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u AdminServiceMock) Refresh(adminId uint) (*dto.AdminAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}
