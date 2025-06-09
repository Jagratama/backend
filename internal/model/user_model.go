package model

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	ImageID      uint      `json:"image_id" gorm:"default:1;not null"`
	RoleID       uint      `json:"role_id"`
	PositionID   uint      `json:"position_id"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	Email        string    `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Password     string    `json:"password" gorm:"not null"`
	Organization string   `json:"organization" gorm:"size:100"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;default:now()"`

	Role     Role     `json:"role" gorm:"foreignKey:RoleID"`
	Position Position `json:"position" gorm:"foreignKey:PositionID"`
	File     File     `json:"image" gorm:"foreignKey:ImageID"`
}
