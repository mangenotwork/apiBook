package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"apiBook/common/conf"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func ginH(h gin.H) gin.H {
	h["Title"] = conf.Conf.Default.App.Name
	h["TimeStamp"] = time.Now().Unix() // define.TimeStamp
	return h
}

func NotFond(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"not_fond.html",
		ginH(gin.H{}),
	)
	return
}

func Index(ctx *gin.Context) {
	userAcc := ctx.GetString("userAcc")

	isAdmin := 0
	if dao.NewUserDao().IsAdmin(userAcc) {
		isAdmin = 1
	}

	ctx.HTML(
		http.StatusOK,
		"index.html",
		ginH(gin.H{
			"nav":      "index",
			"isAdmin":  isAdmin, // 1是管理员
			"userName": ctx.GetString("userName"),
		}),
	)
	return
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
	userAcc := ctx.GetString("userAcc")

	isAdmin := 0
	if dao.NewUserDao().IsAdmin(userAcc) {
		isAdmin = 1
	}

	projectList := dao.NewProjectDao().GetList(userAcc)

	ctx.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"nav":         "home",
			"isAdmin":     isAdmin, // 1是管理员
			"userName":    ctx.GetString("userName"),
			"projectList": projectList,
		},
	)

	return
}

func FormSetCSRF(r *http.Request) template.HTML {
	fragment := fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`,
		define.CsrfName, csrf.Token(r))
	return template.HTML(fragment)
}

func UserMange(ctx *gin.Context) {
	userAcc := ctx.GetString("userAcc")

	if !dao.NewUserDao().IsAdmin(userAcc) {
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   "对不起，你不是管理员无权限操作。",
		})
		return
	}

	ctx.HTML(
		http.StatusOK,
		"user_manage.html",
		gin.H{
			"nav":      "home",
			"userList": dao.NewUserDao().GetAllUser(),
			"userName": ctx.GetString("userName"),
			"isAdmin":  1, // 1是管理员
		},
	)
	return
}
