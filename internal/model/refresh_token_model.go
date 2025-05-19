package model

import "time"

type RefreshToken struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	UserAgent string    `json:"user_agent"`
	ExpiredAt string    `json:"expired_at"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:now()"`
}
