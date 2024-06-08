package service

import (
	"context"

	"ulascan-be/dto"
)

type (
	GeminiService interface {
		Analyze(ctx context.Context, analyzeReq string) (dto.AnalyzeResponse, error)
		Summarize(ctx context.Context, summarizeReq string) (string, error)
	}

	geminiService struct {
	}
)

func NewGeminiService() GeminiService {
	return &geminiService{}
}

func (s *geminiService) Analyze(ctx context.Context, analyzeReq string) (dto.AnalyzeResponse, error) {

	return dto.AnalyzeResponse{}, nil
}

func (s *geminiService) Summarize(ctx context.Context, summarizeReq string) (string, error) {

	return "", nil
}
