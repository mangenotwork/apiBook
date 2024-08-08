package handler

import (
	"apiBook/common/log"
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

func Err(ctx *gin.Context) {
	msg := ctx.Query("msg")
	ctx.HTML(
		http.StatusOK,
		"err.html",
		ginH(gin.H{
			"err": msg,
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

	homeProjectList := make([]*HomeProjectItem, 0)
	for _, v := range projectList {

		item := &HomeProjectItem{
			ProjectId:     v.ProjectId,
			Name:          v.Name,
			Description:   v.Description,
			CreateUserAcc: v.CreateUserAcc,
			CreateDate:    v.CreateDate,
			Private:       v.Private,
		}

		if v.CreateUserAcc == userAcc {
			item.IsOperation = 1
		}

		homeProjectList = append(homeProjectList, item)
	}

	ctx.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"nav":         "home",
			"isAdmin":     isAdmin, // 1是管理员
			"userName":    ctx.GetString("userName"),
			"projectList": homeProjectList,
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
			"userList": dao.NewUserDao().GetAllUser(),
			"userName": ctx.GetString("userName"),
			"isAdmin":  1, // 1是管理员
		},
	)
	return
}

func My(ctx *gin.Context) {
	userAcc := ctx.GetString("userAcc")

	userInfo, err := dao.NewUserDao().Get(userAcc)
	if err != nil || userInfo == nil {
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   "获取用户信息失败。",
		})
	}

	projectList := dao.NewProjectDao().GetList(userAcc)

	ctx.HTML(
		http.StatusOK,
		"my.html",
		gin.H{
			"userName":    ctx.GetString("userName"),
			"isAdmin":     userInfo.IsAdmin, // 1是管理员
			"userInfo":    userInfo,
			"projectList": projectList,
		},
	)
	return
}

func ProjectIndex(ctx *gin.Context) {
	userAcc := ctx.GetString("userAcc")

	userInfo, err := dao.NewUserDao().Get(userAcc)
	if err != nil || userInfo == nil {
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   "获取用户信息失败。",
		})
	}

	pid := ctx.Param("pid")

	project, err := dao.NewProjectDao().Get(pid, userAcc)
	if err != nil {
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   err.Error(),
		})
		return
	}

	userList, _ := dao.NewProjectDao().GetUserList(pid, userAcc)

	user, err := dao.NewUserDao().GetUsers(userList)
	if err != nil {
		ctx.HTML(200, "err.html", gin.H{
			"Title": conf.Conf.Default.App.Name,
			"err":   err.Error(),
		})
		return
	}

	projectUserList := make([]*ProjectUser, 0)

	for _, v := range user {
		log.Info(v, userAcc)
		item := &ProjectUser{
			Name:      v.Name,
			Account:   v.Account,
			IsDisable: v.IsDisable,
			Pid:       project.ProjectId,
		}

		if v.Account == userAcc {
			item.IsCreate = 1
		}

		projectUserList = append(projectUserList, item)
	}

	allUser := GetUserList()

	hasMap := make(map[string]struct{})
	for _, v := range projectUserList {
		hasMap[v.Account] = struct{}{}
	}

	canJsonUser := make([]*UserInfo, 0)

	for _, u := range allUser {
		if u.IsDisable == 1 {
			continue
		}
		if _, ok := hasMap[u.Account]; !ok {
			canJsonUser = append(canJsonUser, u)
		}
	}

	ctx.HTML(
		http.StatusOK,
		"project_configure.html",
		gin.H{
			"userName":        ctx.GetString("userName"),
			"isAdmin":         userInfo.IsAdmin, // 1是管理员
			"userInfo":        userInfo,
			"project":         project,
			"projectUserList": projectUserList,
			"canJsonUser":     canJsonUser,
		},
	)
	return

}
