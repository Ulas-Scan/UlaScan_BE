package controller

import (
	"net/http"

	"ulascan-be/dto"
	"ulascan-be/service"
	"ulascan-be/utils"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserCreateRequest true "User details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/user [post]
func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// Me godoc
// @Summary User info
// @Description Get user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Security BearerAuth
// @Router /api/user/me [get]
func (c *userController) Me(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	result, err := c.userService.GetUserById(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login user
// @Description Login user with the provided creds
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserLoginRequest true "User creds"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/user/login [post]
func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.userService.Verify(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}
