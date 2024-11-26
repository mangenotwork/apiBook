package main

import (
	"apiBook/common/conf"
	"apiBook/common/db"
	"apiBook/common/fenci"
	"apiBook/common/log"
	"apiBook/common/utils"
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
	"runtime/debug"
	"syscall"
	"time"
)

func main() {

	log.InitSysLog()

	defer func() {
		if err := recover(); err != nil {
			stack := debug.Stack()
			log.SendErrorLog(utils.AnyToString(err), string(stack))
			log.Error(err)
		}
	}()

	// 读取配置文件与初始化本地数据文件
	var confPath string
	flag.StringVar(&confPath, "c", "./conf/app.yaml", "配置文件路径")
	flag.Parse()
	log.SendSysLog("读取配置文件: " + confPath)
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
	// todo 需要配置指定什么模式
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.Discard

	gin.SetMode(gin.DebugMode)

	routers.Router = gin.New()

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Conf.Default.HttpServer.Prod),
		Handler:        routers.Routers(),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.InfoF("http服务启动 0.0.0.0:%s", conf.Conf.Default.HttpServer.Prod)
		log.SendSysLog(fmt.Sprintf("http服务启动 0.0.0.0:%s", conf.Conf.Default.HttpServer.Prod))
		if srvErr := srv.ListenAndServe(); srvErr != nil && errors.Is(srvErr, http.ErrServerClosed) {
			log.Error(fmt.Sprintf("http服务出现异常:%v", srvErr))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.SendSysLog("关闭服务 ...")
	log.Info("关闭服务 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if srvErr := srv.Shutdown(ctx); srvErr != nil {
		log.Panic(fmt.Sprintf("关闭服务 Err::%v", srvErr))
	}

	log.SendSysLog("服务已关闭")
	log.Info("服务已关闭")

}
