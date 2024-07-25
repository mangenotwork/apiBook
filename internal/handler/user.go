package handler

import (
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/conf"
	"github.com/mangenotwork/common/log"
	"github.com/mangenotwork/common/utils"
	"net/http"
)

func SetAdmin(ctx *gin.Context) {
	num := dao.NewUserDao().GetUserNum()
	if num == 0 {

		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		password2 := ctx.PostForm("password2")

		if password != password2 {
			ctx.HTML(200, "err.html", gin.H{
				"Title": conf.Conf.Default.App.Name,
				"err":   "两次密码不一致",
			})
			return
		}

		user := &entity.User{
			UserId:   utils.IDMd5(),
			Name:     name,
			Password: password,
			IsAdmin:  1,
		}

		err := dao.NewUserDao().Create(user)
		if err != nil {
			log.Error(err)
			ctx.HTML(200, "err.html", gin.H{
				"Title": conf.Conf.Default.App.Name,
				"err":   "创建用户失败",
			})
			return
		}

		setToken(ctx, user.UserId)
		return
	}

	ctx.HTML(200, "err.html", gin.H{
		"Title": conf.Conf.Default.App.Name,
		"err":   "管理员已经存在",
	})
	return
}

func setToken(ctx *gin.Context, userId string) {
	j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
	j.AddClaims("userId", userId)

	token, tokenErr := j.Token()
	if tokenErr != nil {
		log.Error("生产token错误， err = ", tokenErr)
	}

	ctx.SetCookie(define.UserToken, token, define.TokenExpires, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/home")

	return
}

func Login(ctx *gin.Context) {

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	user, err := dao.NewUserDao().Get(name)
	if err != nil {
		log.Error(err)
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   "登录遇到错误",
		})
		return
	}

	if name == user.Name && password == user.Password {
		setToken(ctx, user.UserId)
		return
	}

	ctx.HTML(200, "err.html", gin.H{
		"Title": conf.Conf.Default.App.Name,
		"err":   "账号或密码错误",
	})
	return
}

func Out(ctx *gin.Context) {
	ctx.SetCookie("sign", "", 60*60*24*7, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}
