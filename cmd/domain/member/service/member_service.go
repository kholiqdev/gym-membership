package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/member/dto"
)

type MemberService interface {
	GetMembers() (*dto.MemberListResponse, error)
	GetMemberById(memberId uint) (*dto.MemberResponse, error)
	Store(request *dto.MemberCreateRequest) (*dto.MemberResponse, error)
	StoreMemberType(request *dto.MemberTypeCreateRequest) (*dto.MemberTypeResponse, error)
	StoreMemberJoin(ctx echo.Context, request *dto.MemberJoinRequest) (*dto.MemberJoinResponse, error)
	Login(request *dto.MemberLoginRequest) (*dto.MemberAuthResponse, error)
	Refresh(memberId uint) (*dto.MemberAuthResponse, error)
}
