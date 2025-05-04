package dto

type CreateDocumentRequest struct {
	FileID         uint     `json:"file_id"`
	UserID         uint     `json:"user_id"`
	CategoryID     uint     `json:"category_id"`
	Title          string   `json:"title"`
	Slug           string   `json:"slug"`
	Description    string   `json:"description"`
	ApproverEmails []string `json:"approvers"`
}
