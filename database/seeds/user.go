package seeds

import (
	"github.com/javakanaya/ulascan-be/entity"

	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	var userSeed = []entity.User{
		{
			Name:     "admin",
			Email:    "admin@example.com",
			Password: "123123123",
			Role:     "admin",
		},
		{
			Name:     "user",
			Email:    "user@example.com",
			Password: "123123123",
			Role:     "user",
		},
	}

	for _, user := range userSeed {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil

}
