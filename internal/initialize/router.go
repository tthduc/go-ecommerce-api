package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/controller"
	"go-ecommerce-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before aa")
		c.Next()
		fmt.Println("after aa")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before bb")
		c.Next()
		fmt.Println("after bb")
	}
}

func CC(c *gin.Context) {
	fmt.Println("before cc")
	c.Next()
	fmt.Println("after cc")
}

func InitRouter() *gin.Engine {
	// gin.Engine is used to create a new Gin server instance
	r := gin.Default()

	// Middleware
	r.Use(middlewares.Auth(), BB(), CC)

	v1 := r.Group("/v1")
	{
		v1.GET("/user/register", controller.NewUserController().Register)
	}

	return r
}
