package repository

import (
	"context"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{
		db: db,
	}
}
func (r *PositionRepository) GetAllPositions(ctx context.Context) ([]*model.Position, error) {
	var positions []*model.Position
	err := r.db.Order("name ASC").Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}
func (r *PositionRepository) GetPositionByID(ctx context.Context, id int) (*model.Position, error) {
	var position model.Position
	err := r.db.First(&position, id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}
