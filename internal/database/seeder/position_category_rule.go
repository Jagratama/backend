package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionCategoryRuleSeeder struct{}

func (s PositionCategoryRuleSeeder) Run(db *gorm.DB) error {
	positionCategoryRules := []model.PositionCategoryRule{
		{ID: 1, CategoryID: 1, PositionID: 20, NeedSignature: true, Description: "kak blm"},
		{ID: 2, CategoryID: 1, PositionID: 6, NeedSignature: false, Description: "kak blm"},
		{ID: 3, CategoryID: 1, PositionID: 7, NeedSignature: false, Description: "kak blm"},
		{ID: 4, CategoryID: 1, PositionID: 8, NeedSignature: false, Description: "kak blm"},
		{ID: 5, CategoryID: 1, PositionID: 9, NeedSignature: false, Description: "kak blm"},
		{ID: 6, CategoryID: 1, PositionID: 10, NeedSignature: false, Description: "kak blm"},
		{ID: 7, CategoryID: 1, PositionID: 11, NeedSignature: true, Description: "kak blm"},
		{ID: 8, CategoryID: 2, PositionID: 5, NeedSignature: true, Description: "kak bem"},
		{ID: 9, CategoryID: 2, PositionID: 6, NeedSignature: false, Description: "kak bem"},
		{ID: 10, CategoryID: 2, PositionID: 7, NeedSignature: false, Description: "kak bem"},
		{ID: 11, CategoryID: 2, PositionID: 8, NeedSignature: false, Description: "kak bem"},
		{ID: 12, CategoryID: 2, PositionID: 9, NeedSignature: false, Description: "kak bem"},
		{ID: 13, CategoryID: 2, PositionID: 10, NeedSignature: false, Description: "kak bem"},
		{ID: 14, CategoryID: 2, PositionID: 11, NeedSignature: true, Description: "kak bem"},
		{ID: 15, CategoryID: 3, PositionID: 19, NeedSignature: true, Description: "kak hmj"},
		{ID: 16, CategoryID: 3, PositionID: 16, NeedSignature: false, Description: "kak hmj"},
		{ID: 17, CategoryID: 3, PositionID: 18, NeedSignature: true, Description: "kak hmj"},
		{ID: 18, CategoryID: 3, PositionID: 5, NeedSignature: false, Description: "kak hmj"},
		{ID: 19, CategoryID: 3, PositionID: 6, NeedSignature: false, Description: "kak hmj"},
		{ID: 20, CategoryID: 3, PositionID: 7, NeedSignature: false, Description: "kak hmj"},
		{ID: 21, CategoryID: 3, PositionID: 8, NeedSignature: false, Description: "kak hmj"},
		{ID: 22, CategoryID: 3, PositionID: 10, NeedSignature: false, Description: "kak hmj"},
		{ID: 23, CategoryID: 3, PositionID: 11, NeedSignature: true, Description: "kak hmj"},
		{ID: 24, CategoryID: 4, PositionID: 15, NeedSignature: true, Description: "kak hmps"},
		{ID: 25, CategoryID: 4, PositionID: 16, NeedSignature: false, Description: "kak hmps"},
		{ID: 26, CategoryID: 4, PositionID: 17, NeedSignature: false, Description: "kak hmps"},
		{ID: 27, CategoryID: 4, PositionID: 18, NeedSignature: false, Description: "kak hmps"},
		{ID: 28, CategoryID: 4, PositionID: 5, NeedSignature: false, Description: "kak hmps"},
		{ID: 29, CategoryID: 4, PositionID: 6, NeedSignature: false, Description: "kak hmps"},
		{ID: 30, CategoryID: 4, PositionID: 7, NeedSignature: false, Description: "kak hmps"},
		{ID: 31, CategoryID: 4, PositionID: 8, NeedSignature: false, Description: "kak hmps"},
		{ID: 32, CategoryID: 4, PositionID: 10, NeedSignature: false, Description: "kak hmps"},
		{ID: 33, CategoryID: 4, PositionID: 11, NeedSignature: true, Description: "kak hmps"},
		{ID: 34, CategoryID: 5, PositionID: 12, NeedSignature: true, Description: "kak ukm"},
		{ID: 35, CategoryID: 5, PositionID: 13, NeedSignature: false, Description: "kak ukm"},
		{ID: 36, CategoryID: 5, PositionID: 14, NeedSignature: false, Description: "kak ukm"},
		{ID: 37, CategoryID: 5, PositionID: 5, NeedSignature: true, Description: "kak ukm"},
		{ID: 38, CategoryID: 5, PositionID: 6, NeedSignature: false, Description: "kak ukm"},
		{ID: 39, CategoryID: 5, PositionID: 7, NeedSignature: false, Description: "kak ukm"},
		{ID: 40, CategoryID: 5, PositionID: 8, NeedSignature: false, Description: "kak ukm"},
		{ID: 41, CategoryID: 5, PositionID: 10, NeedSignature: false, Description: "kak ukm"},
		{ID: 42, CategoryID: 5, PositionID: 11, NeedSignature: true, Description: "kak ukm"},
		{ID: 43, CategoryID: 6, PositionID: 3, NeedSignature: true, Description: "kak sbh"},
		{ID: 44, CategoryID: 6, PositionID: 4, NeedSignature: false, Description: "kak sbh"},
		{ID: 45, CategoryID: 6, PositionID: 5, NeedSignature: false, Description: "kak sbh"},
		{ID: 46, CategoryID: 6, PositionID: 6, NeedSignature: false, Description: "kak sbh"},
		{ID: 47, CategoryID: 6, PositionID: 7, NeedSignature: false, Description: "kak sbh"},
		{ID: 48, CategoryID: 6, PositionID: 8, NeedSignature: false, Description: "kak sbh"},
		{ID: 49, CategoryID: 6, PositionID: 10, NeedSignature: false, Description: "kak sbh"},
		{ID: 50, CategoryID: 6, PositionID: 11, NeedSignature: true, Description: "kak sbh"},
		{ID: 51, CategoryID: 7, PositionID: 20, NeedSignature: true, Description: "surat blm"},
		{ID: 52, CategoryID: 7, PositionID: 6, NeedSignature: false, Description: "surat blm"},
		{ID: 53, CategoryID: 7, PositionID: 7, NeedSignature: false, Description: "surat blm"},
		{ID: 54, CategoryID: 7, PositionID: 8, NeedSignature: false, Description: "surat blm"},
		{ID: 55, CategoryID: 7, PositionID: 9, NeedSignature: true, Description: "surat blm"},
		{ID: 56, CategoryID: 8, PositionID: 5, NeedSignature: true, Description: "surat bem"},
		{ID: 57, CategoryID: 8, PositionID: 6, NeedSignature: false, Description: "surat bem"},
		{ID: 58, CategoryID: 8, PositionID: 7, NeedSignature: false, Description: "surat bem"},
		{ID: 59, CategoryID: 8, PositionID: 8, NeedSignature: false, Description: "surat bem"},
		{ID: 60, CategoryID: 8, PositionID: 9, NeedSignature: true, Description: "surat bem"},
		{ID: 61, CategoryID: 9, PositionID: 19, NeedSignature: true, Description: "surat hmj"},
		{ID: 62, CategoryID: 9, PositionID: 16, NeedSignature: false, Description: "surat hmj"},
		{ID: 63, CategoryID: 9, PositionID: 5, NeedSignature: false, Description: "surat hmj"},
		{ID: 64, CategoryID: 9, PositionID: 6, NeedSignature: false, Description: "surat hmj"},
		{ID: 65, CategoryID: 9, PositionID: 7, NeedSignature: false, Description: "surat hmj"},
		{ID: 66, CategoryID: 9, PositionID: 8, NeedSignature: false, Description: "surat hmj"},
		{ID: 67, CategoryID: 9, PositionID: 9, NeedSignature: true, Description: "surat hmj"},
		{ID: 68, CategoryID: 10, PositionID: 15, NeedSignature: true, Description: "surat hmps"},
		{ID: 69, CategoryID: 10, PositionID: 16, NeedSignature: false, Description: "surat hmps"},
		{ID: 70, CategoryID: 10, PositionID: 5, NeedSignature: false, Description: "surat hmps"},
		{ID: 71, CategoryID: 10, PositionID: 6, NeedSignature: false, Description: "surat hmps"},
		{ID: 72, CategoryID: 10, PositionID: 7, NeedSignature: false, Description: "surat hmps"},
		{ID: 73, CategoryID: 10, PositionID: 8, NeedSignature: false, Description: "surat hmps"},
		{ID: 74, CategoryID: 10, PositionID: 9, NeedSignature: true, Description: "surat hmps"},
		{ID: 75, CategoryID: 11, PositionID: 12, NeedSignature: true, Description: "surat ukm"},
		{ID: 76, CategoryID: 11, PositionID: 13, NeedSignature: false, Description: "surat ukm"},
		{ID: 77, CategoryID: 11, PositionID: 14, NeedSignature: false, Description: "surat ukm"},
		{ID: 78, CategoryID: 11, PositionID: 6, NeedSignature: false, Description: "surat ukm"},
		{ID: 79, CategoryID: 11, PositionID: 7, NeedSignature: false, Description: "surat ukm"},
		{ID: 80, CategoryID: 11, PositionID: 8, NeedSignature: false, Description: "surat ukm"},
		{ID: 81, CategoryID: 11, PositionID: 9, NeedSignature: true, Description: "surat ukm"},
		{ID: 82, CategoryID: 12, PositionID: 19, NeedSignature: true, Description: "surat sbh"},
		{ID: 83, CategoryID: 12, PositionID: 4, NeedSignature: false, Description: "surat sbh"},
		{ID: 84, CategoryID: 12, PositionID: 5, NeedSignature: false, Description: "surat sbh"},
		{ID: 85, CategoryID: 12, PositionID: 6, NeedSignature: false, Description: "surat sbh"},
		{ID: 86, CategoryID: 12, PositionID: 7, NeedSignature: false, Description: "surat sbh"},
		{ID: 87, CategoryID: 12, PositionID: 8, NeedSignature: false, Description: "surat sbh"},
		{ID: 88, CategoryID: 12, PositionID: 9, NeedSignature: true, Description: "surat sbh"},
	}

	for _, positionCategoryRule := range positionCategoryRules {
		err := db.FirstOrCreate(&model.PositionCategoryRule{}, positionCategoryRule).Error
		if err != nil {
			fmt.Printf("Error seeding position category rules %d: %v\n", positionCategoryRule.ID, err)
			return err
		}
	}
	return nil
}
