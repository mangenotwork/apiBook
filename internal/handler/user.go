package handler

import (
	"apiBook/common/conf"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"time"
)

func SetAdmin(ctx *gin.Context) {
	num := dao.NewUserDao().GetUserNum()
	if num == 0 {
		account := ctx.PostForm("account")
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		password2 := ctx.PostForm("password2")

		if password != password2 {
			ctx.HTML(200, "err.html", gin.H{
				"Title":     conf.Conf.Default.App.Name,
				"err":       "两次密码不一致",
				"returnUrl": "/",
			})
			return
		}

		user := &entity.User{
			Account:    account,
			Name:       name,
			Password:   password,
			IsAdmin:    1,
			CreateTime: time.Now().Unix(),
		}

		err := dao.NewUserDao().Create(user)
		if err != nil {
			log.Error(err)
			ctx.HTML(200, "err.html", gin.H{
				"Title":     conf.Conf.Default.App.Name,
				"err":       "创建用户失败",
				"returnUrl": "/",
			})
			return
		}

		setToken(ctx, user.Account)
		return
	}

	ctx.HTML(200, "err.html", gin.H{
		"Title":     conf.Conf.Default.App.Name,
		"err":       "管理员已经存在",
		"returnUrl": "/",
	})

	return
}

func setToken(ctx *gin.Context, userAcc string) {
	j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
	j.AddClaims("userAcc", userAcc)

	token, tokenErr := j.Token()
	if tokenErr != nil {
		log.Error("生产token错误， err = ", tokenErr)
	}

	ctx.SetCookie(define.UserToken, token, define.TokenExpires, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/home")

	return
}

func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	if account == "" || password == "" {
		ctx.HTML(200, "err.html", gin.H{
			"Title":     conf.Conf.Default.App.Name,
			"err":       "账号或密码为空",
			"returnUrl": "/",
		})
	}

	user, err := dao.NewUserDao().Get(account)
	if err != nil {
		log.Error(err)
		ctx.HTML(200, "err.html", gin.H{
			"Title":     conf.Conf.Default.App.Name,
			"err":       "用户不存在",
			"returnUrl": "/",
		})
		return
	}

	if user.IsDisable == 1 {
		ctx.HTML(200, "err.html", gin.H{
			"Title":     conf.Conf.Default.App.Name,
			"err":       "用户被禁用",
			"returnUrl": "/",
		})
		return
	}

	if account == user.Account && password == user.Password {
		log.SendOperationLog(user.Account, "登入了系统")
		setToken(ctx, user.Account)
		return
	}

	ctx.HTML(200, "err.html", gin.H{
		"Title":     conf.Conf.Default.App.Name,
		"err":       "账号或密码错误",
		"returnUrl": "/",
	})
	return
}

func Out(ctx *gin.Context) {

	token, _ := ctx.Cookie(define.UserToken)
	j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
	if err := j.ParseToken(token); err == nil {
		log.SendOperationLog(j.GetString("userAcc"), "登出了系统")
	}

	ctx.SetCookie("sign", "", 60*60*24*7, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

func GetUserInfo(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	acc := ctx.Query("acc")

	data, err := dao.NewUserDao().Get(acc)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("未获取到用户信息"), "未获取到用户信息")
		return
	}

	resp := &UserInfo{
		Name:       data.Name,
		Account:    data.Account,
		CreateTime: data.CreateTime,
		IsDisable:  data.IsDisable,
	}

	ctx.APIOutPut(resp, "")
	return
}

func UserModify(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &UserModifyReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	if dao.NewUserDao().IsAdmin(userAcc) && len(param.Account) > 0 && userAcc != param.Account {
		err = dao.NewUserDao().Modify(param.Account, param.Name)
	} else {
		err = dao.NewUserDao().Modify(userAcc, param.Name)
	}

	if err != nil {
		ctx.APIOutPutError(err, "修改用户信息失败")
		return
	}

	log.SendOperationLog(userAcc, fmt.Sprintf("修改用户信息成功: %s", utils.AnyToJsonNotErr(param)))

	ctx.APIOutPut("修改用户信息成功", "修改用户信息成功")
	return
}

func UserResetPassword(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &UserResetPasswordReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	if param.Password != param.Password2 {
		ctx.APIOutPutError(fmt.Errorf("两次密码不一致"), "两次密码不一致")
		return
	}

	if dao.NewUserDao().IsAdmin(userAcc) && len(param.Account) > 0 && userAcc != param.Account {
		err = dao.NewUserDao().ResetPassword(param.Account, param.Password)
	} else {
		err = dao.NewUserDao().ResetPassword(userAcc, param.Password)
	}

	if err != nil {
		ctx.APIOutPutError(err, "修改密码失败")
		return
	}

	log.SendOperationLog(userAcc, fmt.Sprintf("修改密码成功"))

	ctx.APIOutPut("修改密码成功", "修改密码成功")
	return
}

func UserList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	ctx.APIOutPut(GetUserList(), "")
	return
}

func GetUserList() []*UserInfo {
	list := dao.NewUserDao().GetAllUser()
	resp := make([]*UserInfo, 0)

	for _, v := range list {
		resp = append(resp, &UserInfo{
			Name:       v.Name,
			Account:    v.Account,
			CreateTime: v.CreateTime,
			IsDisable:  v.IsDisable,
		})
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].CreateTime > resp[j].CreateTime {
			return false
		}
		return true
	})

	return resp
}
