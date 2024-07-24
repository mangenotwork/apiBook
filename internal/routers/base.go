package routers

import (
	"apiBook/internal/handler"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/ginHelper"
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
	Svg()
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
	page.GET("/", handler.Login)
	page.GET("/home", handler.Home)
}

func Test() {
	test := Router.Group("")
	test.GET("/test", ginHelper.Handle(handler.Testcase))
}
