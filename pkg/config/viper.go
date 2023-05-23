package config

import "github.com/spf13/viper"

// 初始化配置的路径
// 默认配置文件名为config.toml
func Initialize(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
}
