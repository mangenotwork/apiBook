package routers

import (
	"apiBook/internal/define"
	"apiBook/internal/handler"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
	"html/template"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func Routers() *gin.Engine {
	Router.Use(gzip.Gzip(gzip.DefaultCompression))
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
	return Router
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
	//project.Use(AuthAPI())
	project.GET("/list", handler.ProjectList)        // 项目列表
	project.GET("/item", handler.ProjectItem)        // 项目详情
	project.POST("/create", handler.ProjectCreate)   // 创建项目
	project.POST("/modify", handler.ProjectModify)   // 修改项目
	project.POST("/delete", handler.ProjectDelete)   // 删除项目
	project.POST("/users", handler.ProjectUsers)     // 项目协作人员列表
	project.POST("/adduser", handler.ProjectAddUser) // 项目添加协助人员
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
	document.POST("/modify", handler.DocumentModify)        // 修改文档
	document.POST("/delete", handler.DocumentDelete)        // 删除文档
	document.POST("/sort", handler.DocumentSort)            // 排序文档
}

func User() {
	user := Router.Group("/user")
	user.Use(AuthAPI())
	user.GET("/info", handler.UserInfo)                      // 获取用户信息
	user.POST("/modify", handler.UserModify)                 // 修改用户信息
	user.POST("/reset/password ", handler.UserResetPassword) // 重置用户密码
	user.GET("/list", handler.UserList)                      // 获取所有用户列表
}

func Admin() {
	admin := Router.Group("/mange")
	admin.Use(AuthAPI())
	// todo 是否是管理员中间件
	admin.POST("/create/user", handler.AdminCreateUser)   // 创建用户
	admin.POST("/delete/user", handler.AdminDeleteUser)   // 删除用户
	admin.POST("/disable/user", handler.AdminDisableUser) // 禁用用户
}
