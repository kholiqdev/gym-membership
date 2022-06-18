package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog/log"
	"gym/cmd/domain/member/dto"
	"gym/cmd/domain/member/entity"
	"gym/cmd/domain/member/repository"
	"gym/internal/protocol/http/errors"
	"gym/pkg/auth"
	"gym/pkg/hash"
	"strings"
	"time"
)

type MemberServiceImpl struct {
	RepoMember repository.MemberRepository
	JwtAuth    auth.JwtToken
}

func (s MemberServiceImpl) GetMembers() (*dto.MemberListResponse, error) {
	members, err := s.RepoMember.FindAll()
	if err != nil {
		log.Err(err).Msg("Error fetch members from DB")
		return nil, err
	}
	membersResp := dto.CreateMemberListResponse(members)
	return &membersResp, nil
}

func (s MemberServiceImpl) GetMemberById(memberId uint) (*dto.MemberResponse, error) {
	member, err := s.RepoMember.Find(memberId)
	if err != nil {
		log.Err(err).Msg("Error fetch member from DB")
		return nil, err
	}
	memberResp := dto.CreateMemberResponse(member)
	return &memberResp, nil
}

func (s MemberServiceImpl) Store(request *dto.MemberCreateRequest) (*dto.MemberResponse, error) {
	passwordHashed, err := hash.AppBcryptImpl{}.HashAndSalt([]byte(request.Password))
	if err != nil {
		log.Err(err).Msg("Error hash password to bcrypt")
		return nil, err
	}

	memberRepo, err := s.RepoMember.Insert(&entity.Member{
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHashed,
	})

	if err != nil {
		log.Err(err).Msg("Error insert member to DB")
		return nil, err
	}

	memberResp := dto.CreateMemberResponse(memberRepo)
	log.Info().Msg("Successfully insert to to DB")
	return &memberResp, nil
}

func (s MemberServiceImpl) StoreMemberType(request *dto.MemberTypeCreateRequest) (*dto.MemberTypeResponse, error) {

	memberRepo, err := s.RepoMember.InsertMemberType(&entity.MemberType{
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		Duration:    request.Duration,
		Price:       request.Price,
	})

	if err != nil {
		log.Err(err).Msg("Error insert member type to DB")
		return nil, err
	}

	adminResp := dto.CreateMemberTypeResponse(memberRepo)
	log.Info().Msg("Successfully insert to to DB")
	return &adminResp, nil
}

func (s MemberServiceImpl) StoreMemberJoin(ctx echo.Context, request *dto.MemberJoinRequest) (*dto.MemberJoinResponse, error) {
	memberId := ctx.Get("user_id").(float64)
	memberType, err := s.RepoMember.FindMemberTypeById(uint(memberId))
	if err != nil {
		log.Err(err).Msg("member type not found")
		return nil, err
	}

	invoiceNo := fmt.Sprintf("INV-%s", strings.ToUpper(random.String(16)))
	memberJoin, err := s.RepoMember.InsertMemberJoin(&entity.MemberJoin{
		MemberTypeID:       request.MemberType,
		MemberID:           uint(memberId),
		InvoiceNo:          invoiceNo,
		StartAt:            request.StartAt,
		MemberName:         request.Name,
		MemberNik:          request.Nik,
		MemberPhone:        request.Phone,
		MemberEmail:        request.Email,
		MemberAddress:      request.Address,
		MemberCity:         request.City,
		MemberPostalCode:   request.PostalCode,
		MemberTypeName:     memberType.Name,
		MemberTypeDuration: memberType.Duration,
		MemberTypeImage:    memberType.Image,
		MemberTypePrice:    memberType.Price,
		PaymentMethod:      "BCA",
		Status:             1,
		Total:              memberType.Price,
	})

	if err != nil {
		log.Err(err).Msg("Error insert member type to DB")
		return nil, err
	}

	memberJoinResp := dto.CreateMemberJoinResponse(memberJoin)
	log.Info().Msg("Successfully insert to to DB")
	return &memberJoinResp, nil
}

func (s MemberServiceImpl) Login(request *dto.MemberLoginRequest) (*dto.MemberAuthResponse, error) {
	member, err := s.RepoMember.FindByEmail(request.Email)
	if err != nil {
		log.Err(err).Msg("Error fetch member from DB")
		return nil, errors.FindErrorType(err)
	}

	isMatched := hash.AppBcryptImpl{}.ComparePasswords(member.Password, []byte(request.Password))

	if !isMatched {
		log.Err(err).Msg("email and password didn't match")
		return nil, errors.Unauthorization("email and password didn't match")
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   member.ID,
		"name": member.Name,
		"role": "member",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateMemberAuthResponse(accessToken)

	return &authResp, nil
}

func (s MemberServiceImpl) Refresh(memberId uint) (*dto.MemberAuthResponse, error) {
	member, err := s.RepoMember.Find(memberId)
	if err != nil {
		log.Err(err).Msg("Error fetch member from DB")
		return nil, errors.FindErrorType(err)
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   member.ID,
		"name": member.Name,
		"role": "member",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateMemberAuthResponse(accessToken)

	return &authResp, nil
}
