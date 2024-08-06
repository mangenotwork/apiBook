package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ProjectList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	userAcc := ctx.GetString("userAcc")
	log.Info(userAcc)

	resp := dao.NewProjectDao().GetList(userAcc)
	ctx.APIOutPut(resp, "")
}

func ProjectItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")
	log.Info("pid = ", pid)
	userAcc := ctx.GetString("userAcc")
	log.Info(userAcc)

	resp, err := dao.NewProjectDao().Get(pid, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	// todo 私有化项目获取协作者信息

	ctx.APIOutPut(resp, "")
}

func ProjectCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.Project{}
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

	param.CreateUserAcc = userAcc
	param.ProjectId = utils.IDMd5()
	param.CreateDate = utils.NowDate()
	err = dao.NewProjectDao().Create(param, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, "创建失败")
		return
	}

	// 创建默认和回收站目录

	dirDef := &entity.DocumentDir{
		DirId:   utils.IDStr(),
		DirName: "默认",
		Sort:    1,
	}

	dirRecycleBin := &entity.DocumentDir{
		DirId:   utils.IDStr(),
		DirName: "回收站",
		Sort:    2,
	}

	err = dao.NewDirDao().Create(param.ProjectId, dirDef)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDirDao().Create(param.ProjectId, dirRecycleBin)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("创建成功", "创建成功")
	return
}

func ProjectModify(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.Project{}
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

	err = dao.NewProjectDao().Modify(param, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, "修改失败")
		return
	}

	ctx.APIOutPut("修改成功", "修改成功")
	return
}

func ProjectDelete(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.Project{}
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

	err = dao.NewProjectDao().Delete(param.ProjectId, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, "删除失败")
		return
	}

	ctx.APIOutPut("删除成功", "删除成功")
	return
}

func ProjectUsers(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")
	log.Info("pid = ", pid)

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	userList, err := dao.NewProjectDao().GetUserList(pid, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	user, err := dao.NewUserDao().GetUsers(userList)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	resp := &ProjectUsersResp{
		List: make([]*UserInfo, 0),
	}

	for _, v := range user {
		resp.List = append(resp.List, &UserInfo{
			Name:    v.Name,
			Account: v.Account,
		})
	}

	ctx.APIOutPut(resp, "")
	return
}

func ProjectAddUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &ProjectAddUserReq{}
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

	err = dao.NewProjectDao().AddUser(param.PId, userAcc, param.Account)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("添加协作者成功", "添加协作者成功")
	return
}

func ProjectDelUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &ProjectAddUserReq{}
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

	err = dao.NewProjectDao().DelUser(param.PId, userAcc, param.Account)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("删除协作者成功", "删除协作者成功")
	return
}
