package repository

import (
	"context"
	"errors"
	"jagratama-backend/internal/dto"
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

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Where("confirmed = ?", true).Preload("File").Preload("User").Preload("User.File").Preload("Category").Preload("AddressedUser").Preload("AddressedUser.File").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) GetDocumentByID(ctx context.Context, id int) (*model.Document, error) {
	var document *model.Document

	err := r.db.WithContext(ctx).Where("id = ?", id).Where("confirmed = ?", true).Preload("File").Preload("User").Preload("User.File").Preload("Category").Preload("AddressedUser").Preload("AddressedUser.File").First(&document).Error
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

func (r *DocumentRepository) GetDocumentBySlug(ctx context.Context, slug string) (*model.Document, error) {
	var document model.Document

	err := r.db.WithContext(ctx).Where("slug = ?", slug).Preload("File").Preload("User").Preload("User.File").Preload("Category").Preload("AddressedUser").Preload("AddressedUser.File").First(&document).Error
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

func (r *DocumentRepository) CountAllDocuments(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Count(&count).Error
	return count, err
}

func (r *DocumentRepository) CountAllDocumentsByStatus(ctx context.Context, status string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Where("last_status = ?", status).Count(&count).Error
	return count, err
}

func (r *DocumentRepository) CountAllMyDocuments(ctx context.Context, userID int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *DocumentRepository) CountPendingDocuments(ctx context.Context, userID int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Where("user_id = ?", userID).Where("last_status = ?", "pending").Count(&count).Error
	return count, err
}

func (r *DocumentRepository) CountRejectedDocuments(ctx context.Context, userID int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Where("user_id = ?", userID).Where("last_status = ?", "rejected").Count(&count).Error
	return count, err
}

func (r *DocumentRepository) CountApprovedDocuments(ctx context.Context, userID int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Document{}).Where("user_id = ?", userID).Where("last_status = ?", "approved").Count(&count).Error
	return count, err
}

func (r *DocumentRepository) UpdateDocumentAlreadyApproved(ctx context.Context, documentID int, status string) error {
	var document model.Document

	var approvedAt interface{} = nil
	if status == dto.StatusApprove {
		approvedAt = gorm.Expr("NOW()")
	}

	err := r.db.WithContext(ctx).Model(&document).Where("id = ?", documentID).UpdateColumns(map[string]interface{}{
		"last_status": status,
		"approved_at": approvedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
