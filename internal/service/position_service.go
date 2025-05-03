package service

import (
	"context"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type PositionService struct {
	positionRepository repository.PositionRepository
}

func NewPositionService(positionRepository repository.PositionRepository) *PositionService {
	return &PositionService{
		positionRepository: positionRepository,
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
