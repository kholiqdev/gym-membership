package service

import (
	"gym/cmd/domain/member/dto"
)

type MemberService interface {
	GetMembers() (*dto.MemberListResponse, error)
	GetMemberById(memberId uint) (*dto.MemberResponse, error)
	Store(request *dto.MemberCreateRequest) (*dto.MemberResponse, error)
	Login(request *dto.MemberLoginRequest) (*dto.MemberAuthResponse, error)
	Refresh(memberId uint) (*dto.MemberAuthResponse, error)
}
