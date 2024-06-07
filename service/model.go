package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"ulascan-be/dto"
)

type (
	ModelService interface {
		Predict(ctx context.Context, req dto.PredictRequest) (dto.PredictResponse, error)
	}

	modelService struct {
		predictEndpoint string
	}
)

func NewModelService() ModelService {
	return &modelService{
		predictEndpoint: "http://" + os.Getenv("ML_HOST") + "/predict",
	}
}

func (s *modelService) Predict(ctx context.Context, req dto.PredictRequest) (dto.PredictResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return dto.PredictResponse{}, dto.ErrMarshallJson
	}

	fmt.Println("=========== JSON ================")
	fmt.Println(string(jsonData))

	// Create the HTTP request
	client := &http.Client{}
	httpReq, err := http.NewRequestWithContext(ctx, "GET", s.predictEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return dto.PredictResponse{}, dto.ErrCreateHttpRequest
	}
	httpReq.Header.Add("Content-Type", "application/json")

	// Perform the HTTP request
	res, err := client.Do(httpReq)
	if err != nil {
		return dto.PredictResponse{}, dto.ErrSendsHttpRequest
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return dto.PredictResponse{}, dto.ErrReadHttpResponseBody
	}

	// Check if the HTTP status code is 200 OK
	if res.StatusCode != http.StatusOK {
		return dto.PredictResponse{}, dto.ErrNotOk
	}

	// Parse the response JSON into the response DTO
	var response dto.PredictResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dto.PredictResponse{}, dto.ErrParseJson
	}

	return response, nil
}
