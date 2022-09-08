package service

import (
	"gym/cmd/domain/admin/dto"
)

type AdminService interface {
	GetAdmins() (*dto.AdminListResponse, error)
	GetAdminById(adminId uint) (*dto.AdminResponse, error)
	Store(request *dto.AdminCreateRequest) (*dto.AdminResponse, error)
	Login(request *dto.AdminLoginRequest) (*dto.AdminAuthResponse, error)
	Refresh(adminId uint) (*dto.AdminAuthResponse, error)
}
