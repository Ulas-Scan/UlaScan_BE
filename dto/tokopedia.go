package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_PARSE_URL       = "failed parse url"
	MESSAGE_FAILED_SPLIT_URL       = "failed split url"
	MESSAGE_FAILED_GET_PRODUCT_ID  = "failed get product id"
	MESSAGE_FAILED_GET_REVIEWS     = "failed get product reviews"
	MESSAGE_FAILED_GET_SHOP_AVATAR = "failed get shop avatar"

	// Success
	MESSAGE_SUCCESS_GET_REVIEWS = "success get reviews"
)

var (
	ErrProductUrlMissing     = errors.New("product url is required")
	ErrProductUrlWrongFormat = errors.New("invalid product url format")
	ErrNotTokopediaUrls      = errors.New("invalid domain, only tokopedia.com urls are accepted")
	ErrProductId             = errors.New("failed to extract product id")
	ErrShopAvatarNotFound    = errors.New("shop avatar not found")
)

type ProductReviewResponseTokopedia struct {
	Data struct {
		ProductrevGetProductReviewList struct {
			List []struct {
				Message       string `json:"message"`
				ProductRating int    `json:"productRating"`
			} `json:"list"`
		} `json:"productrevGetProductReviewList"`
	} `json:"data"`
}

type ShopAvatarResponseTokopedia struct {
	Data struct {
		ShopInfoByID struct {
			Result []struct {
				ShopAssets struct {
					Avatar string `json:"avatar"`
				} `json:"shopAssets"`
			} `json:"result"`
		} `json:"shopInfoByID"`
	} `json:"data"`
}

type GetProductRequest struct {
	ProductUrl string
	ProductKey string
	ShopDomain string
}

type GetProductResponse struct {
	ProductName        string
	ProductDescription string
	ShopName           string
	ProductId          string
	ImageUrls          []string
}

type GetReviewsRequest struct {
	ProductUrl string
	ProductId  string
}

type ReviewResponse struct {
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}
