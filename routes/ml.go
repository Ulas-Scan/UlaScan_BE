package routes

import (
	"ulascan-be/controller"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func ML(route *gin.Engine, mlController controller.MLController, jwtService service.JWTService) {
	routes := route.Group("/api/ml")
	{
		routes.GET("/analysis", mlController.GetSentimentAnalysisAndSummarization)
	}
}
