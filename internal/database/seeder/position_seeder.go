package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionSeeder struct{}

func (s PositionSeeder) Run(db *gorm.DB) error {
	positions := []model.Position{
		{
			ID:                 uint(1),
			Name:               "boss",
			RequiresSignatures: true,
		},
		{
			ID:                 uint(2),
			Name:               "staff",
			RequiresSignatures: false,
		},
		{
			ID:                 uint(3),
			Name:               "assistant",
			RequiresSignatures: true,
		},
	}

	for _, position := range positions {
		err := db.FirstOrCreate(&model.Position{}, position).Error
		if err != nil {
			fmt.Printf("Error seeding position %s: %v\n", position.Name, err)
			return err
		}
	}
	return nil
}
