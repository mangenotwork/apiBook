package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func AdminT(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	if !dao.NewUserDao().IsAdmin(userAcc) {
		ctx.APIOutPutError(nil, "不是管理员")
		return
	}

	ctx.APIOutPut("ok", "ok")
	return
}

func AdminCreateUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	param := &AdminCreateUserReq{}
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

	if !dao.NewUserDao().IsAdmin(userAcc) {
		ctx.APIOutPutError(nil, "不是管理员")
		return
	}

	if param.Password != param.Password2 {
		ctx.APIOutPutError(nil, "两次密码不一致")
		return
	}

	if dao.NewUserDao().HasUserAccount(param.Account) {
		ctx.APIOutPutError(nil, "用户账号存在")
		return
	}

	has, err := dao.NewUserDao().HasUserName(param.Name)
	if err != nil {
		log.Info(err)
		ctx.APIOutPutError(nil, "创建用户失败")
		return
	}

	if has {
		ctx.APIOutPutError(nil, "用户名称存在")
		return
	}

	user := &entity.User{
		Account:    param.Account,
		Name:       param.Name,
		Password:   param.Password,
		IsAdmin:    param.IsAdmin,
		CreateTime: time.Now().Unix(),
	}

	err = dao.NewUserDao().Create(user)
	if err != nil {
		ctx.APIOutPutError(nil, "创建用户失败")
		return
	}

	ctx.APIOutPut("创建用户成功", "创建用户成功")
	return
}

func AdminDeleteUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	param := &AdminDeleteUserReq{}
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

	if !dao.NewUserDao().IsAdmin(userAcc) {
		ctx.APIOutPutError(nil, "不是管理员")
		return
	}

	err = dao.NewUserDao().DelUser(param.Account)
	if err != nil {
		ctx.APIOutPutError(nil, "删除用户失败")
		return
	}

	ctx.APIOutPut("删除用户成功", "删除用户成功")
	return
}

func AdminDisableUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	param := &AdminDisableUserReq{}
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

	if !dao.NewUserDao().IsAdmin(userAcc) {
		ctx.APIOutPutError(nil, "不是管理员")
		return
	}

	err = dao.NewUserDao().DisableUser(param.Account, param.IsDisable)
	if err != nil {
		ctx.APIOutPutError(nil, "操作失败")
		return
	}

	ctx.APIOutPut("操作成功", "操作成功")
	return
}
