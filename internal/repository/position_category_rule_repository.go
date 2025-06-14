package repository

import (
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionCategoryRuleRepository struct {
	db *gorm.DB
}

func NewPositionCategoryRuleRepository(db *gorm.DB) *PositionCategoryRuleRepository {
	return &PositionCategoryRuleRepository{
		db: db,
	}
}

func (r *PositionCategoryRuleRepository) GetPositionsRuleByCategoryID(categoryID uint) ([]*model.PositionCategoryRule, error) {
	var rules []*model.PositionCategoryRule
	if err := r.db.Where("category_id = ?", categoryID).Preload("Position").Order("display_order asc").Find(&rules).Error; err != nil {
		return nil, err
	}
	return rules, nil
}

func (r *PositionCategoryRuleRepository) GetPositionRuleByCategoryIDAndPositionID(categoryID, positionID uint) (*model.PositionCategoryRule, error) {
	var rule model.PositionCategoryRule
	if err := r.db.Where("category_id = ? AND position_id = ?", categoryID, positionID).Preload("Position").First(&rule).Error; err != nil {
		return nil, err
	}
	return &rule, nil
}
