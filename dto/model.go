package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_PREDICT = "failed predict"

	// Success
	MESSAGE_SUCCESS_PREDICT = "success predict"
)

var (
	ErrMarshallJson             = errors.New("failed to marshall request body json")
	ErrModelInternalServerError = errors.New("internal server error from ml erver")
)

type PredictRequest struct {
	Statements []string `json:"statements"`
}

type PredictResponse struct {
	CountNegative int `json:"Negative"`
	CountPositive int `json:"Positive"`
}
