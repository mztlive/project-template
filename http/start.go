package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mztlive/project-template/pkg/cors"
	"github.com/spf13/cast"
)

func Start(port int32) {
	engine := gin.Default()
	engine.Use(cors.GinCorsHandler())
	installRouter(engine)

	// 启动服务
	run(engine, port)
}

func run(engin *gin.Engine, port int32) {
	srv := &http.Server{
		Addr:         "127.0.0.1:" + cast.ToString(port),
		Handler:      engin,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 启动服务
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
