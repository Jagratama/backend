package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionSeeder struct{}

func (s PositionSeeder) Run(db *gorm.DB) error {
	positions := []model.Position{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "Pengaju"},
		{ID: 3, Name: "Ketua SBH"},
		{ID: 4, Name: "PEMBINA SBH"},
		{ID: 5, Name: "PRESIDEN BEM"},
		{ID: 6, Name: "KOMISI B BLM"},
		{ID: 7, Name: "PENANGGUNG JAWAB MAHASISWA DAN ALUMNI"},
		{ID: 8, Name: "KA SUB BAG ADM AKADEMIK"},
		{ID: 9, Name: "KA BAG ADM AKADEMIK UMUM"},
		{ID: 10, Name: "WAKIL DIREKTUR III"},
		{ID: 11, Name: "DIREKTUR"},
		{ID: 12, Name: "Ketua UKM"},
		{ID: 13, Name: "PEMBINA UKM"},
		{ID: 14, Name: "MENTERI MINBA BEM"},
		{ID: 15, Name: "Ketua HMPS"},
		{ID: 16, Name: "PJ MAHASISWA DAN ALUMNI JURUSAN"},
		{ID: 17, Name: "KAPRODI"},
		{ID: 18, Name: "KAJUR"},
		{ID: 19, Name: "Ketua HMJ"},
		{ID: 20, Name: "KETUA BLM"},
	}

	for _, position := range positions {
		err := db.FirstOrCreate(&model.Position{}, position).Error
		if err != nil {
			fmt.Printf("Error seeding position %s: %v\n", position.Name, err)
			return err
		}
	}

	err := db.Exec("SELECT setval('positions_id_seq', (SELECT MAX(id) FROM positions))").Error
	if err != nil {
		fmt.Printf("Error resetting positions ID sequence: %v\n", err)
		return err
	}

	return nil
}
