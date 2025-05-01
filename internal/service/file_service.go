package service

import (
	"bytes"
	"context"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/pkg/aws"
	"jagratama-backend/internal/repository"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type FileService struct {
	fileRepository repository.FileRepository
	s3Uploader     *aws.S3Uploader
}

func NewFileService(fileRepository repository.FileRepository, s3Uploader *aws.S3Uploader) *FileService {
	return &FileService{
		fileRepository: fileRepository,
		s3Uploader:     s3Uploader,
	}
}

func (s *FileService) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, uploadType string) (*model.File, error) {
	var resp *model.File

	folder := "profile"
	if uploadType == "document" {
		folder = "document"
	}

	// Compress image file
	defer file.Close()

	// Detect content type
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return resp, err
	}
	contentType := http.DetectContentType(buff)

	// Reset the file reader
	_, err = file.Seek(0, 0)
	if err != nil {
		return resp, err
	}

	var fileToUpload *bytes.Buffer

	extension := filepath.Ext(fileHeader.Filename)
	fileNameFlat, err := helpers.GenerateSlug(strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename)), 8)
	if err != nil {
		return resp, err
	}

	if strings.HasPrefix(contentType, "image/") {
		// Compress image
		fileToUpload, contentType, err = helpers.CompressImage(file, contentType)
		if err != nil {
			return resp, err
		}
	} else {
		// Non-image files use original
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(file)
		if err != nil {
			return resp, err
		}
		fileToUpload = buf
	}

	newFileName := fileNameFlat + extension

	// Upload to S3
	_, err = s.s3Uploader.UploadBuffer(ctx, fileToUpload.Bytes(), newFileName, folder, contentType)
	if err != nil {
		return resp, err
	}

	resp = &model.File{
		FileName:    fileHeader.Filename,
		FilePath:    newFileName,
		ContentType: contentType,
	}

	newFile, err := s.fileRepository.Create(ctx, resp)
	if err != nil {
		return resp, err
	}

	return newFile, nil
}
