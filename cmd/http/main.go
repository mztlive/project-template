package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cztech.com/market-center/http"
	"cztech.com/market-center/pkg/config"
	"cztech.com/market-center/pkg/database"
	"cztech.com/market-center/pkg/logger"
	"cztech.com/market-center/pkg/snowflake"
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

	go http.Start(viper.GetInt32("http.port"))

	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill)
	defer stop()
	<-quit.Done()
	log.Println("Shutdown Server ...")
}
