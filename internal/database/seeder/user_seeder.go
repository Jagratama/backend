package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Run(db *gorm.DB) error {
	imageID := uint(1) // Assuming you have a default image ID
	users := []model.User{
		{
			Name:       "admin",
			Email:      "admin@gmail.com",
			Password:   "$2a$10$478oXQZ.RqD45tCn8GW1RONMLlTE5G.8FxClOIgLF/zcBxSerdTju",
			ImageID:    &imageID,
			RoleID:     1,
			PositionID: 1,
		},
		{
			Name:       "reviewer",
			Email:      "reviewer@gmail.com",
			Password:   "$2a$10$478oXQZ.RqD45tCn8GW1RONMLlTE5G.8FxClOIgLF/zcBxSerdTju",
			ImageID:    &imageID,
			RoleID:     2,
			PositionID: 1,
		},
		{
			Name:       "approver",
			Email:      "approver@gmail.com",
			Password:   "$2a$10$478oXQZ.RqD45tCn8GW1RONMLlTE5G.8FxClOIgLF/zcBxSerdTju",
			ImageID:    &imageID,
			RoleID:     3,
			PositionID: 2,
		},
		{
			Name:       "pengaju",
			Email:      "pengaju@gmail.com",
			Password:   "$2a$10$478oXQZ.RqD45tCn8GW1RONMLlTE5G.8FxClOIgLF/zcBxSerdTju",
			ImageID:    &imageID,
			RoleID:     4,
			PositionID: 3,
		},
	}

	for _, user := range users {
		err := db.FirstOrCreate(&model.User{}, user).Error
		if err != nil {
			fmt.Printf("Error seeding user %s: %v\n", user.Name, err)
			return err
		}
	}
	return nil
}
