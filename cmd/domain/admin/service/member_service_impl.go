package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"gym/cmd/domain/admin/dto"
	"gym/cmd/domain/admin/entity"
	"gym/cmd/domain/admin/repository"
	"gym/internal/protocol/http/errors"
	"gym/pkg/auth"
	"gym/pkg/hash"
	"sync"
	"time"
)

type AdminServiceImpl struct {
	RepoAdmin repository.AdminRepository
	JwtAuth   auth.JwtToken
	repoOnce  sync.Once
}

func (s AdminServiceImpl) GetAdmins() (*dto.AdminListResponse, error) {
	admins, err := s.RepoAdmin.FindAll()
	if err != nil {
		log.Err(err).Msg("Error fetch admins from DB")
		return nil, err
	}
	adminsResp := dto.CreateAdminListResponse(admins)
	return &adminsResp, nil
}

func (s AdminServiceImpl) GetAdminById(adminId uint) (*dto.AdminResponse, error) {
	admin, err := s.RepoAdmin.Find(adminId)
	if err != nil {
		log.Err(err).Msg("Error fetch admin from DB")
		return nil, err
	}
	adminResp := dto.CreateAdminResponse(admin)
	return &adminResp, nil
}

func (s AdminServiceImpl) Store(request *dto.AdminCreateRequest) (*dto.AdminResponse, error) {
	passwordHashed, err := hash.AppBcryptImpl{}.HashAndSalt([]byte(request.Password))
	if err != nil {
		log.Err(err).Msg("Error hash password to bcrypt")
		return nil, err
	}

	adminRepo, err := s.RepoAdmin.Insert(&entity.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHashed,
	})

	if err != nil {
		log.Err(err).Msg("Error insert admin to DB")
		return nil, err
	}

	adminResp := dto.CreateAdminResponse(adminRepo)
	log.Info().Msg("Successfully insert to to DB")
	return &adminResp, nil
}

func (s AdminServiceImpl) Login(request *dto.AdminLoginRequest) (*dto.AdminAuthResponse, error) {
	admin, err := s.RepoAdmin.FindByEmail(request.Email)
	if err != nil {
		log.Err(err).Msg("Error fetch admin from DB")
		return nil, errors.FindErrorType(err)
	}

	isMatched := hash.AppBcryptImpl{}.ComparePasswords(admin.Password, []byte(request.Password))

	if !isMatched {
		log.Err(err).Msg("email and password didn't match")
		return nil, errors.Unauthorization("email and password didn't match")
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   admin.ID,
		"name": admin.Name,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateAdminAuthResponse(accessToken)

	return &authResp, nil
}

func (s AdminServiceImpl) Refresh(adminId uint) (*dto.AdminAuthResponse, error) {
	admin, err := s.RepoAdmin.Find(adminId)
	if err != nil {
		log.Err(err).Msg("Error fetch admin from DB")
		return nil, errors.FindErrorType(err)
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   admin.ID,
		"name": admin.Name,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateAdminAuthResponse(accessToken)

	return &authResp, nil
}
