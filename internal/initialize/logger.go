package initialize

import (
	"go-ecommerce-api/global"
	"go-ecommerce-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
