package config

import (
	"fmt"
	"jagratama-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	mode := GetEnv("APP_ENV", "development")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", GetEnv("DB_HOST", "localhost"), GetEnv("DB_USER", ""), GetEnv("DB_NAME", ""), GetEnv("DB_PORT", ""))
	if mode == "production" {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=require TimeZone=Asia/Jakarta", GetEnv("DB_HOST", "localhost"), GetEnv("DB_USER", ""), GetEnv("DB_NAME", ""), GetEnv("DB_PORT", ""), GetEnv("DB_PASSWORD", ""))
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	mode := GetEnv("APP_ENV", "development")
	if mode == "production" {
		err := db.AutoMigrate(&model.User{}, &model.RefreshToken{}, &model.Document{}, &model.ApprovalRequest{}, &model.PositionCategoryRule{})
		if err != nil {
			return fmt.Errorf("failed to migrate database: %w", err)
		}
		fmt.Println("Successfully migrated database")
	}
	return nil
}
