package service

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"ulascan-be/constants"
	"ulascan-be/dto"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type (
	GeminiService interface {
		Analyze(ctx context.Context, analyzeReq string) (dto.AnalyzeResponse, error)
		Summarize(ctx context.Context, summarizeReq string) (string, error)
		CloseClient() error
	}

	geminiService struct {
		client *genai.Client
		model  *genai.GenerativeModel
	}
)

func NewGeminiService() GeminiService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-pro-latest")
	model.ResponseMIMEType = "application/json"

	return &geminiService{
		client: client,
		model:  model,
	}
}

func (s *geminiService) Analyze(ctx context.Context, analyzeReq string) (dto.AnalyzeResponse, error) {
	prompt := constants.PROMPT_ANALYZE + "\n" + analyzeReq

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return dto.AnalyzeResponse{}, err
	}
	return parseAnalyzeResponse(resp)
}

func (s *geminiService) Summarize(ctx context.Context, summarizeReq string) (string, error) {
	prompt := constants.PROMPT_SUMMARIZE + "\n" + summarizeReq

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	return parseSummaryResponse(resp)
}

func (s *geminiService) CloseClient() error {
	return s.client.Close()
}

func parseAnalyzeResponse(resp *genai.GenerateContentResponse) (dto.AnalyzeResponse, error) {
	var analyzeResp dto.AnalyzeResponse
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					if err := json.Unmarshal([]byte(txt), &analyzeResp); err != nil {
						return dto.AnalyzeResponse{}, err
					}
				}
			}
		}
	}
	return analyzeResp, nil
}

func parseSummaryResponse(resp *genai.GenerateContentResponse) (string, error) {
	var summaryResp struct {
		Summary string `json:"summary"`
	}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					if err := json.Unmarshal([]byte(txt), &summaryResp); err != nil {
						return "", err
					}
				}
			}
		}
	}
	return summaryResp.Summary, nil
}
