package routers

import (
	"apiBook/internal/handler"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/ginHelper"
	"net/http"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func Routers() *gin.Engine {
	Router.Use(gzip.Gzip(gzip.DefaultCompression))
	Router.StaticFS("/js", http.Dir("./assets/js"))
	Router.StaticFS("/css", http.Dir("./assets/css"))
	Router.StaticFS("/images", http.Dir("./assets/images"))
	Router.Delims("{{", "}}")
	//Svg()
	Router.LoadHTMLGlob("assets/html/**/*")
	//Login()
	Page()
	//API()
	//Data()
	Test()
	return Router
}

func Page() {
	page := Router.Group("")
	page.GET("/index", handler.Index)
}

func Test() {
	test := Router.Group("")
	test.GET("/test", ginHelper.Handle(handler.Testcase))
}
