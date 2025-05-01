package handler

import (
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get categories", err.Error())
	}
	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully get all categories", categories)
}
