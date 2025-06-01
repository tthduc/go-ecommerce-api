package initialize

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/global"
	"go-ecommerce-api/internal/routers"
)

func InitRouter() *gin.Engine {
	// gin.Engine is used to create a new Gin server instance
	var r *gin.Engine

	if global.Config.Server.Mode == "development" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Set up global middleware
	r.Use(gin.Recovery())
	// r.Use(middleware.Logger()) // Uncomment if you have a custom logger middleware
	// r.Use(middleware.Cors()) // Uncomment if you have a custom CORS middleware
	// r.Use(middleware.Auth()) // Uncomment if you have a custom authentication middleware

	// Initialize routes
	mangerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/check-status")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		mangerRouter.InitAdminRouter(mainGroup)
	}

	return r
}
