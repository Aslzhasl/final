package routes

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/user/signup", controllers.SignUp())
	incommingRoutes.POST("/user/login", controllers.Login())
	incommingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incommingRoutes.GET("/users/productview", controllers.SerchProduct())
	incommingRoutes.GET("/users/search", controllers.SerchProductByQuery())
}
