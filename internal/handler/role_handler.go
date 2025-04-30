package handler

import (
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) GetRoleByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid role ID", err.Error())
	}

	role, err := h.roleService.GetRoleByID(ctx, id)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get role", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get role", role)
}

func (h *RoleHandler) GetAllRoles(c echo.Context) error {
	ctx := c.Request().Context()
	roles, err := h.roleService.GetAllRoles(ctx)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get all roles", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get all roles", roles)
}
