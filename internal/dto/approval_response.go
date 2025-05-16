package dto

import "time"

type ApprovalDocumentResponse struct {
	ID         uint                 `json:"id"`
	Note       *string              `json:"note"`
	Status     string               `json:"status"`
	ResolvedAt time.Time            `json:"resolved_at"`
	File       string               `json:"file"`
	User       UserDocumentResponse `json:"user"`
}

type ApprovalDocumentDetailResponse struct {
	Title             string `json:"title"`
	File              string `json:"file"`
	RequiresSignature bool   `json:"requires_signature"`
	IsReviewer        bool   `json:"is_reviewer"`
}
