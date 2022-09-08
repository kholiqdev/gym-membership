package admin

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/admin/dto"
	"gym/cmd/domain/admin/service"
	_dtoClass "gym/cmd/domain/class/dto"
	_svcClass "gym/cmd/domain/class/service"
	_dtoClassCategory "gym/cmd/domain/class_category/dto"
	_svcClassCategory "gym/cmd/domain/class_category/service"
	_dtoMember "gym/cmd/domain/member/dto"
	_svcMember "gym/cmd/domain/member/service"
	"gym/internal/protocol/http/response"
	"gym/pkg/database"
	"net/http"
	"strconv"
)

type AdminHandlerImpl struct {
	SvcAdmin         service.AdminService
	SvcMember        _svcMember.MemberService
	SvcClass         _svcClass.ClassService
	SvcClassCategory _svcClassCategory.ClassCategoryService
}

func (h AdminHandlerImpl) Get(ctx echo.Context) error {
	admins, err := h.SvcAdmin.GetAdmins()

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"admins": admins,
	})
	return nil
}

func (h AdminHandlerImpl) GetMember(ctx echo.Context) error {
	members, err := h.SvcMember.GetMembers()

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"members": members,
	})
	return nil
}

func (h AdminHandlerImpl) GetClassCategory(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)
	classCategory, err := h.SvcClassCategory.GetAll(ctx, pagination)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"class_category": map[string]interface{}{
			"data":       classCategory,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h AdminHandlerImpl) Detail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	admin, err := h.SvcAdmin.GetAdminById(uint(id))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", admin)
	return nil
}

func (h AdminHandlerImpl) Create(ctx echo.Context) error {
	var adminDto dto.AdminCreateRequest

	if err := ctx.Bind(&adminDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	admin, err := h.SvcAdmin.Store(&adminDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", admin)
	return nil
}

func (h AdminHandlerImpl) CreateMemberType(ctx echo.Context) error {
	var memberTypeCreateDto _dtoMember.MemberTypeCreateRequest

	if err := ctx.Bind(&memberTypeCreateDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	memberType, err := h.SvcMember.StoreMemberType(&memberTypeCreateDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", memberType)
	return nil
}

func (h AdminHandlerImpl) CreateClass(ctx echo.Context) error {
	var classCreateDto _dtoClass.ClassCreateRequest

	if err := ctx.Bind(&classCreateDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	class, err := h.SvcClass.Create(ctx, &classCreateDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", class)
	return nil
}

func (h AdminHandlerImpl) CreateClassCategory(ctx echo.Context) error {
	var trainerCreateDto _dtoClassCategory.ClassCategoryCreateRequest

	if err := ctx.Bind(&trainerCreateDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	class, err := h.SvcClassCategory.Create(ctx, &trainerCreateDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", class)
	return nil
}

func (h AdminHandlerImpl) Login(ctx echo.Context) error {
	var adminDto dto.AdminLoginRequest

	if err := ctx.Bind(&adminDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	if err := ctx.Validate(adminDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	res, err := h.SvcAdmin.Login(&adminDto)

	if err != nil {
		response.Err(ctx, http.StatusUnauthorized, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}

func (h AdminHandlerImpl) Refresh(ctx echo.Context) error {
	adminId := ctx.Get("user_id").(float64)

	res, err := h.SvcAdmin.Refresh(uint(adminId))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}
	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}
