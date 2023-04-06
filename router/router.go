package router

import (
	"jwt/controllers"
	"jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())

		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", middlewares.ProductAuthorization(), controllers.GetAllProducts)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
