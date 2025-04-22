package service

import (
	"context"
	"fmt"
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
	// Validate approvers can't repeat
	approverMap := make(map[int]bool)
	for _, approverID := range documentRequest.Approvers {
		if approverMap[approverID] {
			return nil, fmt.Errorf("approver with ID %d already exists", approverID)
		}
		approverMap[approverID] = true
	}

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
			UserID:     uint(approverID),
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

func (s *DocumentService) ApprovalAction(ctx context.Context, slug string, userID int, approvalRequest *dto.ApprovalDocumentRequest) error {
	if approvalRequest.Status != dto.StatusApprove && approvalRequest.Status != dto.StatusReject {
		return fmt.Errorf("invalid status: %s", approvalRequest.Status)
	}

	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug, userID)
	if err != nil {
		return err
	}

	err = s.approvalRequestRepository.ApprovalAction(ctx, int(document.ID), userID, approvalRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) GetDocumentApprovalRequest(ctx context.Context, userID int) ([]*dto.DocumentRequestResponse, error) {
	approvalRequests, err := s.approvalRequestRepository.GetApprovalRequest(ctx, userID)
	if err != nil {
		return nil, err
	}

	var response []*dto.DocumentRequestResponse
	for _, approvalRequest := range approvalRequests {
		response = append(response, &dto.DocumentRequestResponse{
			ID:       approvalRequest.Document.ID,
			Title:    approvalRequest.Document.Title,
			Slug:     approvalRequest.Document.Slug,
			FilePath: approvalRequest.Document.FilePath,
			Status:   approvalRequest.Status,
			User: dto.UserDocumentResponse{
				ID:        approvalRequest.Document.User.ID,
				Name:      approvalRequest.Document.User.Name,
				Email:     approvalRequest.Document.User.Email,
				ImagePath: approvalRequest.Document.User.ImagePath,
			},
			Category: dto.CategoryResponse{
				ID:   approvalRequest.Document.Category.ID,
				Name: approvalRequest.Document.Category.Name,
			},
		})
	}

	return response, nil
}
