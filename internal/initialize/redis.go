package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-ecommerce-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: "",
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis connect failed", zap.Error(err))
		panic("failed to connect to Redis: " + err.Error())
	}

	global.Logger.Info("redis connect successfully")
	global.Rdb = rdb
}
