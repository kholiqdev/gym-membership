package dto

import (
	"gym/cmd/domain/admin/entity"
	"gym/pkg/auth/dto"
	"time"
)

type AdminResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AdminAuthResponse struct {
	Type         string `json:"type,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AdminListResponse []*AdminResponse

func CreateAdminResponse(member *entity.Admin) AdminResponse {
	return AdminResponse{
		ID:        member.ID,
		Name:      member.Name,
		Phone:     member.Phone,
		Email:     member.Email,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}
}

func CreateAdminListResponse(members *entity.AdminList) AdminListResponse {
	memberResp := AdminListResponse{}
	for _, p := range *members {
		member := CreateAdminResponse(p)
		memberResp = append(memberResp, &member)
	}
	return memberResp
}

func CreateAdminAuthResponse(token dto.AccessToken) AdminAuthResponse {
	return AdminAuthResponse{
		Type:         token.Type,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
	}
}
