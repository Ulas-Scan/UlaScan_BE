package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type History struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	URL         string    `json:"url" gorm:"not null"`
	ProductID   string    `json:"product_id" gorm:"not null" `
	ProductName string    `json:"product_name" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	UserID      uuid.UUID `json:"user_id" gorm:"not null" `
	User        User      `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`

	Timestamp
}

func (h *History) BeforeCreate(tx *gorm.DB) (err error) {
	if h.UserID == uuid.Nil {
		return gorm.ErrEmptySlice
	}
	return nil
}
