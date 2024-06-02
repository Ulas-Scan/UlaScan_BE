package routes

import (
	"ulascan-be/controller"
	"ulascan-be/middleware"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func History(route *gin.Engine, historyController controller.HistoryController, jwtService service.JWTService) {
	routes := route.Group("/api/history")
	{
		// User
		routes.GET("", middleware.Authenticate(jwtService), historyController.GetHistories)
		routes.GET("/:id", middleware.Authenticate(jwtService), historyController.GetHistory)
	}
}
