package model

import "time"

type ApprovalRequest struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	DocumentID uint       `json:"document_id" gorm:"not null"`
	UserID     uint       `json:"user_id" gorm:"not null"`
	FilePath   string     `json:"file_path"`
	Note       *string    `json:"note"`
	Status     string     `json:"status" gorm:"type:text;check:status IN ('pending','approved','rejected');default:'pending';not null"`
	ResolvedAt *time.Time `json:"resolved_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	Document Document `json:"document" gorm:"foreignKey:DocumentID"`
	User     User     `json:"user" gorm:"foreignKey:UserID"`
}
