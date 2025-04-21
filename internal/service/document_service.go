package service

import (
	"context"
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type DocumentService struct {
	documentRepository        repository.DocumentRepository
	approvalRequestRepository repository.ApprovalRequestRepository
}

func NewDocumentService(documentRepository repository.DocumentRepository, approvalRequestRepository repository.ApprovalRequestRepository) *DocumentService {
	return &DocumentService{
		documentRepository:        documentRepository,
		approvalRequestRepository: approvalRequestRepository,
	}
}

func (s *DocumentService) GetAllDocuments(ctx context.Context, userID int) ([]*dto.DocumentResponse, error) {
	documents, err := s.documentRepository.GetAllDocuments(ctx, userID)
	if err != nil {
		return nil, err
	}

	var response []*dto.DocumentResponse
	for _, document := range documents {
		response = append(response, &dto.DocumentResponse{
			ID:       document.ID,
			Title:    document.Title,
			Slug:     document.Slug,
			FilePath: document.FilePath,
			User: dto.UserDocumentResponse{
				ID:        document.User.ID,
				Name:      document.User.Name,
				Email:     document.User.Email,
				ImagePath: document.User.ImagePath,
			},
			Category: dto.CategoryResponse{
				ID:   document.Category.ID,
				Name: document.Category.Name,
			},
		})
	}
	return response, nil
}

func (s *DocumentService) GetDocumentBySlug(ctx context.Context, slug string, userID int) (*dto.DocumentResponse, error) {
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug, userID)
	if err != nil {
		return nil, err
	}

	response := &dto.DocumentResponse{
		ID:       document.ID,
		Title:    document.Title,
		Slug:     document.Slug,
		FilePath: document.FilePath,
		User: dto.UserDocumentResponse{
			ID:        document.User.ID,
			Name:      document.User.Name,
			Email:     document.User.Email,
			ImagePath: document.User.ImagePath,
		},
		Category: dto.CategoryResponse{
			ID:   document.Category.ID,
			Name: document.Category.Name,
		},
	}
	return response, err
}

func (s *DocumentService) CreateDocument(ctx context.Context, documentRequest *dto.CreateDocumentRequest) (*dto.DocumentResponse, error) {
	slug, err := helpers.GenerateSlug(documentRequest.Title)
	if err != nil {
		return nil, err
	}
	documentRequest.Slug = slug

	// Create the document
	document := &model.Document{
		UserID:      documentRequest.UserID,
		CategoryID:  documentRequest.CategoryID,
		Title:       documentRequest.Title,
		Slug:        documentRequest.Slug,
		Description: documentRequest.Description,
		FilePath:    documentRequest.FilePath,
	}
	newDocument, err := s.documentRepository.CreateDocument(ctx, document)
	if err != nil {
		return nil, err
	}

	// Create approvers
	for _, approverID := range documentRequest.Approvers {
		approver := &model.ApprovalRequest{
			DocumentID: newDocument.ID,
			UserID:     approverID,
		}
		err = s.approvalRequestRepository.CreateDocumentApprovalRequest(ctx, approver)
		if err != nil {
			return nil, err
		}
	}

	response := &dto.DocumentResponse{
		ID:       document.ID,
		Title:    document.Title,
		Slug:     document.Slug,
		FilePath: document.FilePath,
		User: dto.UserDocumentResponse{
			ID:        document.User.ID,
			Name:      document.User.Name,
			Email:     document.User.Email,
			ImagePath: document.User.ImagePath,
		},
		Category: dto.CategoryResponse{
			ID:   document.Category.ID,
			Name: document.Category.Name,
		},
	}
	return response, nil
}

func (s *DocumentService) UpdateDocument(ctx context.Context, document *model.Document, slug string, userID int) (*model.Document, error) {
	updatedDocument, err := s.documentRepository.UpdateDocumentBySlug(ctx, document, slug, userID)
	return updatedDocument, err
}

func (s *DocumentService) DeleteDocument(ctx context.Context, slug string, userID int) error {
	_, err := s.documentRepository.GetDocumentBySlug(ctx, slug, userID)
	if err != nil {
		return err
	}

	err = s.documentRepository.DeleteDocument(ctx, slug, userID)
	return err
}

func (s *DocumentService) GetDocumentProgress(ctx context.Context, slug string, userID int) ([]*dto.ApprovalDocumentResponse, error) {
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug, userID)
	if err != nil {
		return nil, err
	}

	approvalRequests, err := s.approvalRequestRepository.GetApprovalRequestsByDocumentID(ctx, int(document.ID))
	if err != nil {
		return nil, err
	}

	if len(approvalRequests) == 0 {
		return nil, nil
	}

	var response []*dto.ApprovalDocumentResponse
	for _, approvalRequest := range approvalRequests {
		response = append(response, &dto.ApprovalDocumentResponse{
			ID:         approvalRequest.ID,
			Note:       approvalRequest.Note,
			Status:     approvalRequest.Status,
			ResolvedAt: approvalRequest.ResolvedAt,
			User: dto.UserDocumentResponse{
				ID:        approvalRequest.User.ID,
				Name:      approvalRequest.User.Name,
				Email:     approvalRequest.User.Email,
				ImagePath: approvalRequest.User.ImagePath,
			},
		})
	}
	return response, nil
}
