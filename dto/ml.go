package dto

type MLResult struct {
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	Rating             int      `json:"rating"`
	Ulasan             int      `json:"ulasan"`
	Bintang            float64  `json:"bintang"`
	ImageUrls          []string `json:"image_urls"`
	ShopName           string   `json:"shop_name"`
	CountNegative      int      `json:"count_negative"`
	CountPositive      int      `json:"count_positive"`
	Packaging          float32  `json:"packaging"`
	Delivery           float32  `json:"delivery"`
	AdminResponse      float32  `json:"admin_response"`
	ProductCondition   float32  `json:"product_condition"`
	Summary            string   `json:"summary"`
}
