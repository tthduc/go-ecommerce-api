package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/controller"
)

func NewRouter() *gin.Engine {
	// gin.Engine is used to create a new Gin server instance
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/user/register", controller.NewUserController().Register)
	}

	return r
}
