package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"ulascan-be/dto"
)

type (
	TokopediaService interface {
		GetProductId(ctx context.Context, req dto.GetProductIdRequest) (string, error)
		GetReviews(ctx context.Context, req dto.GetReviewsRequest) ([]dto.ReviewResponse, error)
	}

	tokopediaService struct {
		url string
	}
)

func NewTokopediaService() TokopediaService {
	return &tokopediaService{
		url: "https://gql.tokopedia.com/graphql/",
	}
}

func (s *tokopediaService) GetProductId(ctx context.Context, req dto.GetProductIdRequest) (string, error) {
	payload := strings.NewReader(fmt.Sprintf(`{
		"operationName": "PDPGetLayoutQuery",
		"variables": {
			"shopDomain": "%s",
			"productKey": "%s",
			"apiVersion": 1
		},
		"query": "query PDPGetLayoutQuery($shopDomain: String, $productKey: String, $apiVersion: Float) {\n  pdpGetLayout(shopDomain: $shopDomain, productKey: $productKey, apiVersion: $apiVersion) {\n    basicInfo {\n      id: productID\n    }\n  }\n}\n"
	}`, req.ShopDomain, req.ProductKey))

	client := &http.Client{}
	tokopediaReq, err := http.NewRequest("POST", s.url, payload)
	if err != nil {
		fmt.Println(err)
		return "", dto.ErrCreateHttpRequest
	}

	tokopediaReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	tokopediaReq.Header.Add("X-Source", "tokopedia-lite")
	tokopediaReq.Header.Add("X-Tkpd-Lite-Service", "zeus")
	tokopediaReq.Header.Add("Referer", req.ProductUrl)
	tokopediaReq.Header.Add("X-TKPD-AKAMAI", "pdpGetLayout")
	tokopediaReq.Header.Add("Content-Type", "application/json")

	res, err := client.Do(tokopediaReq)
	if err != nil {
		return "", dto.ErrSendsHttpRequest
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", dto.ErrReadHttpResponseBody
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", dto.ErrParseJson
	}

	id, ok := response["data"].(map[string]interface{})["pdpGetLayout"].(map[string]interface{})["basicInfo"].(map[string]interface{})["id"].(string)
	if !ok {
		return "", dto.ErrProductId
	}

	return id, nil
}

func (s *tokopediaService) GetReviews(ctx context.Context, req dto.GetReviewsRequest) ([]dto.ReviewResponse, error) {
	var allReviews []dto.ReviewResponse

	for page := 1; page <= 2; page++ {
		// Prepare the request payload
		payload := fmt.Sprintf(`{
			"operationName": "productReviewList",
			"variables": {
				"productID": "%s",
				"page": %d,
				"limit": 50,
				"sortBy": "create_time desc"
			},
			"query": "query productReviewList($productID: String!, $page: Int!, $limit: Int!, $sortBy: String) {\n  productrevGetProductReviewList(productID: $productID, page: $page, limit: $limit, sortBy: $sortBy) {\n    list {\n      message\n      productRating\n    }\n  }\n}\n"
		}`, req.ProductId, page)

		client := &http.Client{}

		tokopediaReq, err := http.NewRequest("POST", s.url, strings.NewReader(payload))
		if err != nil {
			return nil, dto.ErrCreateHttpRequest
		}

		tokopediaReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
		tokopediaReq.Header.Add("X-Source", "tokopedia-lite")
		tokopediaReq.Header.Add("X-Tkpd-Lite-Service", "zeus")
		tokopediaReq.Header.Add("Referer", req.ProductUrl)
		tokopediaReq.Header.Add("Content-Type", "application/json")

		res, err := client.Do(tokopediaReq)
		if err != nil {
			return nil, dto.ErrSendsHttpRequest
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, dto.ErrReadHttpResponseBody
		}

		var response dto.ProductReviewResponseTokopedia
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, dto.ErrParseJson
		}

		reviews := response.Data.ProductrevGetProductReviewList.List
		for _, review := range reviews {
			allReviews = append(allReviews, dto.ReviewResponse{
				Message: review.Message,
				Rating:  review.ProductRating,
			})
		}

		if len(reviews) < 50 {
			break
		}
	}

	return allReviews, nil
}
