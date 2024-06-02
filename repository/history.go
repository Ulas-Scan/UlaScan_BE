package repository

import (
	"context"

	"ulascan-be/dto"
	"ulascan-be/entity"

	"gorm.io/gorm"
)

type (
	HistoryRepository interface {
		CreateHistory(ctx context.Context, tx *gorm.DB, history entity.History) (entity.History, error)
		GetHistories(ctx context.Context, tx *gorm.DB, dto dto.HistoriesGetRequest, userId string) ([]entity.History, int64, error)
		GetHistoryById(ctx context.Context, tx *gorm.DB, historyId string, userId string) (entity.History, error)
	}

	historyRepository struct {
		db *gorm.DB
	}
)

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{
		db: db,
	}
}

func (r *historyRepository) CreateHistory(ctx context.Context, tx *gorm.DB, history entity.History) (entity.History, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&history).Error; err != nil {
		return entity.History{}, err
	}

	return history, nil
}

func (r *historyRepository) GetHistories(ctx context.Context, tx *gorm.DB, dto dto.HistoriesGetRequest, userId string) ([]entity.History, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var histories []entity.History
	var totalCount int64

	limit := dto.Limit
	page := dto.Page
	offset := (page - 1) * limit

	// Count the total number of records
	err := tx.WithContext(ctx).
		Model(&entity.History{}).
		Where("user_id = ?", userId).
		Count(&totalCount).Error
	if err != nil {
		return []entity.History{}, 0, err
	}

	// Query the paginated records
	err = tx.WithContext(ctx).
		Where("user_id = ?", userId).
		Limit(limit).Offset(offset).
		Find(&histories).Error
	if err != nil {
		return []entity.History{}, 0, err
	}

	return histories, totalCount, nil
}

func (r *historyRepository) GetHistoryById(ctx context.Context, tx *gorm.DB, historyId string, userId string) (entity.History, error) {
	if tx == nil {
		tx = r.db
	}

	var history entity.History
	err := tx.WithContext(ctx).
		Where("id = ?", historyId).
		Where("user_id = ?", userId).
		Take(&history).Error
	if err != nil {
		return entity.History{}, err
	}

	return history, nil
}
