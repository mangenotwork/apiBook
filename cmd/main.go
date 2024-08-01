package main

import (
	"apiBook/common/conf"
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/routers"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	conf.InitConf("./conf/")
	db.Init()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	routers.Router = gin.New()

	// 初始化全局变量
	define.CsrfAuthKey, _ = conf.YamlGetString("csrfAuthKey")
	define.CsrfName, _ = conf.YamlGetString("csrfName")

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Conf.Default.HttpServer.Prod),
		Handler:        routers.Routers(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.InfoF("http服务启动 0.0.0.0:%s", conf.Conf.Default.HttpServer.Prod)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.ErrorF("http服务出现异常:%s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("关闭服务 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("关闭服务 Err:", err)
	}
	log.Info("服务已关闭")
}
