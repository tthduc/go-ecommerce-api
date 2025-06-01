package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public routes
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.GET("/register")
		userRouterPublic.POST("/otp")
	}

	// private routes
	userRouterPrivate := router.Group("/user")
	//userRouterPrivate.Use(AuthMiddleware())
	{
		userRouterPrivate.GET("/profile")
	}
}
