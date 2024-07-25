package main

import (
	"apiBook/common/db"
	"apiBook/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/conf"
	"log"
	"net/http"
	"time"
)

func main() {
	conf.InitConf("./conf/")
	db.Init()
	// 启动 https servers
	gin.SetMode(gin.DebugMode)
	server := &http.Server{
		Addr:           ":" + conf.Conf.Default.HttpServer.Prod,
		Handler:        routers.Routers(),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
