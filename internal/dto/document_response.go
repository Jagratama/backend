package dto

import "time"

type DocumentResponse struct {
	ID              uint       `json:"id"`
	UserID          uint       `json:"user_id"`
	AddressedUserID uint       `json:"addressed_user_id"`
	CategoryID      uint       `json:"category_id"`
	Title           string     `json:"title"`
	Slug            string     `json:"slug"`
	Description     string     `json:"description"`
	File            string     `json:"file"`
	LastStatus      string     `json:"last_status"`
	ApprovedAt      *time.Time `json:"approved_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	User          UserDocumentResponse `json:"user"`
	AddressedUser UserDocumentResponse `json:"addressed_user"`
	Category      CategoryResponse     `json:"category"`
}

type DocumentRequestResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	File        string    `json:"file"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User     UserDocumentResponse `json:"user"`
	Category CategoryResponse     `json:"category"`
}

type DocumentCountResponse struct {
	TotalDocument int `json:"total_document"`
	TotalRejected int `json:"total_rejected"`
	TotalPending  int `json:"total_pending"`
	TotalApproved int `json:"total_approved"`
	TotalUsers    int `json:"total_users"`
}
