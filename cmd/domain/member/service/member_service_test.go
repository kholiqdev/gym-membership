package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gym/cmd/domain/member/dto"
	"gym/cmd/domain/member/entity"
	"gym/cmd/domain/member/repository"
	"testing"
)

var memberRepository = &repository.MemberRepositoryMock{Mock: mock.Mock{}}
var memberService = MemberServiceImpl{RepoMember: memberRepository}

func TestMemberService_GetMemberByIdFail(t *testing.T) {

	// program mock
	memberRepository.Mock.On("Find", uint(1)).Return(nil, errors.New("member not found")).Once()

	member, err := memberService.GetMemberById(1)
	assert.Nil(t, member)
	assert.NotNil(t, err)
}

func TestMemberService_GetMemberByIdSuccess(t *testing.T) {
	member := entity.Member{
		ID:       2,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	memberRepository.Mock.On("Find", member.ID).Return(member, nil).Once()

	result, err := memberService.GetMemberById(member.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, member.ID, result.ID)
	assert.Equal(t, member.Name, result.Name)
}

func TestMemberService_GetMembersFail(t *testing.T) {

	// program mock
	memberRepository.Mock.On("FindAll").Return(nil, errors.New("members not found")).Once()

	members, err := memberService.GetMembers()
	assert.Nil(t, members)
	assert.NotNil(t, err)
}

func TestMemberService_GetMembersSuccess(t *testing.T) {
	member := entity.Member{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	memberList := entity.MemberList{&member}

	// program mock
	memberRepository.Mock.On("FindAll").Return(memberList, nil).Once()

	members, err := memberService.GetMembers()

	assert.NotNil(t, members)
	assert.Nil(t, err)
	assert.IsType(t, dto.MemberListResponse{}, *members)
	assert.Equal(t, member.ID, (*members)[0].ID)
}

func TestMemberService_StoreFail(t *testing.T) {
	memberReq := dto.MemberCreateRequest{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	// program mock
	memberRepository.Mock.On("Insert", mock.Anything).Return(nil, errors.New("can't insert to db")).Once()

	result, err := memberService.Store(&memberReq)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestMemberService_StoreSuccess(t *testing.T) {
	member := entity.Member{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	memberReq := dto.MemberCreateRequest{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	// program mock
	memberRepository.Mock.On("Insert", mock.Anything).Return(member, nil).Once()

	result, err := memberService.Store(&memberReq)

	assert.NotNil(t, *result)
	assert.Nil(t, err)
	assert.IsType(t, dto.MemberResponse{}, *result)
	assert.Equal(t, member.ID, result.ID)
	assert.Equal(t, member.Email, result.Email)
}
