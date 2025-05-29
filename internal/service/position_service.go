package service

import (
	"context"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type PositionService struct {
	positionRepository   repository.PositionRepository
	positionCategoryRule repository.PositionCategoryRuleRepository
}

func NewPositionService(positionRepository repository.PositionRepository, positionCategoryRule repository.PositionCategoryRuleRepository) *PositionService {
	return &PositionService{
		positionRepository:   positionRepository,
		positionCategoryRule: positionCategoryRule,
	}
}

func (s *PositionService) GetAllPositions(ctx context.Context) ([]*model.Position, error) {
	positions, err := s.positionRepository.GetAllPositions(ctx)
	if err != nil {
		return []*model.Position{}, err
	}
	return positions, nil
}

func (s *PositionService) GetPositionByID(ctx context.Context, id int) (*model.Position, error) {
	position, err := s.positionRepository.GetPositionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (s *PositionService) GetPositionsRequiredByCategoryID(ctx context.Context, categoryID int) ([]*model.Position, error) {
	var positions []*model.Position
	rules, err := s.positionCategoryRule.GetPositionsRuleByCategoryID(uint(categoryID))
	if err != nil {
		return []*model.Position{}, err
	}
	for _, rule := range rules {
		positions = append(positions, &rule.Position)
	}
	return positions, nil
}
