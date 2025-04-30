package service

import (
	"context"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/repository"
)

type RoleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) GetRoleByID(ctx context.Context, id int) (*model.Role, error) {
	role, err := s.roleRepository.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) GetAllRoles(ctx context.Context) ([]*model.Role, error) {
	roles, err := s.roleRepository.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
