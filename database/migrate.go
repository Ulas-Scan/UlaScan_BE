package database

import (
	"ulascan-be/entity"

	"gorm.io/gorm"
)

func MigrateFresh(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&entity.User{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}
	return nil
}
