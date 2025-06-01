package global

import (
	"github.com/redis/go-redis/v9"
	"go-ecommerce-api/pkg/logger"
	"go-ecommerce-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
