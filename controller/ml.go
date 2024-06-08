package controller

import (
	"net/http"
	"net/url"
	"strings"

	"ulascan-be/dto"
	"ulascan-be/service"
	"ulascan-be/utils"

	"github.com/gin-gonic/gin"
)

type (
	MLController interface {
		GetSentimentAnalysisAndSummarization(ctx *gin.Context)
	}

	mlController struct {
		tokopediaService service.TokopediaService
		modelService     service.ModelService
		geminiService    service.GeminiService
	}
)

func NewMLController(
	ts service.TokopediaService,
	ms service.ModelService,
	gs service.GeminiService,
) MLController {
	return &mlController{
		tokopediaService: ts,
		modelService:     ms,
		geminiService:    gs,
	}
}

func (c *mlController) GetSentimentAnalysisAndSummarization(ctx *gin.Context) {
	productUrl := ctx.Query("product_url")
	if productUrl == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEWS, dto.ErrProductUrlMissing.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	parsedUrl, err := url.Parse(productUrl)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEWS, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	pathParts := strings.Split(parsedUrl.Path, "/")
	if len(pathParts) < 3 {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEWS, dto.ErrProductUrlWrongFormat.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	productReq := dto.GetProductRequest{
		ShopDomain: pathParts[1],
		ProductKey: pathParts[2],
		ProductUrl: "https://www.tokopedia.com/" + pathParts[1] + "/" + pathParts[2],
	}

	product, err := c.tokopediaService.GetProduct(ctx, productReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// fmt.Println("=== PRODUCT ID ===")
	// fmt.Println(product)

	reviewsReq := dto.GetReviewsRequest{
		ProductUrl: productReq.ProductUrl,
		ProductId:  product.ProductId,
	}

	reviews, err := c.tokopediaService.GetReviews(ctx, reviewsReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEWS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// fmt.Println("=== REVIEWS ===")
	// fmt.Println(reviews)

	statements := make([]string, len(reviews))
	for i, review := range reviews {
		statements[i] = review.Message
	}

	predictReq := dto.PredictRequest{
		Statements: statements,
	}

	// fmt.Println("=== PREDICT REQ ===")
	// fmt.Println(predictReq)

	predictResult, err := c.modelService.Predict(ctx, predictReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PREDICT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var builder strings.Builder
	for _, review := range reviews {
		builder.WriteString(review.Message)
		builder.WriteString("\n")
	}
	concatenatedMessage := builder.String()

	analyzeResult, err := c.geminiService.Analyze(ctx, concatenatedMessage)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_ANALYZE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	summarizeResult, err := c.geminiService.Summarize(ctx, concatenatedMessage)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_ANALYZE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEWS, dto.MLResult{
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		ImageUrls:          product.ImageUrls,
		ShopName:           product.ShopName,
		CountNegative:      predictResult.CountNegative,
		CountPositive:      predictResult.CountPositive,
		Packaging:          analyzeResult.Packaging,
		Delivery:           analyzeResult.Delivery,
		AdminResponse:      analyzeResult.AdminResponse,
		ProductCondition:   analyzeResult.ProductCondition,
		Summary:            summarizeResult,
	})
	ctx.JSON(http.StatusOK, res)
}