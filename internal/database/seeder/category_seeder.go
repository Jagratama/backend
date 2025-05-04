package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type CategorySeeder struct{}

func (s CategorySeeder) Run(db *gorm.DB) error {
	categories := []model.Category{
		{
			ID:   uint(1),
			Name: "Cat 1",
		},
		{
			ID:   uint(2),
			Name: "Cat 2",
		},
		{
			ID:   uint(3),
			Name: "Cat 3",
		},
	}

	for _, category := range categories {
		err := db.FirstOrCreate(&model.Category{}, category).Error
		if err != nil {
			fmt.Printf("Error seeding category %s: %v\n", category.Name, err)
			return err
		}
	}
	return nil
}
