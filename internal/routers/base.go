package routers

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/handler"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func Routers() *gin.Engine {
	Router.Use(gzip.Gzip(gzip.DefaultCompression), Base())
	Router.StaticFS("/js", http.Dir("./assets/js"))
	Router.StaticFS("/css", http.Dir("./assets/css"))
	Router.StaticFS("/images", http.Dir("./assets/images"))

	Router.Delims("{{", "}}")
	FuncMap()
	Router.LoadHTMLGlob("assets/html/**/*")

	Login()    // 登录
	Page()     // 页面
	Project()  // 项目
	Document() // 文档
	User()     // 用户相关
	Admin()    // 管理员

	Mock() // mock

	return Router
}

// Base  拦截器，限流，记录， 商户，认证等等
func Base() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()
		// 设置请求端ip
		setIP(ctx)
		// 全局多语言语种
		ctx.Set(ginHelper.Lang, ctx.GetHeader(ginHelper.Lang))
		// 全局来源
		ctx.Set(ginHelper.Source, ctx.GetHeader(ginHelper.Source))

		ctx.Next()

		// 记录请求
		reqLog(ctx, startTime)

	}
}

func setIP(ctx *gin.Context) {
	ctx.Set(ginHelper.ReqIP, ginHelper.GetIP(ctx.Request))
}

func reqLog(ctx *gin.Context, startTime time.Time) {
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)
	reqMethod := ctx.Request.Method
	reqUri := ctx.Request.RequestURI
	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()

	// 大于300ms的接口需要记录
	if latencyTime.Milliseconds() > 300 {
		log.WarnF(" %3d | %13v | %15s | %s | %s | >>> 慢接口，请优化!!!",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri)
	} else {
		log.HttpInfoF(" %10v | %10s | %s:%d | %s ",
			latencyTime,
			clientIP,
			reqMethod,
			statusCode,
			reqUri)
	}
}

func CSRFMiddleware() gin.HandlerFunc {
	csrfMiddleware := csrf.Protect(
		[]byte(define.CsrfAuthKey),
		csrf.Secure(false),
		csrf.HttpOnly(true),
		csrf.CookieName(define.CsrfName),
		csrf.FieldName(define.CsrfName),
		csrf.RequestHeader(define.CsrfName),
		csrf.ErrorHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusForbidden)
			_, _ = writer.Write([]byte(`403`))
		})),
	)
	return adapter.Wrap(csrfMiddleware)
}

func FuncMap() {
	Router.SetFuncMap(template.FuncMap{
		"SVG":         SvgHtml,
		"InputModule": Input,
		"ApiBookInfo": ApiBookInfo,
	})
}

func Login() {
	login := Router.Group("")
	login.Use(CSRFMiddleware())
	login.GET("/", handler.LoginPage)
	login.POST("/set/admin", handler.SetAdmin)
	login.POST("/login", handler.Login)
	login.GET("/out", handler.Out)
	login.GET("/clear/project", handler.ClearData)
}

func Page() {
	Router.NoRoute(handler.NotFond)
	Router.NoMethod(handler.NotFond)
	page := Router.Group("")
	page.Use(AuthPG())
	page.GET("/index", handler.Index)
	page.GET("/home", handler.Home)
}

func Project() {
	project := Router.Group("/project")
	project.Use(AuthAPI())
	project.GET("/list", handler.ProjectList)        // 项目列表
	project.GET("/item", handler.ProjectItem)        // 项目详情
	project.POST("/create", handler.ProjectCreate)   // 创建项目
	project.POST("/modify", handler.ProjectModify)   // 修改项目
	project.POST("/delete", handler.ProjectDelete)   // 删除项目
	project.GET("/users", handler.ProjectUsers)      // 项目协作人员列表
	project.POST("/adduser", handler.ProjectAddUser) // 项目添加协助人员
	project.POST("/deluser", handler.ProjectDelUser) // 项目移除协作者
}

func Document() {
	document := Router.Group("/document")
	document.Use(AuthAPI())
	document.GET("/dir/list", handler.DocumentDirList)      // 文档目录列表
	document.POST("/dir/create", handler.DocumentDirCreate) // 创建文档目录
	document.POST("/dir/delete", handler.DocumentDirDelete) // 删除文档目录
	document.POST("/dir/modify", handler.DocumentDirModify) // 修改文档目录
	document.POST("/dir/sort", handler.DocumentDirSort)     // 排序文档目录
	document.POST("/list", handler.DocumentList)            // 文档列表
	document.POST("/create", handler.DocumentCreate)        // 创建文档
	document.POST("/item", handler.DocumentItem)            // 文档详情
	document.POST("/modify", handler.DocumentModify)        // 修改文档
	document.POST("/delete", handler.DocumentDelete)        // 删除文档
	document.POST("/sort", handler.DocumentSort)            // 排序文档
}

func User() {
	user := Router.Group("/user")
	user.Use(AuthAPI())
	user.GET("/info", handler.GetUserInfo)                   // 获取用户信息
	user.POST("/modify", handler.UserModify)                 // 修改用户信息
	user.POST("/reset/password ", handler.UserResetPassword) // 重置用户密码
	user.GET("/list", handler.UserList)                      // 获取所有用户列表
}

func Admin() {
	admin := Router.Group("/mange")
	admin.Use(AuthAPI())
	admin.GET("/t", handler.AdminT)
	admin.POST("/create/user", handler.AdminCreateUser)   // 创建用户
	admin.POST("/delete/user", handler.AdminDeleteUser)   // 删除用户
	admin.POST("/disable/user", handler.AdminDisableUser) // 禁用用户
}

func Mock() {
	Router.Any("/simulator/:path", handler.Simulator) // mock 模拟器
	mock := Router.Group("/mock")
	mock.Use(AuthAPI())
	// 新增mock
}
