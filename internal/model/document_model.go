package model

import "time"

type Document struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FileID      uint      `json:"file_id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	LastStatus  string    `json:"last_status" gorm:"type:text;check:last_status IN ('pending','approved','rejected');default:'pending';not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User     User     `json:"user" gorm:"foreignKey:UserID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	File     File     `json:"file" gorm:"foreignKey:FileID"`
}
