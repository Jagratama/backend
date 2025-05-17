package service

import (
	"context"
	"errors"
	"fmt"
	"jagratama-backend/internal/config"
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
	"time"

	"gorm.io/gorm"
)

type DocumentService struct {
	documentRepository        repository.DocumentRepository
	approvalRequestRepository repository.ApprovalRequestRepository
	userRepository            repository.UserRepository
}

func NewDocumentService(documentRepository repository.DocumentRepository, approvalRequestRepository repository.ApprovalRequestRepository, userRepository repository.UserRepository) *DocumentService {
	return &DocumentService{
		documentRepository:        documentRepository,
		approvalRequestRepository: approvalRequestRepository,
		userRepository:            userRepository,
	}
}

func (s *DocumentService) GetAllDocuments(ctx context.Context, userID int) ([]*dto.DocumentResponse, error) {
	documents, err := s.documentRepository.GetAllDocuments(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.DocumentResponse, 0)
	for _, document := range documents {
		response = append(response, &dto.DocumentResponse{
			ID:         document.ID,
			Title:      document.Title,
			Slug:       document.Slug,
			File:       config.GetEnv("AWS_S3_URL", "") + document.File.FilePath,
			LastStatus: document.LastStatus,
			ApprovedAt: document.ApprovedAt,
			User: dto.UserDocumentResponse{
				ID:    document.User.ID,
				Name:  document.User.Name,
				Email: document.User.Email,
				Image: config.GetEnv("AWS_S3_URL", "") + document.User.File.FilePath,
			},
			AddressedUser: dto.UserDocumentResponse{
				ID:    document.AddressedUser.ID,
				Name:  document.AddressedUser.Name,
				Email: document.AddressedUser.Email,
				Image: config.GetEnv("AWS_S3_URL", "") + document.AddressedUser.File.FilePath,
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
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	// authorization validation

	response := &dto.DocumentResponse{
		ID:          document.ID,
		Title:       document.Title,
		Description: document.Description,
		Slug:        document.Slug,
		File:        config.GetEnv("AWS_S3_URL", "") + document.File.FilePath,
		LastStatus:  document.LastStatus,
		ApprovedAt:  document.ApprovedAt,
		User: dto.UserDocumentResponse{
			ID:    document.User.ID,
			Name:  document.User.Name,
			Email: document.User.Email,
			Image: config.GetEnv("AWS_S3_URL", "") + document.User.File.FilePath,
		},
		AddressedUser: dto.UserDocumentResponse{
			ID:    document.AddressedUser.ID,
			Name:  document.AddressedUser.Name,
			Email: document.AddressedUser.Email,
			Image: config.GetEnv("AWS_S3_URL", "") + document.AddressedUser.File.FilePath,
		},
		Category: dto.CategoryResponse{
			ID:   document.Category.ID,
			Name: document.Category.Name,
		},
	}
	return response, err
}

func (s *DocumentService) CreateDocument(ctx context.Context, documentRequest *dto.CreateDocumentRequest) (*dto.DocumentResponse, error) {
	if len(documentRequest.ApproverEmails) < 2 {
		return nil, fmt.Errorf("approver emails must be at least 2")
	}

	// Validate approvers can't repeat
	approverMap := make(map[string]bool)
	approverIDs := []int{}
	for _, approverEmail := range documentRequest.ApproverEmails {
		if approverMap[approverEmail] {
			return nil, fmt.Errorf("approver with ID %s already exists", approverEmail)
		}
		approverMap[approverEmail] = true

		user, err := s.userRepository.GetUserByEmail(ctx, approverEmail)
		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, fmt.Errorf("user with email %s not found", approverEmail)
		}

		approverIDs = append(approverIDs, int(user.ID))
	}

	slug, err := helpers.GenerateSlug(documentRequest.Title, 8)
	if err != nil {
		return nil, err
	}
	documentRequest.Slug = slug

	lastApproverID := approverIDs[len(approverIDs)-1]
	// Create the document
	document := &model.Document{
		UserID:          documentRequest.UserID,
		AddressedUserID: uint(lastApproverID),
		FileID:          documentRequest.FileID,
		CategoryID:      documentRequest.CategoryID,
		Title:           documentRequest.Title,
		Slug:            documentRequest.Slug,
		Description:     documentRequest.Description,
	}
	newDocument, err := s.documentRepository.CreateDocument(ctx, document)
	if err != nil {
		return nil, err
	}

	// Create approvers
	for _, approverID := range approverIDs {
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
		ID:    document.ID,
		Title: document.Title,
		Slug:  document.Slug,
		File:  config.GetEnv("AWS_S3_URL", "") + document.File.FilePath,
		User: dto.UserDocumentResponse{
			ID:    document.User.ID,
			Name:  document.User.Name,
			Email: document.User.Email,
			Image: config.GetEnv("AWS_S3_URL", "") + document.User.File.FilePath,
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
	_, err := s.documentRepository.GetDocumentBySlug(ctx, slug)
	if err != nil {
		return err
	}

	err = s.documentRepository.DeleteDocument(ctx, slug, userID)
	return err
}

func (s *DocumentService) GetDocumentProgress(ctx context.Context, slug string, userID int) ([]*dto.ApprovalDocumentResponse, error) {
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	// authorization validation

	approvalRequests, err := s.approvalRequestRepository.GetApprovalRequestsByDocumentID(ctx, int(document.ID))
	if err != nil {
		return nil, err
	}

	if len(approvalRequests) == 0 {
		return nil, nil
	}

	var response []*dto.ApprovalDocumentResponse
	for _, approvalRequest := range approvalRequests {
		approvalFilePath := config.GetEnv("AWS_S3_URL", "") + approvalRequest.File.FilePath
		if approvalRequest.FileID == nil {
			approvalFilePath = ""
		}

		response = append(response, &dto.ApprovalDocumentResponse{
			ID:         approvalRequest.ID,
			Note:       approvalRequest.Note,
			Status:     approvalRequest.Status,
			File:       approvalFilePath,
			ResolvedAt: approvalRequest.ResolvedAt,
			User: dto.UserDocumentResponse{
				ID:    approvalRequest.User.ID,
				Name:  approvalRequest.User.Name,
				Email: approvalRequest.User.Email,
				Image: config.GetEnv("AWS_S3_URL", "") + document.User.File.FilePath,
			},
		})
	}
	return response, nil
}

func (s *DocumentService) ApprovalAction(ctx context.Context, slug string, userID int, approvalRequest *dto.ApprovalDocumentRequest) error {
	if approvalRequest.Status != dto.StatusApprove && approvalRequest.Status != dto.StatusReject {
		return fmt.Errorf("invalid status: %s", approvalRequest.Status)
	}

	if approvalRequest.Status == dto.StatusReject && approvalRequest.Note == nil {
		return fmt.Errorf("note is required when rejecting")
	}

	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug)
	if err != nil {
		return err
	}

	userApprovals, err := s.approvalRequestRepository.GetApprovalRequestsByDocumentID(ctx, int(document.ID))
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	requiresSignature := false
	if len(user.Position.RequiresSignatureByCategoryType) > 0 {
		for _, categoryType := range user.Position.RequiresSignatureByCategoryType {
			if categoryType == document.Category.Type {
				requiresSignature = true
				break
			}
		}
	}

	approvalData := &model.ApprovalRequest{}
	found := false

	for i, approval := range userApprovals {
		if int(approval.User.ID) == userID {
			found = true

			if approval.Status != "pending" {
				return fmt.Errorf("you have already approved or rejected this document")
			}

			if i != 0 && userApprovals[i-1].Status == "pending" {
				return fmt.Errorf("the previous approver has not yet approved or rejected this document")
			}

			approvalData = approval
			approvalData.Status = approvalRequest.Status
			approvalData.Note = approvalRequest.Note
			approvalData.ResolvedAt = time.Now()
			if approvalRequest.FileID != 0 {
				if requiresSignature {
					fileID := uint(approvalRequest.FileID)
					approvalData.FileID = &fileID
				} else {
					// user does not require signature, use file id from before approver
					if i > 0 {
						approvalData.FileID = userApprovals[i-1].FileID
					} else {
						approvalData.FileID = &document.FileID
					}
				}
			} else {
				if i > 0 {
					approvalData.FileID = userApprovals[i-1].FileID
				} else {
					approvalData.FileID = &document.FileID
				}
			}

			break
		}
	}

	if !found {
		return fmt.Errorf("you are not authorized to approve this document")
	}

	err = s.approvalRequestRepository.UpdateApprovalRequest(ctx, int(document.ID), userID, approvalData)
	if err != nil {
		return err
	}

	unApprovedApprovals, err := s.approvalRequestRepository.GetUnApprovedApprovalByDocumentID(ctx, int(document.ID))
	if err != nil {
		return err
	}

	documentLastStatus := dto.StatusPending

	if approvalData.Status == dto.StatusReject {
		documentLastStatus = dto.StatusReject
	} else if len(unApprovedApprovals) == 0 && approvalData.Status == dto.StatusApprove {
		documentLastStatus = dto.StatusApprove
	}

	err = s.documentRepository.UpdateDocumentAlreadyApproved(ctx, int(document.ID), documentLastStatus)
	if err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) GetDocumentApprovalRequest(ctx context.Context, userID int) ([]*dto.DocumentRequestResponse, error) {
	myApprovalRequests, err := s.approvalRequestRepository.GetPendingApprovalRequest(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.DocumentRequestResponse, 0)
	for _, myApprovalRequest := range myApprovalRequests {
		allApprovals, err := s.approvalRequestRepository.GetApprovalRequestsByDocumentID(ctx, int(myApprovalRequest.Document.ID))
		if err != nil {
			return nil, err
		}

		var canReview = false
		for idx, approval := range allApprovals {
			if approval.User.ID == uint(userID) {
				if idx == 0 {
					canReview = true
				} else {
					if allApprovals[idx-1].Status == "approved" {
						canReview = true
						break
					}
				}
			}

		}

		if canReview {
			response = append(response, &dto.DocumentRequestResponse{
				ID:     myApprovalRequest.Document.ID,
				Title:  myApprovalRequest.Document.Title,
				Slug:   myApprovalRequest.Document.Slug,
				File:   config.GetEnv("AWS_S3_URL", "") + myApprovalRequest.Document.File.FilePath,
				Status: myApprovalRequest.Status,
				User: dto.UserDocumentResponse{
					ID:    myApprovalRequest.Document.User.ID,
					Name:  myApprovalRequest.Document.User.Name,
					Email: myApprovalRequest.Document.User.Email,
					Image: config.GetEnv("AWS_S3_URL", "") + myApprovalRequest.Document.User.File.FilePath,
				},
				Category: dto.CategoryResponse{
					ID:   myApprovalRequest.Document.Category.ID,
					Name: myApprovalRequest.Document.Category.Name,
				},
			})
		}
	}

	return response, nil
}

func (s *DocumentService) GetDocumentApprovalHistory(ctx context.Context, userID int) ([]*dto.DocumentRequestResponse, error) {
	approvalRequests, err := s.approvalRequestRepository.GetApprovalRequest(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.DocumentRequestResponse, 0)
	for _, approvalRequest := range approvalRequests {
		response = append(response, &dto.DocumentRequestResponse{
			ID:     approvalRequest.Document.ID,
			Title:  approvalRequest.Document.Title,
			Slug:   approvalRequest.Document.Slug,
			File:   config.GetEnv("AWS_S3_URL", "") + approvalRequest.Document.File.FilePath,
			Status: approvalRequest.Status,
			User: dto.UserDocumentResponse{
				ID:    approvalRequest.Document.User.ID,
				Name:  approvalRequest.Document.User.Name,
				Email: approvalRequest.Document.User.Email,
				Image: config.GetEnv("AWS_S3_URL", "") + approvalRequest.Document.User.File.FilePath,
			},
			Category: dto.CategoryResponse{
				ID:   approvalRequest.Document.Category.ID,
				Name: approvalRequest.Document.Category.Name,
			},
		})
	}

	return response, nil
}

func (s *DocumentService) GetCountAllMyDocuments(ctx context.Context, userID int) (dto.DocumentCountResponse, error) {
	response := dto.DocumentCountResponse{}

	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return response, err
	}

	countAllDocuments := int64(0)
	countPendingDocuments := int64(0)
	countRejectedDocuments := int64(0)
	countApprovedDocuments := int64(0)
	countUsers := int64(0)

	if user.Role.Name == "requester" {
		countAllDocuments, err = s.documentRepository.CountAllMyDocuments(ctx, userID)
		if err != nil {
			return response, err
		}

		countPendingDocuments, err = s.documentRepository.CountPendingDocuments(ctx, userID)
		if err != nil {
			return response, err
		}

		countRejectedDocuments, err = s.documentRepository.CountRejectedDocuments(ctx, userID)
		if err != nil {
			return response, err
		}

		countApprovedDocuments, err = s.documentRepository.CountApprovedDocuments(ctx, userID)
		if err != nil {
			return response, err
		}
	}

	if (user.Role.Name == "reviewer") || (user.Role.Name == "approver") {
		myApprovalRequests, err := s.approvalRequestRepository.GetPendingApprovalRequest(ctx, userID)
		if err != nil {
			return response, err
		}

		for _, myApprovalRequest := range myApprovalRequests {
			allApprovals, err := s.approvalRequestRepository.GetApprovalRequestsByDocumentID(ctx, int(myApprovalRequest.Document.ID))
			if err != nil {
				return response, err
			}

			for idx, approval := range allApprovals {
				if approval.User.ID == uint(userID) {
					if idx == 0 {
						countPendingDocuments++
					} else {
						if allApprovals[idx-1].Status == "approved" {
							countPendingDocuments++
							break
						}
					}
				}

			}
		}

		countRejectedDocuments, err = s.approvalRequestRepository.CountApprovalDocumentsByStatus(ctx, userID, dto.StatusReject)
		if err != nil {
			return response, err
		}

		countApprovedDocuments, err = s.approvalRequestRepository.CountApprovalDocumentsByStatus(ctx, userID, dto.StatusApprove)
		if err != nil {
			return response, err
		}

		countAllDocuments = countPendingDocuments + countRejectedDocuments + countApprovedDocuments
	}

	if user.Role.Name == "admin" {
		countAllDocuments, err = s.documentRepository.CountAllDocuments(ctx)
		if err != nil {
			return response, err
		}

		countRejectedDocuments, err = s.documentRepository.CountAllDocumentsByStatus(ctx, dto.StatusReject)
		if err != nil {
			return response, err
		}
		countApprovedDocuments, err = s.documentRepository.CountAllDocumentsByStatus(ctx, dto.StatusApprove)
		if err != nil {
			return response, err
		}

		countUsers, err = s.userRepository.CountAllUsers(ctx)
		if err != nil {
			return response, err
		}
	}

	response = dto.DocumentCountResponse{
		TotalDocument: int(countAllDocuments),
		TotalRejected: int(countRejectedDocuments),
		TotalPending:  int(countPendingDocuments),
		TotalApproved: int(countApprovedDocuments),
		TotalUsers:    int(countUsers),
	}
	return response, nil
}

func (s *DocumentService) GetDocumentApprovalReviewDetail(ctx context.Context, slug string, userID int) (*dto.ApprovalDocumentDetailResponse, error) {
	document, err := s.documentRepository.GetDocumentBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	approvalRequests, err := s.approvalRequestRepository.GetLatestApprovalRequestApproved(ctx, int(document.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var filePath string
	if approvalRequests != nil && approvalRequests.FileID != nil {
		// approved by someone, get file from before approver
		filePath = approvalRequests.File.FilePath
	} else {
		// no one approved, get file from document
		filePath = document.File.FilePath
	}

	requiresSignature := false
	if len(user.Position.RequiresSignatureByCategoryType) > 0 {
		for _, categoryType := range user.Position.RequiresSignatureByCategoryType {
			if categoryType == document.Category.Type {
				requiresSignature = true
				break
			}
		}
	}

	IsReviewer := false
	if user.Role.Name == "reviewer" {
		IsReviewer = true
	}

	return &dto.ApprovalDocumentDetailResponse{
		Title:             document.Title,
		File:              config.GetEnv("AWS_S3_URL", "") + filePath,
		RequiresSignature: requiresSignature,
		IsReviewer:        IsReviewer,
	}, nil
}
