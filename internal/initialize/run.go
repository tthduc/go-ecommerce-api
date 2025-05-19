package initialize

import (
	"go-ecommerce-api/global"
	"go.uber.org/zap"
)

func Run() {
	// Initialize configuration
	LoadConfig()

	// Initialize the logger
	InitLogger()
	global.Logger.Info("config log ok!!!", zap.String("ok", "success"))

	// Initialize MySQL connection
	InitMysql()

	// Initialize Redis connection
	InitRedis()

	// Initialize the router
	r := InitRouter()

	r.Run(":8080")
}
