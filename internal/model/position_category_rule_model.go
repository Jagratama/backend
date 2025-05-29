package model

type PositionCategoryRule struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CategoryID uint `json:"category_id" gorm:"not null"`
	PositionID uint `json:"position_id" gorm:"not null"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Position Position `json:"position" gorm:"foreignKey:PositionID"`
}
