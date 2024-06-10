package dto

const (
	// Failed
	MESSAGE_FAILED_ANALYZE = "failed analyze"

// Success
// MESSAGE_SUCCESS_PREDICT = "success predict"
)

var (
// ErrMarshallJson = errors.New("failed to marshall request body json")
)

type AnalyzeResponse struct {
	Packaging        float32 `json:"packaging"`
	Delivery         float32 `json:"delivery"`
	AdminResponse    float32 `json:"admin_response"`
	ProductCondition float32 `json:"product_condition"`
}
