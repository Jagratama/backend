// database/seeder/seeder.go
package seeder

import "gorm.io/gorm"

type Seeder interface {
	Run(db *gorm.DB) error
}
