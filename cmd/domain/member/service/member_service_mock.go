package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/member/dto"
)

type MemberServiceMock struct {
	Mock mock.Mock
}

func (u MemberServiceMock) GetMembers() (*dto.MemberListResponse, error) {
	arguments := u.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		members := arguments.Get(0).(dto.MemberListResponse)
		return &members, nil
	}
}

func (u MemberServiceMock) GetMemberById(memberId uint) (*dto.MemberResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u MemberServiceMock) Store(request *dto.MemberCreateRequest) (*dto.MemberResponse, error) {
	arguments := u.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		member := arguments.Get(0).(dto.MemberResponse)
		return &member, nil
	}
}

func (u MemberServiceMock) Login(request *dto.MemberLoginRequest) (*dto.MemberAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u MemberServiceMock) Refresh(memberId uint) (*dto.MemberAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}
