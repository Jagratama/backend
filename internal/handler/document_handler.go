package handler

import (
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DocumentHandler struct {
	documentService service.DocumentService
}

func NewDocumentHandler(documentService service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentService: documentService,
	}
}

func (h *DocumentHandler) GetAllDocuments(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	documents, err := h.documentService.GetAllDocuments(ctx, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get documents", err.Error())
	}
	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get documents", documents)
}

func (h *DocumentHandler) GetDocumentBySlug(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	document, err := h.documentService.GetDocumentBySlug(ctx, slug, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get document", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get document", document)
}

func (h *DocumentHandler) CreateDocument(c echo.Context) error {
	ctx := c.Request().Context()
	document := &dto.CreateDocumentRequest{}
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	if err := c.Bind(document); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Invalid request body", err.Error())
	}

	document.UserID = uint(userID)
	newDocument, err := h.documentService.CreateDocument(ctx, document)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to create document", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to create document", newDocument)
}

func (h *DocumentHandler) UpdateDocument(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")
	document := &model.Document{}
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	if err := c.Bind(document); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Invalid request body", err.Error())
	}

	newDocument, err := h.documentService.UpdateDocument(ctx, document, slug, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to update document", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to update document", newDocument)
}

func (h *DocumentHandler) DeleteDocument(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	err := h.documentService.DeleteDocument(ctx, slug, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to delete document", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to delete document", nil)
}

func (h *DocumentHandler) GetDocumentProgress(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	document, err := h.documentService.GetDocumentProgress(ctx, slug, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get document progress", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get document progress", document)
}
