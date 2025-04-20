package service

import (
	"context"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type DocumentService struct {
	documentRepository repository.DocumentRepository
}

func NewDocumentService(documentRepository repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepository: documentRepository,
	}
}

func (s *DocumentService) GetAllDocuments(ctx context.Context, userID int) ([]*model.Document, error) {
	documents, err := s.documentRepository.GetAllDocuments(ctx, userID)
	return documents, err
}

func (s *DocumentService) GetDocumentBySlug(ctx context.Context, slug string, userID int) (*model.Document, error) {
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug, userID)
	return document, err
}

func (s *DocumentService) CreateDocument(ctx context.Context, document *model.Document) (*model.Document, error) {
	slug, err := helpers.GenerateSlug(document.Title)
	if err != nil {
		return nil, err
	}
	document.Slug = slug
	newDocument, err := s.documentRepository.CreateDocument(ctx, document)
	return newDocument, err
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
