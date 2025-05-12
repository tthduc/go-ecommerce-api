package initialize

import (
	"github.com/spf13/viper"
	"go-ecommerce-api/global"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
