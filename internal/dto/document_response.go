package dto

import "time"

type DocumentResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User     UserDocumentResponse `json:"user"`
	Category CategoryResponse     `json:"category"`
}

type DocumentProgressResponse struct {
	Note       *string              `json:"note"`
	Status     string               `json:"status"`
	ResolvedAt *time.Time           `json:"resolved_at"`
	User       UserDocumentResponse `json:"user"`
}
