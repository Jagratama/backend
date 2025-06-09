package seeder

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/gorm"
)

type PositionCategoryRuleSeeder struct{}

func (s PositionCategoryRuleSeeder) Run(db *gorm.DB) error {
	positionCategoryRules := []model.PositionCategoryRule{
		{ID: 1, CategoryID: 1, PositionID: 20, NeedSignature: true, Description: "kak blm", DisplayOrder: 1},
		{ID: 2, CategoryID: 1, PositionID: 6, NeedSignature: false, Description: "kak blm", DisplayOrder: 2},
		{ID: 3, CategoryID: 1, PositionID: 7, NeedSignature: false, Description: "kak blm", DisplayOrder: 3},
		{ID: 4, CategoryID: 1, PositionID: 8, NeedSignature: false, Description: "kak blm", DisplayOrder: 4},
		{ID: 5, CategoryID: 1, PositionID: 9, NeedSignature: false, Description: "kak blm", DisplayOrder: 6},
		{ID: 6, CategoryID: 1, PositionID: 10, NeedSignature: false, Description: "kak blm", DisplayOrder: 7},
		{ID: 7, CategoryID: 1, PositionID: 11, NeedSignature: true, Description: "kak blm", DisplayOrder: 8},
		{ID: 8, CategoryID: 2, PositionID: 5, NeedSignature: true, Description: "kak bem", DisplayOrder: 1},
		{ID: 9, CategoryID: 2, PositionID: 6, NeedSignature: false, Description: "kak bem", DisplayOrder: 2},
		{ID: 10, CategoryID: 2, PositionID: 7, NeedSignature: false, Description: "kak bem", DisplayOrder: 3},
		{ID: 11, CategoryID: 2, PositionID: 8, NeedSignature: false, Description: "kak bem", DisplayOrder: 4},
		{ID: 12, CategoryID: 2, PositionID: 9, NeedSignature: false, Description: "kak bem", DisplayOrder: 5},
		{ID: 13, CategoryID: 2, PositionID: 10, NeedSignature: false, Description: "kak bem", DisplayOrder: 6},
		{ID: 14, CategoryID: 2, PositionID: 11, NeedSignature: true, Description: "kak bem", DisplayOrder: 7},
		{ID: 15, CategoryID: 3, PositionID: 19, NeedSignature: true, Description: "kak hmj", DisplayOrder: 1},
		{ID: 16, CategoryID: 3, PositionID: 16, NeedSignature: false, Description: "kak hmj", DisplayOrder: 2},
		{ID: 17, CategoryID: 3, PositionID: 18, NeedSignature: true, Description: "kak hmj", DisplayOrder: 3},
		{ID: 18, CategoryID: 3, PositionID: 5, NeedSignature: false, Description: "kak hmj", DisplayOrder: 4},
		{ID: 19, CategoryID: 3, PositionID: 6, NeedSignature: false, Description: "kak hmj", DisplayOrder: 5},
		{ID: 20, CategoryID: 3, PositionID: 7, NeedSignature: false, Description: "kak hmj", DisplayOrder: 6},
		{ID: 21, CategoryID: 3, PositionID: 8, NeedSignature: false, Description: "kak hmj", DisplayOrder: 7},
		{ID: 22, CategoryID: 3, PositionID: 10, NeedSignature: false, Description: "kak hmj", DisplayOrder: 8},
		{ID: 23, CategoryID: 3, PositionID: 11, NeedSignature: true, Description: "kak hmj", DisplayOrder: 9},
		{ID: 24, CategoryID: 4, PositionID: 15, NeedSignature: true, Description: "kak hmps", DisplayOrder: 1},
		{ID: 25, CategoryID: 4, PositionID: 16, NeedSignature: false, Description: "kak hmps", DisplayOrder: 2},
		{ID: 26, CategoryID: 4, PositionID: 17, NeedSignature: false, Description: "kak hmps", DisplayOrder: 3},
		{ID: 27, CategoryID: 4, PositionID: 18, NeedSignature: false, Description: "kak hmps", DisplayOrder: 4},
		{ID: 28, CategoryID: 4, PositionID: 5, NeedSignature: false, Description: "kak hmps", DisplayOrder: 5},
		{ID: 29, CategoryID: 4, PositionID: 6, NeedSignature: false, Description: "kak hmps", DisplayOrder: 6},
		{ID: 30, CategoryID: 4, PositionID: 7, NeedSignature: false, Description: "kak hmps", DisplayOrder: 7},
		{ID: 31, CategoryID: 4, PositionID: 8, NeedSignature: false, Description: "kak hmps", DisplayOrder: 8},
		{ID: 32, CategoryID: 4, PositionID: 10, NeedSignature: false, Description: "kak hmps", DisplayOrder: 9},
		{ID: 33, CategoryID: 4, PositionID: 11, NeedSignature: true, Description: "kak hmps", DisplayOrder: 10},
		{ID: 34, CategoryID: 5, PositionID: 12, NeedSignature: true, Description: "kak ukm", DisplayOrder: 1},
		{ID: 35, CategoryID: 5, PositionID: 13, NeedSignature: false, Description: "kak ukm", DisplayOrder: 2},
		{ID: 36, CategoryID: 5, PositionID: 14, NeedSignature: false, Description: "kak ukm", DisplayOrder: 3},
		{ID: 37, CategoryID: 5, PositionID: 5, NeedSignature: true, Description: "kak ukm", DisplayOrder: 4},
		{ID: 38, CategoryID: 5, PositionID: 6, NeedSignature: false, Description: "kak ukm", DisplayOrder: 5},
		{ID: 39, CategoryID: 5, PositionID: 7, NeedSignature: false, Description: "kak ukm", DisplayOrder: 6},
		{ID: 40, CategoryID: 5, PositionID: 8, NeedSignature: false, Description: "kak ukm", DisplayOrder: 7},
		{ID: 41, CategoryID: 5, PositionID: 10, NeedSignature: false, Description: "kak ukm", DisplayOrder: 8},
		{ID: 42, CategoryID: 5, PositionID: 11, NeedSignature: true, Description: "kak ukm", DisplayOrder: 9},
		{ID: 43, CategoryID: 6, PositionID: 3, NeedSignature: true, Description: "kak sbh", DisplayOrder: 1},
		{ID: 44, CategoryID: 6, PositionID: 4, NeedSignature: false, Description: "kak sbh", DisplayOrder: 2},
		{ID: 45, CategoryID: 6, PositionID: 5, NeedSignature: false, Description: "kak sbh", DisplayOrder: 3},
		{ID: 46, CategoryID: 6, PositionID: 6, NeedSignature: false, Description: "kak sbh", DisplayOrder: 4},
		{ID: 47, CategoryID: 6, PositionID: 7, NeedSignature: false, Description: "kak sbh", DisplayOrder: 5},
		{ID: 48, CategoryID: 6, PositionID: 8, NeedSignature: false, Description: "kak sbh", DisplayOrder: 6},
		{ID: 49, CategoryID: 6, PositionID: 10, NeedSignature: false, Description: "kak sbh", DisplayOrder: 7},
		{ID: 50, CategoryID: 6, PositionID: 11, NeedSignature: true, Description: "kak sbh", DisplayOrder: 8},
		{ID: 51, CategoryID: 7, PositionID: 20, NeedSignature: true, Description: "surat blm", DisplayOrder: 1},
		{ID: 52, CategoryID: 7, PositionID: 6, NeedSignature: false, Description: "surat blm", DisplayOrder: 2},
		{ID: 53, CategoryID: 7, PositionID: 7, NeedSignature: false, Description: "surat blm", DisplayOrder: 3},
		{ID: 54, CategoryID: 7, PositionID: 8, NeedSignature: false, Description: "surat blm", DisplayOrder: 4},
		{ID: 55, CategoryID: 7, PositionID: 9, NeedSignature: true, Description: "surat blm", DisplayOrder: 5},
		{ID: 56, CategoryID: 8, PositionID: 5, NeedSignature: true, Description: "surat bem", DisplayOrder: 1},
		{ID: 57, CategoryID: 8, PositionID: 6, NeedSignature: false, Description: "surat bem", DisplayOrder: 2},
		{ID: 58, CategoryID: 8, PositionID: 7, NeedSignature: false, Description: "surat bem", DisplayOrder: 3},
		{ID: 59, CategoryID: 8, PositionID: 8, NeedSignature: false, Description: "surat bem", DisplayOrder: 4},
		{ID: 60, CategoryID: 8, PositionID: 9, NeedSignature: true, Description: "surat bem", DisplayOrder: 5},
		{ID: 61, CategoryID: 9, PositionID: 19, NeedSignature: true, Description: "surat hmj", DisplayOrder: 1},
		{ID: 62, CategoryID: 9, PositionID: 16, NeedSignature: false, Description: "surat hmj", DisplayOrder: 2},
		{ID: 63, CategoryID: 9, PositionID: 5, NeedSignature: false, Description: "surat hmj", DisplayOrder: 3},
		{ID: 64, CategoryID: 9, PositionID: 6, NeedSignature: false, Description: "surat hmj", DisplayOrder: 4},
		{ID: 65, CategoryID: 9, PositionID: 7, NeedSignature: false, Description: "surat hmj", DisplayOrder: 5},
		{ID: 66, CategoryID: 9, PositionID: 8, NeedSignature: false, Description: "surat hmj", DisplayOrder: 6},
		{ID: 67, CategoryID: 9, PositionID: 9, NeedSignature: true, Description: "surat hmj", DisplayOrder: 7},
		{ID: 68, CategoryID: 10, PositionID: 15, NeedSignature: true, Description: "surat hmps", DisplayOrder: 1},
		{ID: 69, CategoryID: 10, PositionID: 16, NeedSignature: false, Description: "surat hmps", DisplayOrder: 2},
		{ID: 70, CategoryID: 10, PositionID: 5, NeedSignature: false, Description: "surat hmps", DisplayOrder: 3},
		{ID: 71, CategoryID: 10, PositionID: 6, NeedSignature: false, Description: "surat hmps", DisplayOrder: 4},
		{ID: 72, CategoryID: 10, PositionID: 7, NeedSignature: false, Description: "surat hmps", DisplayOrder: 5},
		{ID: 73, CategoryID: 10, PositionID: 8, NeedSignature: false, Description: "surat hmps", DisplayOrder: 6},
		{ID: 74, CategoryID: 10, PositionID: 9, NeedSignature: true, Description: "surat hmps", DisplayOrder: 7},
		{ID: 75, CategoryID: 11, PositionID: 12, NeedSignature: true, Description: "surat ukm", DisplayOrder: 1},
		{ID: 76, CategoryID: 11, PositionID: 13, NeedSignature: false, Description: "surat ukm", DisplayOrder: 2},
		{ID: 77, CategoryID: 11, PositionID: 14, NeedSignature: false, Description: "surat ukm", DisplayOrder: 3},
		{ID: 78, CategoryID: 11, PositionID: 6, NeedSignature: false, Description: "surat ukm", DisplayOrder: 4},
		{ID: 79, CategoryID: 11, PositionID: 7, NeedSignature: false, Description: "surat ukm", DisplayOrder: 5},
		{ID: 80, CategoryID: 11, PositionID: 8, NeedSignature: false, Description: "surat ukm", DisplayOrder: 6},
		{ID: 81, CategoryID: 11, PositionID: 9, NeedSignature: true, Description: "surat ukm", DisplayOrder: 7},
		{ID: 82, CategoryID: 12, PositionID: 19, NeedSignature: true, Description: "surat sbh", DisplayOrder: 1},
		{ID: 83, CategoryID: 12, PositionID: 4, NeedSignature: false, Description: "surat sbh", DisplayOrder: 2},
		{ID: 84, CategoryID: 12, PositionID: 5, NeedSignature: false, Description: "surat sbh", DisplayOrder: 3},
		{ID: 85, CategoryID: 12, PositionID: 6, NeedSignature: false, Description: "surat sbh", DisplayOrder: 4},
		{ID: 86, CategoryID: 12, PositionID: 7, NeedSignature: false, Description: "surat sbh", DisplayOrder: 5},
		{ID: 87, CategoryID: 12, PositionID: 8, NeedSignature: false, Description: "surat sbh", DisplayOrder: 6},
		{ID: 88, CategoryID: 12, PositionID: 9, NeedSignature: true, Description: "surat sbh", DisplayOrder: 7},
	}

	for _, positionCategoryRule := range positionCategoryRules {
		err := db.FirstOrCreate(&model.PositionCategoryRule{}, positionCategoryRule).Error
		if err != nil {
			fmt.Printf("Error seeding position category rules %d: %v\n", positionCategoryRule.ID, err)
			return err
		}
	}

	err := db.Exec("SELECT setval('position_category_rules_id_seq', (SELECT MAX(id) FROM position_category_rules))").Error
	if err != nil {
		fmt.Printf("Error resetting sequence for position_category_rules: %v\n", err)
		return err
	}

	return nil
}
