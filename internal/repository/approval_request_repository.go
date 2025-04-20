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

	err := r.db.WithContext(ctx).Where("document_id = ?", documentID).Preload("User").Find(&approvalRequests).Error
	if err != nil {
		return nil, err
	}
	return approvalRequests, nil
}
