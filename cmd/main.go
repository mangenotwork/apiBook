package main

import (
	"apiBook/common/conf"
	"apiBook/common/db"
	"apiBook/common/fenci"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/routers"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 读取配置文件与初始化本地数据文件
	var confPath string
	flag.StringVar(&confPath, "c", "./conf/app.yaml", "配置文件路径")
	flag.Parse()
	conf.InitConf(confPath)
	db.Init()

	// 初始化分词
	fenci.Seg.SkipLog = true
	err := fenci.Seg.LoadDict()
	if err != nil {
		log.Panic("初始化分词失败")
	}

	// 全局变量初始化
	define.CsrfAuthKey, _ = conf.YamlGetString("csrfAuthKey")
	define.CsrfName, _ = conf.YamlGetString("csrfName")

	// 初始化gin框架
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	//gin.DefaultWriter = io.Discard
	routers.Router = gin.New()

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
