package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type FileSeeder struct{}

func (s FileSeeder) Run(db *gorm.DB) error {
	files := []model.File{
		{
			ID:          uint(1),
			FileName:    "Default Avatar",
			FilePath:    "profile/avatar.png",
			ContentType: "image/png",
		},
	}

	for _, file := range files {
		err := db.FirstOrCreate(&model.File{}, file).Error
		if err != nil {
			fmt.Printf("Error seeding file %s: %v\n", file.FileName, err)
			return err
		}
	}

	err := db.Exec("SELECT setval('files_id_seq', (SELECT MAX(id) FROM files))").Error
	if err != nil {
		fmt.Printf("Error resetting sequence for files: %v\n", err)
		return err
	}

	return nil
}
