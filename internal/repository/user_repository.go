package repository

import (
	"context"
	"errors"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context, name string, position_id, page, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	query := r.db.WithContext(ctx)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if position_id > 0 {
		query = query.Where("position_id = ?", position_id)
	}

	err := query.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.
		Offset(offset).
		Limit(limit).
		Preload("Role").
		Preload("Position").
		Preload("File").
		Order("created_at DESC").
		Find(&users).Error

	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", id).Preload("Role").Preload("Position").Preload("File").First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).Where("email = ?", email).Preload("Role").Preload("Position").Preload("File").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.WithContext(ctx).Model(model.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CountAllUsers(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) GetApproverReviewerUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).
		Joins("JOIN roles ON roles.id = users.role_id").
		Where("roles.name IN (?)", []string{"approver", "reviewer"}).
		Preload("Role").
		Preload("Position").
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}
