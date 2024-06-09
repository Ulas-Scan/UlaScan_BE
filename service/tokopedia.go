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
		GetProduct(ctx context.Context, req dto.GetProductRequest) (dto.GetProductResponse, error)
		GetReviews(ctx context.Context, req dto.GetReviewsRequest) ([]dto.ReviewResponse, error)
		GetShopAvatar(ctx context.Context, shopDomain string) (string, error)
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

func (s *tokopediaService) GetProduct(ctx context.Context, req dto.GetProductRequest) (dto.GetProductResponse, error) {
	payload := strings.NewReader(fmt.Sprintf(`{
		"operationName": "PDPGetLayoutQuery",
		"variables": {
			"shopDomain": "%s",
			"productKey": "%s",
			"apiVersion": 1
		},
		"query": "fragment ProductVariant on pdpDataProductVariant {\n  errorCode\n  parentID\n  defaultChild\n  sizeChart\n  totalStockFmt\n  variants {\n    productVariantID\n    variantID\n    name\n    identifier\n    option {\n      picture {\n        urlOriginal: url\n        urlThumbnail: url100\n        __typename\n      }\n      productVariantOptionID\n      variantUnitValueID\n      value\n      hex\n      stock\n      __typename\n    }\n    __typename\n  }\n  children {\n    productID\n    price\n    priceFmt\n    slashPriceFmt\n    discPercentage\n    optionID\n    optionName\n    productName\n    productURL\n    picture {\n      urlOriginal: url\n      urlThumbnail: url100\n      __typename\n    }\n    stock {\n      stock\n      isBuyable\n      stockWordingHTML\n      minimumOrder\n      maximumOrder\n      __typename\n    }\n    isCOD\n    isWishlist\n    campaignInfo {\n      campaignID\n      campaignType\n      campaignTypeName\n      campaignIdentifier\n      background\n      discountPercentage\n      originalPrice\n      discountPrice\n      stock\n      stockSoldPercentage\n      startDate\n      endDate\n      endDateUnix\n         isAppsOnly\n      isActive\n      hideGimmick\n      isCheckImei\n      minOrder\n      __typename\n    }\n    thematicCampaign {\n      additionalInfo\n      background\n      campaignName\n       __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ProductMedia on pdpDataProductMedia {\n  media {\n    type\n    urlOriginal: URLOriginal\n    urlThumbnail: URLThumbnail\n    urlMaxRes: URLMaxRes\n    videoUrl: videoURLAndroid\n    prefix\n    suffix\n    description\n    variantOptionID\n    __typename\n  }\n  videos {\n    source\n    url\n    __typename\n  }\n  __typename\n}\n\nfragment ProductCategoryCarousel on pdpDataCategoryCarousel {\n  linkText\n  titleCarousel\n   list {\n    categoryID\n     title\n        __typename\n  }\n  __typename\n}\n\nfragment ProductHighlight on pdpDataProductContent {\n  name\n }\n\nfragment ProductCustomInfo on pdpDataCustomInfo {\n  title\n   separator\n  description\n  __typename\n}\n\nfragment ProductInfo on pdpDataProductInfo {\n  row\n  content {\n    title\n    subtitle\n      __typename\n  }\n  __typename\n}\n\nfragment ProductDetail on pdpDataProductDetail {\n  content {\n    title\n    subtitle\n  }\n  __typename\n}\n\nfragment ProductDataInfo on pdpDataInfo {\n  title\n   __typename\n}\n\nfragment ProductSocial on pdpDataSocialProof {\n  row\n  content {\n    title\n    subtitle\n     type\n    rating\n    __typename\n  }\n  __typename\n}\n\nfragment ProductDetailMediaComponent on pdpDataProductDetailMediaComponent {\n  title\n  description\n  contentMedia {\n    url\n    ratio\n    type\n    __typename\n  }\n  show\n  ctaText\n  __typename\n}\n\nquery PDPGetLayoutQuery($shopDomain: String, $productKey: String, $layoutID: String, $apiVersion: Float, $userLocation: pdpUserLocation, $extParam: String, $tokonow: pdpTokoNow, $deviceID: String) {\n  pdpGetLayout(shopDomain: $shopDomain, productKey: $productKey, layoutID: $layoutID, apiVersion: $apiVersion, userLocation: $userLocation, extParam: $extParam, tokonow: $tokonow, deviceID: $deviceID) {\n           basicInfo {\n          id: productID\n        shopName\n    }\n    components {\n      name\n      data {\n        ...ProductMedia        ...ProductHighlight\n        ...ProductInfo\n        ...ProductDetail\n        ...ProductSocial\n        ...ProductDataInfo\n        ...ProductCustomInfo\n        ...ProductVariant\n        ...ProductCategoryCarousel\n        ...ProductDetailMediaComponent\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"
	}`, req.ShopDomain, req.ProductKey))

	client := &http.Client{}
	tokopediaReq, err := http.NewRequest("POST", s.url, payload)
	if err != nil {
		fmt.Println(err)
		return dto.GetProductResponse{}, dto.ErrCreateHttpRequest
	}

	tokopediaReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	tokopediaReq.Header.Add("X-Source", "tokopedia-lite")
	tokopediaReq.Header.Add("X-Tkpd-Lite-Service", "zeus")
	tokopediaReq.Header.Add("Referer", req.ProductUrl)
	tokopediaReq.Header.Add("X-TKPD-AKAMAI", "pdpGetLayout")
	tokopediaReq.Header.Add("Content-Type", "application/json")

	res, err := client.Do(tokopediaReq)
	if err != nil {
		return dto.GetProductResponse{}, dto.ErrSendsHttpRequest
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return dto.GetProductResponse{}, dto.ErrReadHttpResponseBody
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dto.GetProductResponse{}, dto.ErrParseJson
	}

	// Extracting necessary data
	productData := response["data"].(map[string]interface{})["pdpGetLayout"].(map[string]interface{})["components"].([]interface{})
	var productName string
	var description string
	for _, component := range productData {
		if component.(map[string]interface{})["name"].(string) == "product_detail" {
			content := component.(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["content"].([]interface{})
			for _, c := range content {
				if c.(map[string]interface{})["title"].(string) == "Deskripsi" {
					description = c.(map[string]interface{})["subtitle"].(string)
					break
				}
			}
		} else if component.(map[string]interface{})["name"].(string) == "product_content" {
			productName = component.(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["name"].(string)
		}
	}

	// Extracting images
	productMedia := response["data"].(map[string]interface{})["pdpGetLayout"].(map[string]interface{})["components"].([]interface{})[1].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["media"].([]interface{})
	var imageUrls []string
	for _, media := range productMedia {
		if media.(map[string]interface{})["type"].(string) == "image" {
			imageUrls = append(imageUrls, media.(map[string]interface{})["urlOriginal"].(string))
		}
	}

	productResponse := dto.GetProductResponse{
		ProductName:        productName,
		ProductDescription: description,
		ShopName:           response["data"].(map[string]interface{})["pdpGetLayout"].(map[string]interface{})["basicInfo"].(map[string]interface{})["shopName"].(string),
		ProductId:          response["data"].(map[string]interface{})["pdpGetLayout"].(map[string]interface{})["basicInfo"].(map[string]interface{})["id"].(string),
		ImageUrls:          imageUrls,
	}

	return productResponse, nil

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

func (s *tokopediaService) GetShopAvatar(ctx context.Context, shopDomain string) (string, error) {
	payload := strings.NewReader(fmt.Sprintf(` {
        "operationName": "ShopInfoCore",
        "variables": {
            "id": 0,
            "domain": "%s"
        },
        "query": "query ShopInfoCore($id: Int!, $domain: String) {\n  shopInfoByID(input: {shopIDs: [$id], fields: [\"active_product\", \"allow_manage_all\", \"assets\", \"core\", \"closed_info\", \"create_info\", \"favorite\", \"location\", \"status\", \"is_open\", \"other-goldos\", \"shipment\", \"shopstats\", \"shop-snippet\", \"other-shiploc\", \"shopHomeType\", \"branch-link\", \"goapotik\", \"fs_type\"], domain: $domain, source: \"shoppage\"}) {\n    result {\n                 shopAssets {\n        avatar\n          }\n                   }\n     }\n}\n"
    }`, shopDomain))

	client := &http.Client{}
	tokopediaReq, err := http.NewRequest("POST", s.url, payload)
	if err != nil {
		return "", dto.ErrCreateHttpRequest
	}

	tokopediaReq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	tokopediaReq.Header.Add("X-Source", "tokopedia-lite")
	tokopediaReq.Header.Add("X-Tkpd-Lite-Service", "zeus")
	tokopediaReq.Header.Add("Referer", "https://www.tokopedia.com/"+shopDomain)
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

	var response dto.ShopAvatarResponseTokopedia
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", dto.ErrParseJson
	}

	if len(response.Data.ShopInfoByID.Result) > 0 {
		shopAvatar := response.Data.ShopInfoByID.Result[0].ShopAssets.Avatar
		return shopAvatar, nil
	}

	return "", dto.ErrShopAvatarNotFound

}
