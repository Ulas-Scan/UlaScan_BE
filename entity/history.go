package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type History struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	ProductID   string    `gorm:"primaryKey" json:"product_id"`
	URL         string    `json:"url" gorm:"not null"`
	ProductName string    `json:"product_name" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`

	Timestamp
}

func (h *History) BeforeCreate(tx *gorm.DB) (err error) {
	if h.UserID == uuid.Nil {
		return gorm.ErrEmptySlice
	}
	return nil
}
