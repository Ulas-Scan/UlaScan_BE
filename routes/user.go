package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javakanaya/ulascan-be/controller"
	"github.com/javakanaya/ulascan-be/middleware"
	"github.com/javakanaya/ulascan-be/service"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		// User
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
		routes.PATCH("", middleware.Authenticate(jwtService), userController.Update)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
	}
}
