package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/member_order/entity"
)

type MemberOrderRepository interface {
	FindAll(ctx echo.Context) (*entity.MemberOrder, error)
	FindByInvoice(ctx echo.Context, trainerID uint) (*entity.MemberOrder, error)
	Insert(ctx echo.Context, entity *entity.MemberOrder) (*entity.MemberOrder, error)
}
