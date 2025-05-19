package model

import (
	"time"

	"github.com/lib/pq"
)

type Position struct {
	ID                              uint           `json:"id" gorm:"primaryKey"`
	Name                            string         `json:"name"`
	RequiresSignatureByCategoryType pq.StringArray `json:"requires_signature_by_category_type" gorm:"type:text[]"`
	CreatedAt                       time.Time      `json:"created_at" gorm:"type:timestamp;default:now()"`
	UpdatedAt                       time.Time      `json:"updated_at" gorm:"type:timestamp;default:now()"`
}
