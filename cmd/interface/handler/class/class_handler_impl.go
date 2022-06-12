package class

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class/service"
	"gym/internal/protocol/http/response"
	"gym/pkg/database"
	"net/http"
	"strconv"
)

type ClassHandlerImpl struct {
	SvcClass service.ClassService
}

func (h ClassHandlerImpl) Get(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)

	classes, err := h.SvcClass.GetAll(ctx, pagination)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"classes": map[string]interface{}{
			"data":       classes,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h ClassHandlerImpl) Detail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	class, err := h.SvcClass.GetById(ctx, uint(id))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", class)
	return nil
}
