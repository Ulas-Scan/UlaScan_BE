package controller

import (
	"net/http"
	"strconv"

	"ulascan-be/dto"
	"ulascan-be/service"
	"ulascan-be/utils"

	"github.com/gin-gonic/gin"
)

type (
	HistoryController interface {
		GetHistories(ctx *gin.Context)
		GetHistory(ctx *gin.Context)
	}

	historyController struct {
		historyService service.HistoryService
	}
)

func NewHistoryController(hs service.HistoryService) HistoryController {
	return &historyController{
		historyService: hs,
	}
}

// GetHistories godoc
// @Summary Retrieve the user's analysis histories.
// @Description Retrieve the user's analysis histories.
// @Tags History
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Maximum number of results per page"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Security BearerAuth
// @Router /api/history [get]
func (c *historyController) GetHistories(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_HISTORIES, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_HISTORIES, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productName := ctx.Query("product_name")

	req := dto.HistoriesGetRequest{
		Page:        page,
		Limit:       limit,
		ProductName: productName,
	}

	result, err := c.historyService.GetHistories(ctx.Request.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_HISTORIES, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_HISTORIES, result)
	ctx.JSON(http.StatusOK, res)
}

// GetHistory godoc
// @Summary Retrieve the user's analysis history by id.
// @Description Retrieve the user's analysis history by id.
// @Tags History
// @Accept json
// @Produce json
// @Param id path string true "History ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Security BearerAuth
// @Router /api/history/{id} [get]
func (c *historyController) GetHistory(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)
	id := ctx.Param("id")

	result, err := c.historyService.GetHistoryById(ctx.Request.Context(), id, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_HISTORIES, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_HISTORIES, result)
	ctx.JSON(http.StatusOK, res)
}
