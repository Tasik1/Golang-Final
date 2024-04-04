package route

import (
	"github.com/gin-gonic/gin"
	"go_final/handlers"
	"go_final/middleware"
	"net/http"
)

func RunAPI(address string) error {

	userHandler := handlers.NewUserHandler()

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
		userSecuredRoutes.PUT("/:id", userHandler.UpdateUser)
		userSecuredRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	return r.Run(address)
}
