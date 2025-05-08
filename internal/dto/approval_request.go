package dto

const (
	StatusApprove = "approved"
	StatusReject  = "rejected"
	StatusPending = "pending"
)

type ApprovalDocumentRequest struct {
	Status string  `json:"status" binding:"required,oneof=approved rejected"`
	Note   *string `json:"note" binding:"omitempty,max=500"`
	FileID int     `json:"file_id" binding:"required"`
}
