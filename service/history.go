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
	isExist := s.historyRepo.CheckByProductId(ctx, nil, req.ProductID, req.UserID.String())
	if isExist {
		err := s.historyRepo.DeleteByProductId(ctx, nil, req.ProductID, req.UserID.String())
		if err != nil {
			return dto.HistoryResponse{}, dto.ErrDeleteHistory
		}
	}

	history := entity.History{
		URL:              req.URL,
		ProductID:        req.ProductID,
		ProductName:      req.ProductName,
		Rating:           req.Rating,
		Ulasan:           req.Ulasan,
		Bintang:          req.Bintang,
		CountPositive:    req.CountPositive,
		CountNegative:    req.CountNegative,
		Packaging:        req.Packaging,
		Delivery:         req.Delivery,
		AdminResponse:    req.AdminResponse,
		ProductCondition: req.ProductCondition,
		Content:          req.Content,
		UserID:           req.UserID,
	}

	historyCreated, err := s.historyRepo.CreateHistory(ctx, nil, history)
	if err != nil {
		return dto.HistoryResponse{}, dto.ErrCreateHistory
	}

	return dto.HistoryResponse{
		UserID:           historyCreated.UserID,
		URL:              historyCreated.URL,
		ProductID:        historyCreated.ProductID,
		Rating:           historyCreated.Rating,
		Ulasan:           historyCreated.Ulasan,
		Bintang:          historyCreated.Bintang,
		ProductName:      historyCreated.ProductName,
		CountPositive:    historyCreated.CountPositive,
		CountNegative:    historyCreated.CountNegative,
		Packaging:        historyCreated.Packaging,
		Delivery:         historyCreated.Delivery,
		AdminResponse:    historyCreated.AdminResponse,
		ProductCondition: historyCreated.ProductCondition,
		Content:          historyCreated.Content,
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
		UserID:           history.UserID,
		ProductID:        history.ProductID,
		ID:               history.ID,
		URL:              history.URL,
		Rating:           history.Rating,
		Ulasan:           history.Ulasan,
		Bintang:          history.Bintang,
		ProductName:      history.ProductName,
		CountPositive:    history.CountPositive,
		CountNegative:    history.CountNegative,
		Packaging:        history.Packaging,
		Delivery:         history.Delivery,
		AdminResponse:    history.AdminResponse,
		ProductCondition: history.ProductCondition,
		Content:          history.Content,
	}, nil
}
