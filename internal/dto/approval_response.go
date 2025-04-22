package dto

import "time"

type ApprovalDocumentResponse struct {
	ID         uint                 `json:"id"`
	Note       *string              `json:"note"`
	Status     string               `json:"status"`
	ResolvedAt time.Time            `json:"resolved_at"`
	User       UserDocumentResponse `json:"user"`
}
