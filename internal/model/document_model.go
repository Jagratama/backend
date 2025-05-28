package model

import "time"

type Document struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	UserID          uint       `json:"user_id"`
	AddressedUserID uint       `json:"addressed_user_id"`
	FileID          uint       `json:"file_id"`
	CategoryID      uint       `json:"category_id"`
	Title           string     `json:"title"`
	Slug            string     `json:"slug"`
	Description     string     `json:"description"`
	Confirmed       bool       `json:"confirmed" gorm:"default:false;not null"`
	LastStatus      string     `json:"last_status" gorm:"type:text;check:last_status IN ('pending','approved','rejected');default:'pending';not null"`
	ApprovedAt      *time.Time `json:"approved_at" gorm:"default:null"`
	CreatedAt       time.Time  `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"type:timestamp;default:now()"`

	User          User     `json:"user" gorm:"foreignKey:UserID"`
	AddressedUser User     `json:"addressed_user" gorm:"foreignKey:AddressedUserID"`
	Category      Category `json:"category" gorm:"foreignKey:CategoryID"`
	File          File     `json:"file" gorm:"foreignKey:FileID"`
}
