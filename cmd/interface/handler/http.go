package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gym/cmd/interface/handler/admin"
	"gym/cmd/interface/handler/health"
	"gym/cmd/interface/handler/member"
	"gym/internal/protocol/http/middleware/auth"
)

type HttpHandlerImpl struct {
	member member.MemberHandler
	admin  admin.AdminHandler
	health health.HealthHandler
}

func NewHttpHandler(
	member member.MemberHandler,
	admin admin.AdminHandler,
	health health.HealthHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		member: member,
		admin:  admin,
		health: health,
	}
}

func (h *HttpHandlerImpl) RegisterPath(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", h.health.GetHealth)
	// Auth Member
	memberAuthGroup := e.Group("member/auth")
	{
		memberAuthGroup.POST("/login", h.member.Login)
		memberAuthGroup.POST("/register", h.member.Create)
		memberAuthGroup.POST("/refresh", h.member.Refresh, auth.JwtVerifyRefresh("member"))
	}

	// Auth Admin
	adminAuthGroup := e.Group("admin/auth")
	{
		adminAuthGroup.POST("/login", h.admin.Login)
		adminAuthGroup.POST("/create", h.admin.Create)
		adminAuthGroup.POST("/refresh", h.admin.Refresh, auth.JwtVerifyRefresh("admin"))
	}

	// Member group
	memberGroup := e.Group("member")
	{
		memberGroup.GET("", h.member.Get, auth.JwtVerifyAccess("member"))
		memberGroup.GET("/:id", h.member.Detail)
	}

	// Admin group
	adminGroup := e.Group("admin")
	{
		adminGroup.GET("/list", h.admin.Get, auth.JwtVerifyAccess("admin"))
		adminGroup.GET("/:id", h.admin.Detail, auth.JwtVerifyAccess("admin"))
		adminGroup.POST("/create", h.admin.Create)
		adminGroup.POST("/member-type/create", h.admin.CreateMemberType, auth.JwtVerifyAccess("admin"))
	}

}
