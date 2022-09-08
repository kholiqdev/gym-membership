package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/member_order/dto"
)

type MemberOrderService interface {
	GetAll(ctx echo.Context) (*dto.MemberOrderListResponse, error)
	GetByInvoice(ctx echo.Context, invoice string) (*dto.MemberOrderResponse, error)
	Create(ctx echo.Context, request *dto.MemberOrderCreateRequest) (*dto.MemberOrderResponse, error)
}
