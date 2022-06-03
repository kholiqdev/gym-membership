package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/member/entity"
)

type MemberRepositoryMock struct {
	Mock mock.Mock
}

func (r *MemberRepositoryMock) InsertMemberType(member *entity.MemberType) (*entity.MemberType, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberRepositoryMock) FindAll() (*entity.MemberList, error) {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("not found arguments")
	} else {
		members := arguments.Get(0).(entity.MemberList)
		return &members, nil
	}
}

func (r *MemberRepositoryMock) Find(memberId uint) (*entity.Member, error) {
	arguments := r.Mock.Called(memberId)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found memberId")
	}
	member := arguments.Get(0).(entity.Member)
	return &member, nil
}

func (r *MemberRepositoryMock) FindByMembername(membername string) (*entity.Member, error) {
	arguments := r.Mock.Called(membername)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		member := arguments.Get(0).(entity.Member)
		return &member, nil
	}
}

func (r *MemberRepositoryMock) FindByEmail(email string) (*entity.Member, error) {
	arguments := r.Mock.Called(email)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		member := arguments.Get(0).(entity.Member)
		return &member, nil
	}
}

func (r *MemberRepositoryMock) Insert(member *entity.Member) (*entity.Member, error) {
	arguments := r.Mock.Called(member)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		member := arguments.Get(0).(entity.Member)
		return &member, nil
	}
}
