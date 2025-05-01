package service

import (
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	categories, err := s.categoryRepository.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
