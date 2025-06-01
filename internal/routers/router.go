package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// gin.Engine is used to create a new Gin server instance
	r := gin.Default()

	return r
}
