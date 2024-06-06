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
	TokopediaController interface {
		GetReviews(ctx *gin.Context)
	}

	tokopediaController struct {
		tokopediaService service.TokopediaService
	}
)

func NewTokopediaController(ts service.TokopediaService) TokopediaController {
	return &tokopediaController{
		tokopediaService: ts,
	}
}

func (c *tokopediaController) GetReviews(ctx *gin.Context) {
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

	productReq := dto.GetProductIdRequest{
		ShopDomain: pathParts[1],
		ProductKey: pathParts[2],
		ProductUrl: "https://www.tokopedia.com/" + pathParts[1] + "/" + pathParts[2],
	}

	productId, err := c.tokopediaService.GetProductId(ctx, productReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewsReq := dto.GetReviewsRequest{
		ProductUrl: productReq.ProductUrl,
		ProductId:  productId,
	}

	result, err := c.tokopediaService.GetReviews(ctx, reviewsReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEWS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEWS, result)
	ctx.JSON(http.StatusOK, res)
}
