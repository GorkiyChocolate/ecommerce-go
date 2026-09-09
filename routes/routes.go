package routes

import (
	"ecommerce-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.POST("/users/signup", controllers.Signup())
	routes.POST("/users/login", controllers.Login())
	routes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	routes.GET("/users/productview", controllers.ProductView())
	routes.GET("/users/search", controllers.SearchProductByQuery())
}
