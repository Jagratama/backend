package repository

import (
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAllCategories() ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.Order("id ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
