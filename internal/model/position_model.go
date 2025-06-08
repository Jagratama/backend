package model

import (
	"time"
)

type Position struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:now()"`
}
