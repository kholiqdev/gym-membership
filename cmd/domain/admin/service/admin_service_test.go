package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/admin/dto"
	"gym/cmd/domain/admin/entity"
	"gym/cmd/domain/admin/repository"
	"testing"
)

var adminRepository = &repository.AdminRepositoryMock{Mock: mock.Mock{}}
var adminService = AdminServiceImpl{RepoAdmin: adminRepository}

func TestAdminService_GetAdminByIdFail(t *testing.T) {

	// program mock
	adminRepository.Mock.On("Find", uint(1)).Return(nil, errors.New("admin not found")).Once()

	admin, err := adminService.GetAdminById(1)
	assert.Nil(t, admin)
	assert.NotNil(t, err)
}

func TestAdminService_GetAdminByIdSuccess(t *testing.T) {
	admin := entity.Admin{
		ID:       2,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	adminRepository.Mock.On("Find", admin.ID).Return(admin, nil).Once()

	result, err := adminService.GetAdminById(admin.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, admin.ID, result.ID)
	assert.Equal(t, admin.Name, result.Name)
}

func TestAdminService_GetAdminsFail(t *testing.T) {

	// program mock
	adminRepository.Mock.On("FindAll").Return(nil, errors.New("admins not found")).Once()

	admins, err := adminService.GetAdmins()
	assert.Nil(t, admins)
	assert.NotNil(t, err)
}

func TestAdminService_GetAdminsSuccess(t *testing.T) {
	admin := entity.Admin{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	adminList := entity.AdminList{&admin}

	// program mock
	adminRepository.Mock.On("FindAll").Return(adminList, nil).Once()

	admins, err := adminService.GetAdmins()

	assert.NotNil(t, admins)
	assert.Nil(t, err)
	assert.IsType(t, dto.AdminListResponse{}, *admins)
	assert.Equal(t, admin.ID, (*admins)[0].ID)
}

func TestAdminService_StoreFail(t *testing.T) {
	adminReq := dto.AdminCreateRequest{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	// program mock
	adminRepository.Mock.On("Insert", mock.Anything).Return(nil, errors.New("can't insert to db")).Once()

	result, err := adminService.Store(&adminReq)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestAdminService_StoreSuccess(t *testing.T) {
	admin := entity.Admin{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	adminReq := dto.AdminCreateRequest{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	// program mock
	adminRepository.Mock.On("Insert", mock.Anything).Return(admin, nil).Once()

	result, err := adminService.Store(&adminReq)

	assert.NotNil(t, *result)
	assert.Nil(t, err)
	assert.IsType(t, dto.AdminResponse{}, *result)
	assert.Equal(t, admin.ID, result.ID)
	assert.Equal(t, admin.Email, result.Email)
}
