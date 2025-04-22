package repository

import (
	"context"
	"jagratama-backend/internal/dto"
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

func (r *ApprovalRequestRepository) ApprovalAction(ctx context.Context, documentID int, userID int, approvalRequestData *dto.ApprovalDocumentRequest) error {
	var approvalRequest *model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("document_id = ?", documentID).Where("user_id = ?", userID).Where("status = ?", "pending").First(&approvalRequest).Error
	if err != nil {
		return err
	}

	approvalRequest.Status = approvalRequestData.Status
	approvalRequest.Note = approvalRequestData.Note
	err = r.db.WithContext(ctx).Save(&approvalRequest).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ApprovalRequestRepository) GetApprovalRequest(ctx context.Context, userID int) ([]*model.ApprovalRequest, error) {
	var approvalRequests []*model.ApprovalRequest

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("Document").Preload("Document.User").Preload("Document.Category").Find(&approvalRequests).Error
	if err != nil {
		return nil, err
	}
	return approvalRequests, nil
}
