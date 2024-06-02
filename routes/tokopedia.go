package routes

import (
	"ulascan-be/controller"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
)

func Tokopedia(route *gin.Engine, tokopediaController controller.TokopediaController, jwtService service.JWTService) {
	routes := route.Group("/api/tokopedia")
	{
		routes.GET("/reviews", tokopediaController.GetReviews)
	}
}
