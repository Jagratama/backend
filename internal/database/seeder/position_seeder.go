package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PositionSeeder struct{}

func (s PositionSeeder) Run(db *gorm.DB) error {
	positions := []model.Position{
		{ID: 1, Name: "Admin", RequiresSignatureByCategoryType: nil},
		{ID: 2, Name: "Pengaju", RequiresSignatureByCategoryType: nil},
		{ID: 3, Name: "Ketua SBH", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj", "persuratan"}},
		{ID: 4, Name: "PEMBINA SBH", RequiresSignatureByCategoryType: nil},
		{ID: 5, Name: "PRESIDEN BEM", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj"}},
		{ID: 6, Name: "KOMISI B BLM", RequiresSignatureByCategoryType: nil},
		{ID: 7, Name: "PENANGGUNG JAWAB MAHASISWA DAN ALUMNI", RequiresSignatureByCategoryType: nil},
		{ID: 8, Name: "KA SUB BAG ADM AKADEMIK", RequiresSignatureByCategoryType: pq.StringArray{"persuratan"}},
		{ID: 9, Name: "KA BAG ADM AKADEMIK UMUM", RequiresSignatureByCategoryType: pq.StringArray{"persuratan"}},
		{ID: 10, Name: "WAKIL DIREKTUR III", RequiresSignatureByCategoryType: nil},
		{ID: 11, Name: "DIREKTUR", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj"}},
		{ID: 12, Name: "Ketua UKM", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj", "persuratan"}},
		{ID: 13, Name: "PEMBINA UKM", RequiresSignatureByCategoryType: nil},
		{ID: 14, Name: "MENTERI MINBA BEM", RequiresSignatureByCategoryType: nil},
		{ID: 15, Name: "Ketua HMPS", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj", "persuratan"}},
		{ID: 16, Name: "PJ MAHASISWA DAN ALUMNI JURUSAN", RequiresSignatureByCategoryType: nil},
		{ID: 17, Name: "KAPRODI", RequiresSignatureByCategoryType: nil},
		{ID: 18, Name: "KAJUR", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj"}},
		{ID: 19, Name: "Ketua HMJ", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj", "persuratan"}},
		{ID: 20, Name: "KETUA BLM", RequiresSignatureByCategoryType: pq.StringArray{"kak/lpj", "persuratan"}},
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
