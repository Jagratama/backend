package dto

type CreateDocumentRequest struct {
	UserID         uint     `json:"user_id"`
	CategoryID     uint     `json:"category_id"`
	Title          string   `json:"title"`
	Slug           string   `json:"slug"`
	Description    string   `json:"description"`
	FilePath       string   `json:"file_path"`
	ApproverEmails []string `json:"approvers"`
}
