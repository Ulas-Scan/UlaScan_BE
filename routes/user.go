package routes

import (
	"ulascan-be/controller"
	"ulascan-be/middleware"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func User(route *gin.RouterGroup, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/user")
	{
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
	}
}
