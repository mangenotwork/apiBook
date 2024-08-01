package handler

import (
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
	"time"

	"apiBook/common/conf"
	"github.com/gin-gonic/gin"
)

func ginH(h gin.H) gin.H {
	h["Title"] = conf.Conf.Default.App.Name
	h["TimeStamp"] = time.Now().Unix() // define.TimeStamp
	return h
}

func NotFond(ctx *gin.Context) {
	// 实现内部重定向
	ctx.HTML(
		http.StatusOK,
		"not_fond.html",
		ginH(gin.H{}),
	)
}

func Index(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		ginH(gin.H{
			"nav": "index",
		}),
	)
}

func LoginPage(ctx *gin.Context) {

	token, _ := ctx.Cookie(define.UserToken)
	if token != "" {
		j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
		if err := j.ParseToken(token); err == nil {
			ctx.Redirect(http.StatusFound, "/home")
			return
		}
	}

	// 检查是否存在用户表
	num := dao.NewUserDao().GetUserNum()
	if num == 0 {
		ctx.HTML(
			http.StatusOK,
			"init.html",
			gin.H{
				"nav":  "init",
				"csrf": FormSetCSRF(ctx.Request),
			},
		)
		return
	}

	ctx.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"nav":  "login",
			"csrf": FormSetCSRF(ctx.Request),
		},
	)
	return
}

func Home(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"nav": "home",
		},
	)
}

func FormSetCSRF(r *http.Request) template.HTML {
	fragment := fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`,
		define.CsrfName, csrf.Token(r))
	return template.HTML(fragment)
}
