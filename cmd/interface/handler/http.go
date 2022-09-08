package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gym/cmd/interface/handler/admin"
	"gym/cmd/interface/handler/class"
	"gym/cmd/interface/handler/health"
	"gym/cmd/interface/handler/member"
	"gym/internal/protocol/http/middleware/auth"
)

type HttpHandlerImpl struct {
	member member.MemberHandler
	admin  admin.AdminHandler
	health health.HealthHandler
	class  class.ClassHandler
}

func NewHttpHandler(
	member member.MemberHandler,
	admin admin.AdminHandler,
	health health.HealthHandler,
	class class.ClassHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		member: member,
		admin:  admin,
		health: health,
		class:  class,
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
		memberGroup.POST("/join", h.member.Join, auth.JwtVerifyAccess("member"))
	}

	// Admin group
	adminGroup := e.Group("admin")
	{
		adminGroup.GET("/user/list", h.admin.GetMember, auth.JwtVerifyAccess("admin"))
		adminGroup.GET("/list", h.admin.Get, auth.JwtVerifyAccess("admin"))
		adminGroup.GET("/:id", h.admin.Detail, auth.JwtVerifyAccess("admin"))
		adminGroup.POST("/create", h.admin.Create)
		adminGroup.POST("/class/create", h.admin.CreateClass, auth.JwtVerifyAccess("admin"))
		adminGroup.POST("/member-type/create", h.admin.CreateMemberType, auth.JwtVerifyAccess("admin"))
		adminGroup.POST("/class-category/create", h.admin.CreateClassCategory, auth.JwtVerifyAccess("admin"))
		adminGroup.GET("/class-category", h.admin.GetClassCategory, auth.JwtVerifyAccess("admin"))
	}

	// Class group
	classGroup := e.Group("class")
	{
		classGroup.GET("", h.class.Get, auth.JwtVerifyAccess("member"))
		classGroup.GET("/:id", h.class.Detail, auth.JwtVerifyAccess("member"))
	}
}
