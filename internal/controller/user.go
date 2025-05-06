package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/service"
	"go-ecommerce-api/response"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) Register(c *gin.Context) {
	response.SuccessResponse(c, 20001, uc.UserService.Register())
}
