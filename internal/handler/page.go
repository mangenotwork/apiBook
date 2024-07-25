package handler

import (
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"github.com/mangenotwork/common/ginHelper"
	"github.com/mangenotwork/common/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/conf"
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
				"csrf": ginHelper.FormSetCSRF(ctx.Request),
			},
		)
		return
	}

	ctx.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"nav":  "login",
			"csrf": ginHelper.FormSetCSRF(ctx.Request),
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
