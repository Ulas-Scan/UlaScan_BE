package routes

import (
	"ulascan-be/controller"
	"ulascan-be/middleware"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func ML(route *gin.RouterGroup, mlController controller.MLController, jwtService service.JWTService) {
	routes := route.Group("/ml")
	{
		routes.GET("/guest/analysis", mlController.GetSentimentAnalysisAndSummarizationAsGuest)
		routes.GET("/analysis", middleware.Authenticate(jwtService), mlController.GetSentimentAnalysisAndSummarization)
	}
}
