package dto

import (
	"gym/cmd/domain/member/entity"
	"gym/pkg/auth/dto"
	"time"
)

type MemberResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MemberAuthResponse struct {
	Type         string `json:"type,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type MemberListResponse []*MemberResponse

func CreateMemberResponse(member *entity.Member) MemberResponse {
	return MemberResponse{
		ID:        member.ID,
		Name:      member.Name,
		Phone:     member.Phone,
		Email:     member.Email,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}
}

func CreateMemberListResponse(members *entity.MemberList) MemberListResponse {
	memberResp := MemberListResponse{}
	for _, p := range *members {
		member := CreateMemberResponse(p)
		memberResp = append(memberResp, &member)
	}
	return memberResp
}

func CreateMemberAuthResponse(token dto.AccessToken) MemberAuthResponse {
	return MemberAuthResponse{
		Type:         token.Type,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
	}
}
