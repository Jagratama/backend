package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type CategorySeeder struct{}

func (s CategorySeeder) Run(db *gorm.DB) error {
	categories := []model.Category{
		{ID: uint(1), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (BLM)", Type: "kak/lpj"},
		{ID: uint(2), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (BEM)", Type: "kak/lpj"},
		{ID: uint(3), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (HMJ)", Type: "kak/lpj"},
		{ID: uint(4), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (HMPS)", Type: "kak/lpj"},
		{ID: uint(5), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (UKM)", Type: "kak/lpj"},
		{ID: uint(6), Name: "ALUR TOR, KAK/LPJ/PROPOSAL SPONSORSHIP (SBH)", Type: "kak/lpj"},
		{ID: uint(7), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (BLM)", Type: "persuratan"},
		{ID: uint(8), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (BEM)", Type: "persuratan"},
		{ID: uint(9), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (HMJ)", Type: "persuratan"},
		{ID: uint(10), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (HMPS)", Type: "persuratan"},
		{ID: uint(11), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (UKM)", Type: "persuratan"},
		{ID: uint(12), Name: "ALUR PENGAJUAN PERSURATAN DIREKTORAT DAN DESAIN SERTIFIKAT (SBH)", Type: "persuratan"},
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
