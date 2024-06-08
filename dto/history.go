package dto

import (
	"errors"
	"ulascan-be/entity"

	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_HISTORY = "failed create history"
	MESSAGE_FAILED_GET_HISTORIES  = "failed get histories"
	MESSAGE_FAILED_GET_HISTORY    = "failed get history"

	// Success
	MESSAGE_SUCCESS_CREATE_HISTORY = "success create history"
	MESSAGE_SUCCESS_GET_HISTORIES  = "success get histories"
	MESSAGE_SUCCESS_GET_HISTORY    = "success get history"
)

var (
	ErrCreateHistory = errors.New("failed to create history")
	ErrDeleteHistory = errors.New("failed to delete history")
	ErrGetHistories  = errors.New("failed to get histories")
	ErrGetHistory    = errors.New("failed to get history")
)

type (
	HistoriesGetRequest struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}

	HistoriesResponse struct {
		Histories []entity.History `json:"histories"`
		Page      int              `json:"page"`
		Pages     int              `json:"pages"`
		Limit     int              `json:"limit"`
		Total     int64            `json:"total"`
	}

	HistoryCreateRequest struct {
		UserID           uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
		ProductID        string    `json:"product_id" form:"product_id" binding:"required"`
		URL              string    `json:"url" form:"url" binding:"required"`
		ProductName      string    `json:"product_name" form:"product_name" binding:"required"`
		PositiveCount    int       `json:"positive_count" form:"positive_count" binding:"required"`
		NegativeCount    int       `json:"negative_count" form:"negative_count" binding:"required"`
		Packaging        float32   `json:"packaging"  form:"packaging" binding:"required"`
		Delivery         float32   `json:"delivery" form:"delivery" binding:"required"`
		AdminResponse    float32   `json:"admin_response" form:"admin_response" binding:"required"`
		ProductCondition float32   `json:"product_condition" form:"product_condition" binding:"required"`
		Content          string    `json:"content" form:"content"`
	}

	HistoryResponse struct {
		ID               uuid.UUID `json:"id"`
		UserID           uuid.UUID `json:"user_id"`
		ProductID        string    `json:"product_id"`
		URL              string    `json:"url"`
		ProductName      string    `json:"product_name"`
		PositiveCount    int       `json:"positive_count" `
		NegativeCount    int       `json:"negative_count" `
		Packaging        float32   `json:"packaging"`
		Delivery         float32   `json:"delivery"`
		AdminResponse    float32   `json:"admin_response"`
		ProductCondition float32   `json:"product_condition"`
		Content          string    `json:"content"`
	}
)
