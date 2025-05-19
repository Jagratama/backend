package model

import "time"

type ApprovalRequest struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	DocumentID uint      `json:"document_id" gorm:"not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	FileID     *uint     `json:"file_id"`
	Note       *string   `json:"note"`
	Status     string    `json:"status" gorm:"type:text;check:status IN ('pending','approved','rejected');default:'pending';not null"`
	ResolvedAt time.Time `json:"resolved_at"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;default:now()"`

	Document Document `json:"document" gorm:"foreignKey:DocumentID"`
	User     User     `json:"user" gorm:"foreignKey:UserID"`
	File     File     `json:"file" gorm:"foreignKey:FileID"`
}
