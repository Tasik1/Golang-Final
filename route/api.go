package route

import (
	"github.com/gin-gonic/gin"
	"go_final/handlers"
	"go_final/middleware"
	"net/http"
)

func RunAPI(address string) error {

	userHandler := handlers.NewUserHandler()
	productHandler := handlers.NewProductHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to SDU Canteen!")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userSecuredRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userSecuredRoutes.GET("/", userHandler.GetAllUsers)
		userSecuredRoutes.GET("/:id", userHandler.GetUser)
		adminRoutes := userSecuredRoutes.Group("/", middleware.CheckAdmin())
		{
			adminRoutes.PUT("/:id", userHandler.UpdateUser)
			adminRoutes.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)
		productRoutes.GET("/:productID", productHandler.GetProduct)
		adminRoutes := productRoutes.Group("/", middleware.CheckAdmin())
		{
			adminRoutes.POST("/", productHandler.CreateProduct)
			adminRoutes.PUT("/:productID", productHandler.UpdateProduct)
			adminRoutes.DELETE("/:productID", productHandler.DeleteProduct)
		}
	}

	return r.Run(address)
}
