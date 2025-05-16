package repository

import (
	"context"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type ApprovalRequestRepository struct {
	db *gorm.DB
}

func NewApprovalRequestRepository(db *gorm.DB) *ApprovalRequestRepository {
	return &ApprovalRequestRepository{
		db: db,
	}
}

func (r *ApprovalRequestRepository) CreateDocumentApprovalRequest(ctx context.Context, approvalRequest *model.ApprovalRequest) error {
	err := r.db.WithContext(ctx).Create(&approvalRequest).Error
	return err
}

func (r *ApprovalRequestRepository) GetApprovalRequestsByDocumentID(ctx context.Context, documentID int) ([]*model.ApprovalRequest, error) {
	var approvalRequests []*model.ApprovalRequest

	err := r.db.WithContext(ctx).
		Where("document_id = ?", documentID).
		Preload("User").
		Preload("User.File").
		Preload("File").
		Order("id ASC").
		Find(&approvalRequests).Error

	if err != nil {
		return nil, err
	}

	return approvalRequests, nil
}

func (r *ApprovalRequestRepository) UpdateApprovalRequest(ctx context.Context, documentID int, userID int, approvalData *model.ApprovalRequest) error {
	err := r.db.WithContext(ctx).Model(&model.ApprovalRequest{}).
		Where("document_id = ? AND user_id = ?", documentID, userID).
		Updates(approvalData).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *ApprovalRequestRepository) GetApprovalRequest(ctx context.Context, userID int) ([]*model.ApprovalRequest, error) {
	var approvalRequests []*model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("user_id = ? AND status != ?", userID, "pending").Preload("Document").Preload("Document.Category").Preload("Document.File").Preload("Document.User").Preload("Document.User.File").Find(&approvalRequests).Error
	if err != nil {
		return nil, err
	}
	return approvalRequests, nil
}

func (r *ApprovalRequestRepository) GetPendingApprovalRequest(ctx context.Context, userID int) ([]*model.ApprovalRequest, error) {
	var approvalRequests []*model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, "pending").Preload("Document").Preload("Document.Category").Preload("Document.File").Preload("Document.User").Preload("Document.User.File").Find(&approvalRequests).Error
	if err != nil {
		return nil, err
	}
	return approvalRequests, nil
}

func (r *ApprovalRequestRepository) GetUnApprovedApprovalByDocumentID(ctx context.Context, documentID int) ([]*model.ApprovalRequest, error) {
	var approvalRequests []*model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("document_id = ? AND status != ?", documentID, "approved").Find(&approvalRequests).Error
	if err != nil {
		return nil, err
	}

	return approvalRequests, nil
}

func (r *ApprovalRequestRepository) GetLatestApprovalRequestApproved(ctx context.Context, documentID int) (*model.ApprovalRequest, error) {
	var approvalRequest model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("document_id = ? AND status = ?", documentID, "approved").Order("id DESC").Preload("File").First(&approvalRequest).Error
	if err != nil {
		return nil, err
	}

	return &approvalRequest, nil
}
