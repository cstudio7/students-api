package database

import (
	"students-api/internal/services/student"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&student.Student{}); err != nil {
		return err
	}
	return nil
}
