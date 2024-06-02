package service

import (
	"context"
	"math"

	"ulascan-be/dto"
	"ulascan-be/entity"
	"ulascan-be/repository"
)

type (
	HistoryService interface {
		CreateHistory(ctx context.Context, req dto.HistoryCreateRequest) (dto.HistoryResponse, error)
		GetHistories(ctx context.Context, req dto.HistoriesGetRequest, userId string) (dto.HistoriesResponse, error)
		GetHistoryById(ctx context.Context, historyId string, userId string) (dto.HistoryResponse, error)
	}

	historyService struct {
		historyRepo repository.HistoryRepository
	}
)

func NewHistoryService(historyRepo repository.HistoryRepository) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
	}
}

func (s *historyService) CreateHistory(ctx context.Context, req dto.HistoryCreateRequest) (dto.HistoryResponse, error) {
	history := entity.History{
		URL:         req.ProductID,
		ProductID:   req.ProductID,
		ProductName: req.ProductName,
		Content:     req.Content,
		UserID:      req.UserID,
	}

	historyCreated, err := s.historyRepo.CreateHistory(ctx, nil, history)
	if err != nil {
		return dto.HistoryResponse{}, dto.ErrCreateHistory
	}

	return dto.HistoryResponse{
		UserID:      historyCreated.UserID,
		URL:         historyCreated.URL,
		ProductID:   historyCreated.ProductID,
		ProductName: historyCreated.ProductName,
		Content:     historyCreated.Content,
	}, nil
}

func (s *historyService) GetHistories(ctx context.Context, req dto.HistoriesGetRequest, userId string) (dto.HistoriesResponse, error) {
	histories, total, err := s.historyRepo.GetHistories(ctx, nil, req, userId)
	if err != nil {
		return dto.HistoriesResponse{}, dto.ErrGetHistories
	}

	pages := int(math.Ceil(float64(total) / float64(req.Limit)))

	return dto.HistoriesResponse{
		Histories: histories,
		Page:      req.Page,
		Limit:     req.Limit,
		Total:     total,
		Pages:     pages,
	}, nil
}

func (s *historyService) GetHistoryById(ctx context.Context, historyId string, userId string) (dto.HistoryResponse, error) {
	history, err := s.historyRepo.GetHistoryById(ctx, nil, historyId, userId)
	if err != nil {
		return dto.HistoryResponse{}, dto.ErrGetHistory
	}

	return dto.HistoryResponse{
		ID:          history.ID,
		URL:         history.URL,
		ProductID:   history.ProductID,
		ProductName: history.ProductName,
		Content:     history.Content,
		UserID:      history.UserID,
	}, nil
}
