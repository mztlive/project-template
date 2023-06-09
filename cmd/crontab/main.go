package main

import (
	"flag"

	"github.com/mztlive/project-template/pkg/config"
	"github.com/mztlive/project-template/pkg/database"
	"github.com/mztlive/project-template/pkg/logger"
	"github.com/mztlive/project-template/pkg/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logPath = flag.String("log", "./log", "the log path")

	configPath = flag.String("c", "./configs/local.toml", "the config file")

	zapcore *zap.Logger
)

func main() {

	defer zapcore.Sync()

	// 初始化配置文件
	flag.Parse()
	config.Initialize(*configPath)

	// 初始化日志
	zapcore = logger.Initialize(*logPath, viper.GetBool("app.debug"))

	// 初始化数据库连接
	database.Initialize(
		viper.GetString("database.dsn"),
		viper.GetString("database.driver"),
		viper.GetInt("database.max_open_conns"),
		viper.GetInt("database.max_idle_conns"),
	)

	snowflake.Initialize(viper.GetInt64("snowflake.worker_id"))
}
