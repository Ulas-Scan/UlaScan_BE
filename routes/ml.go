package routes

import (
	"ulascan-be/controller"
	"ulascan-be/middleware"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func ML(route *gin.Engine, mlController controller.MLController, jwtService service.JWTService) {
	routes := route.Group("/api/ml")
	{
		routes.GET("/guest/analysis", mlController.GetSentimentAnalysisAndSummarization)
		routes.GET("/analysis", middleware.Authenticate(jwtService), mlController.GetSentimentAnalysisAndSummarization)
	}
}
