package model

import "time"

type User struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	RoleID     uint       `json:"role_id"`
	PositionID uint       `json:"position_id"`
	Name       string     `json:"name" gorm:"size:100;not null"`
	Birthdate  *time.Time `json:"birthdate" gorm:"type:date"`
	Email      string     `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Password   string     `json:"_" gorm:"not null"`
	ImagePath  string     `json:"image_path"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
