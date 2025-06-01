package initialize

import (
	"fmt"
	"go-ecommerce-api/global"
	"go-ecommerce-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysql() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		global.Logger.Error("failed to connect to database", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("connect to database successfully")
	global.Mdb = db

	// Set the connection pool settings
	setPool()

	// Migrate the database tables
	migrateTables()
}

func setPool() {
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("failed to get sqlDB from gorm", zap.Error(err))
		panic(err)
	}

	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdle)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.Mysql.MaxLife))
}

func migrateTables() {
	if err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	); err != nil {
		global.Logger.Error("failed to migrate tables", zap.Error(err))
		panic(err)
	}
}
