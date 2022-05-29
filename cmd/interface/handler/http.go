package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gym/cmd/interface/handler/health"
	"gym/cmd/interface/handler/member"
	"gym/internal/protocol/http/middleware/auth"
)

type HttpHandlerImpl struct {
	member member.MemberHandler
	health health.HealthHandler
}

func NewHttpHandler(
	member member.MemberHandler,
	health health.HealthHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		member: member,
		health: health,
	}
}

func (h *HttpHandlerImpl) RegisterPath(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", h.health.GetHealth)
	// Auth
	authGroup := e.Group("member/auth")
	{
		authGroup.POST("/login", h.member.Login)
		authGroup.POST("/register", h.member.Create)
		authGroup.POST("/refresh", h.member.Refresh, auth.JwtVerifyRefresh())
	}

	// Member group
	memberGroup := e.Group("member")
	{
		memberGroup.GET("", h.member.Get, auth.JwtVerifyAccess())
		memberGroup.GET("/:id", h.member.Detail)
	}
}
