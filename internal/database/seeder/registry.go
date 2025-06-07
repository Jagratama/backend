// database/seeder/registry.go
package seeder

import "gorm.io/gorm"

func RunAll(db *gorm.DB) error {
	seeders := []Seeder{
		FileSeeder{},
		RoleSeeder{},
		PositionSeeder{},
		UserSeeder{},
		CategorySeeder{},
		PositionCategoryRuleSeeder{},
	}

	for _, s := range seeders {
		if err := s.Run(db); err != nil {
			return err
		}
	}
	return nil
}
