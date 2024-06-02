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

	req := dto.HistoriesGetRequest{
		Page:  page,
		Limit: limit,
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
