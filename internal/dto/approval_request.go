package dto

const (
	StatusApprove = "approved"
	StatusReject  = "rejected"
)

type ApprovalDocumentRequest struct {
	Status string  `json:"status" binding:"required,oneof=approved rejected"`
	Note   *string `json:"note" binding:"omitempty,max=500"`
}
