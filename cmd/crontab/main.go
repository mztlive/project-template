package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/mztlive/go-pkgs/config"
	"github.com/mztlive/go-pkgs/database"
	"github.com/mztlive/go-pkgs/logger"
	"github.com/mztlive/go-pkgs/snowflake"
	"github.com/robfig/cron/v3"
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
	zap.ReplaceGlobals(zapcore)

	// 初始化数据库连接
	database.Initialize(
		viper.GetString("database.dsn"),
		viper.GetString("database.driver"),
		viper.GetInt("database.max_open_conns"),
		viper.GetInt("database.max_idle_conns"),
	)

	snowflake.Initialize(viper.GetInt64("snowflake.worker_id"))

	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill)
	defer stop()
	<-quit.Done()
	log.Println("Shutdown Crontab Server ...")
}

// start crontab
func StartCrontab() {
	crontab := cron.New(cron.WithSeconds())

	// 这里添加定时任务

	zap.L().Info("所有定时任务添加完成")
	crontab.Start()
}

// 包装任务以捕获和记录panic
func WrapJobWithRecovery(job func()) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error(
					"crontab server panic",
					zap.Any("err", err),
					zap.String("stack", string(debug.Stack())),
				)
			}
		}()
		job()
	}
}
