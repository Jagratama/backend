package handler

import (
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PositionHandler struct {
	positionService service.PositionService
}

func NewPositionHandler(positionService service.PositionService) *PositionHandler {
	return &PositionHandler{
		positionService: positionService,
	}
}

func (h *PositionHandler) GetAllPositions(c echo.Context) error {
	ctx := c.Request().Context()
	positions, err := h.positionService.GetAllPositions(ctx)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get all positions", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get all positions", positions)
}

func (h *PositionHandler) GetPositionByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid position ID", err.Error())
	}

	position, err := h.positionService.GetPositionByID(ctx, id)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get position", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get position", position)
}
