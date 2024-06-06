package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_PARSE_URL      = "failed parse url"
	MESSAGE_FAILED_SPLIT_URL      = "failed split url"
	MESSAGE_FAILED_GET_PRODUCT_ID = "failed get product id"
	MESSAGE_FAILED_GET_REVIEWS    = "failed get product reviews"

	// Success
	MESSAGE_SUCCESS_GET_REVIEWS = "success get reviews"
)

var (
	ErrProductUrlMissing         = errors.New("product url is required")
	ErrProductUrlWrongFormat     = errors.New("invalid product url format")
	ErrCreateTokopediaRequest    = errors.New("failed to create http request")
	ErrSendsTokopediaRequest     = errors.New("failed to sends http request")
	ErrReadTokopediaResponseBody = errors.New("failed to read http response body")
	ErrParseJson                 = errors.New("failed to parse response json")
	ErrProductId                 = errors.New("failed to extract product id")
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

type GetProductIdRequest struct {
	ProductUrl string
	ProductKey string
	ShopDomain string
}

type GetReviewsRequest struct {
	ProductUrl string
	ProductId  string
}

type ReviewResponse struct {
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}
