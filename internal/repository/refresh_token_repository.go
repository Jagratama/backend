package repository

import (
	"context"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, refreshToken *model.RefreshToken) error {
	return r.db.Create(refreshToken).Error
}

func (r *RefreshTokenRepository) GetByUserID(ctx context.Context, userID int) (*model.RefreshToken, error) {
	refreshToken := &model.RefreshToken{}
	err := r.db.Where("user_id = ?", userID).First(refreshToken).Error
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (r *RefreshTokenRepository) GetByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	refreshToken := &model.RefreshToken{}
	err := r.db.Where("token = ?", token).First(refreshToken).Error
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (r *RefreshTokenRepository) DeleteByUserID(ctx context.Context, userID int) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error
}

func (r *RefreshTokenRepository) DeleteByToken(ctx context.Context, token string) error {
	return r.db.Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}
