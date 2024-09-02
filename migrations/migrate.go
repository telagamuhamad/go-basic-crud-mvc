package migrations

import (
	"log"
	"testing-app/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
