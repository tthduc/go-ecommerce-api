package routers

import (
	"go-ecommerce-api/internal/routers/manger"
	"go-ecommerce-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manger.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
