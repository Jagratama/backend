package repository

import (
	"context"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	if err := r.db.WithContext(ctx).Create(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}
