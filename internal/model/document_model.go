package model

import "time"

type Document struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CategoryID  uint      `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	FilePath    string    `json:"file"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
