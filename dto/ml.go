package dto

type MLResult struct {
	CountNegative    int     `json:"count_negative"`
	CountPositive    int     `json:"count_positive"`
	Packaging        float32 `json:"packaging"`
	Delivery         float32 `json:"delivery"`
	AdminResponse    float32 `json:"admin_response"`
	ProductCondition float32 `json:"product_condition"`
	Summary          string  `json:"summary"`
}
