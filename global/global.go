package global

import (
	"go-ecommerce-api/pkg/logger"
	"go-ecommerce-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
