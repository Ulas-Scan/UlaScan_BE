package database

import (
	"ulascan-be/entity"

	"gorm.io/gorm"
)

func MigrateFresh(db *gorm.DB) error {
	// Drop the tables if they exist
	if err := db.Migrator().DropTable(
		&entity.History{},
		&entity.User{},
	); err != nil {
		return err
	}

	// Auto-migrate the tables
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.History{},
	); err != nil {
		return err
	}

	return nil
}
