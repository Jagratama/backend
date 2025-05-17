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
			Name: "KAK/ LPJ",
			Type: "kak/lpj",
		},
		{
			ID:   uint(2),
			Name: "Persuratan",
			Type: "persuratan",
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
