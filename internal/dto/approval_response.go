package dto

import "time"

type ApprovalDocumentResponse struct {
	ID         uint                 `json:"id"`
	Note       *string              `json:"note"`
	Status     string               `json:"status"`
	ResolvedAt time.Time            `json:"resolved_at"`
	FilePath   string               `json:"file_path"`
	User       UserDocumentResponse `json:"user"`
}
