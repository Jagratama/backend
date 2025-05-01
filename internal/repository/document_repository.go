package repository

import (
	"context"
	"errors"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

func (r *DocumentRepository) GetAllDocuments(ctx context.Context, userID int) ([]*model.Document, error) {
	var documents []*model.Document

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User").Preload("User.File").Preload("Category").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) GetDocumentByID(ctx context.Context, id int) (*model.Document, error) {
	var document *model.Document

	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("User").Preload("User.File").Preload("Category").First(&document).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle not found case specifically
			return nil, gorm.ErrRecordNotFound
		}
		// Handle other errors
		return nil, err
	}
	return document, err
}

func (r *DocumentRepository) GetDocumentBySlug(ctx context.Context, slug string, userID int) (*model.Document, error) {
	var document model.Document

	err := r.db.WithContext(ctx).Where("slug = ?", slug).Where("user_id", userID).Preload("User").Preload("User.File").Preload("Category").First(&document).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle not found case specifically
			return nil, gorm.ErrRecordNotFound
		}
		// Handle other errors
		return nil, err
	}
	return &document, nil
}

func (r *DocumentRepository) CreateDocument(ctx context.Context, document *model.Document) (*model.Document, error) {
	err := r.db.WithContext(ctx).Create(&document).Error
	return document, err
}

func (r *DocumentRepository) UpdateDocumentBySlug(ctx context.Context, documentData *model.Document, slug string, userID int) (*model.Document, error) {
	var document model.Document
	err := r.db.WithContext(ctx).Where("slug = ?", slug).Where("user_id", userID).First(&document).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&document).Updates(documentData).Error
	if err != nil {
		return nil, err
	}
	return &document, err
}

func (r *DocumentRepository) DeleteDocument(ctx context.Context, slug string, userID int) error {
	var document model.Document
	err := r.db.WithContext(ctx).Where("slug = ?", slug).Where("user_id", userID).Delete(&document).Error
	return err
}

func (r *DocumentRepository) GetAllDocumentsNeedApprove(ctx context.Context, userID int) ([]*model.Document, error) {
	var documents []*model.Document

	err := r.db.WithContext(ctx).Preload("ApprovalRequest").Find(&documents).Error
	return documents, err
}
