//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"go-ecommerce-api/internal/controller"
	"go-ecommerce-api/internal/repo"
	"go-ecommerce-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
