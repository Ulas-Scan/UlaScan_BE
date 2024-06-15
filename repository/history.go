package repository

import (
	"context"
	"strings"

	"ulascan-be/dto"
	"ulascan-be/entity"

	"gorm.io/gorm"
)

type (
	HistoryRepository interface {
		CreateHistory(ctx context.Context, tx *gorm.DB, history entity.History) (entity.History, error)
		GetHistories(ctx context.Context, tx *gorm.DB, dto dto.HistoriesGetRequest, userId string) ([]entity.History, int64, error)
		GetHistoryById(ctx context.Context, tx *gorm.DB, historyId string, userId string) (entity.History, error)
		CheckByProductId(ctx context.Context, tx *gorm.DB, productId string, userId string) bool
		DeleteByProductId(ctx context.Context, tx *gorm.DB, productId string, userId string) error
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

	scope := tx.WithContext(ctx)
	if dto.ProductName != "" {
		scope = scope.Where("LOWER(product_name) LIKE ?", "%"+strings.ToLower(dto.ProductName)+"%")
	}

	// Query the paginated records
	err = scope.
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

func (r *historyRepository) CheckByProductId(ctx context.Context, tx *gorm.DB, productId string, userId string) bool {
	if tx == nil {
		tx = r.db
	}

	var history entity.History
	err := tx.WithContext(ctx).
		Where("product_id = ?", productId).
		Where("user_id = ?", userId).
		Take(&history).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}

	return true
}

func (r *historyRepository) DeleteByProductId(ctx context.Context, tx *gorm.DB, productId string, userId string) error {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).Delete(&entity.History{}, "product_id = ? AND user_id = ?", productId, userId).Error
	if err != nil {
		return err
	}

	return nil
}
