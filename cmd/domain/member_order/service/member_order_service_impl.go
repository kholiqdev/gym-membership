package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/member_order/dto"
	"gym/cmd/domain/member_order/repository"
)

type MemberOrderImpl struct {
	Repo repository.MemberOrderRepository
}

func (s *MemberOrderImpl) GetAll(ctx echo.Context) (*dto.MemberOrderListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MemberOrderImpl) GetByInvoice(ctx echo.Context, invoice string) (*dto.MemberOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MemberOrderImpl) Create(ctx echo.Context, request *dto.MemberOrderCreateRequest) (*dto.MemberOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}
