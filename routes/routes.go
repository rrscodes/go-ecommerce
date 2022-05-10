package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rrscodes/go-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
}