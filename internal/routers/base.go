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
	Login()
	Page()
	//API()
	//Data()
	return Router
}

func Login() {
	login := Router.Group("")
	login.Use(ginHelper.CSRFMiddleware())
	login.GET("/", handler.LoginPage)
	login.POST("/set/admin", handler.SetAdmin)
	login.POST("/login", handler.Login)
	login.GET("/out", handler.Out)
}

func Page() {
	Router.NoRoute(handler.NotFond)
	Router.NoMethod(handler.NotFond)
	page := Router.Group("")
	page.Use(AuthPG())
	page.GET("/index", handler.Index)
	page.GET("/home", handler.Home)
}
