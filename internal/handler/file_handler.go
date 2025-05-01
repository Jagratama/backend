package handler

import (
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	fileService service.FileService
}

func NewFileHandler(uploadService service.FileService) *FileHandler {
	return &FileHandler{
		fileService: uploadService,
	}
}

func (h *FileHandler) UploadFile(c echo.Context) error {
	ctx := c.Request().Context()
	file, err := c.FormFile("file")
	uploadType := c.Param("type")
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Failed to get file", err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer src.Close()

	newFile, err := h.fileService.UploadFile(ctx, src, file, uploadType)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to upload file", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "File uploaded successfully", newFile)
}
