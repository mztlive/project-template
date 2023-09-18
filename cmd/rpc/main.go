package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mztlive/go-pkgs/config"
	"github.com/mztlive/go-pkgs/database"
	"github.com/mztlive/go-pkgs/logger"
	"github.com/mztlive/go-pkgs/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logPath = flag.String("log", "./log", "the log path")

	configPath = flag.String("c", "./local.toml", "the config file")

	zapcore *zap.Logger
)

func main() {

	defer zapcore.Sync()

	// 初始化配置文件
	flag.Parse()
	config.Initialize(*configPath)

	// 初始化日志
	zapcore = logger.Initialize(*logPath, viper.GetBool("app.debug"))

	// 初始化数据库信息
	database.Initialize(
		viper.GetString("database.dsn"),
		viper.GetString("database.driver"),
		viper.GetInt("database.max_open_conns"),
		viper.GetInt("database.max_idle_conns"),
	)

	snowflake.Initialize(viper.GetInt64("snowflake.worker_id"))

	go startRpcServer()

	log.Println("start rpc server ...")
	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill)
	defer stop()
	<-quit.Done()
	log.Println("Shutdown Server ...")
}

func startRpcServer() {
	lis, err := net.Listen("tcp", ":"+viper.GetString("rpc.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 初始化rpc服务
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
