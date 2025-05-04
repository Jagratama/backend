package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type RoleSeeder struct{}

func (s RoleSeeder) Run(db *gorm.DB) error {
	roles := []model.Role{
		{
			ID:   uint(1),
			Name: "admin",
		},
		{
			ID:   uint(2),
			Name: "reviewer",
		},
		{
			ID:   uint(3),
			Name: "approver",
		},
		{
			ID:   uint(4),
			Name: "requester",
		},
	}

	for _, role := range roles {
		err := db.FirstOrCreate(&model.Role{}, role).Error
		if err != nil {
			fmt.Printf("Error seeding role %s: %v\n", role.Name, err)
			return err
		}
	}
	return nil
}
