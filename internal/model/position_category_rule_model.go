package model

type PositionCategoryRule struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	CategoryID    uint   `json:"category_id" gorm:"not null"`
	PositionID    uint   `json:"position_id" gorm:"not null"`
	DisplayOrder  int    `json:"display_order" gorm:"not null"`
	NeedSignature bool   `json:"need_signature" gorm:"default:false;not null"`
	Description   string `json:"description" gorm:"type:text"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Position Position `json:"position" gorm:"foreignKey:PositionID"`
}
