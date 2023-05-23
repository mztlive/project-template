package config

import (
	"log"

	"github.com/spf13/viper"
)

// 初始化配置的路径
// 默认配置文件名为local.toml
func Initialize(configPath string) {
	viper.SetConfigName("local")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatalf("Fatal error config file: %s \n", err.Error())
	}
}
