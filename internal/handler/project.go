package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ProjectList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	userAcc := ctx.GetString("userAcc")
	resp := dao.NewProjectDao().GetList(userAcc)
	ctx.APIOutPut(resp, "")
	return
}

func ProjectItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")
	userAcc := ctx.GetString("userAcc")

	resp, err := dao.NewProjectDao().Get(pid, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut(resp, "")
	return
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
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDirDao().CreateInit(param.ProjectId)
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
		ctx.APIOutPutError(err, err.Error())
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
		List:      make([]*UserInfo, 0),
		CreateAcc: "",
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

	for _, v := range strings.Split(param.Accounts, ",") {
		err = dao.NewProjectDao().AddUser(param.PId, userAcc, v)
		if err != nil {
			log.Error(err)
		}
	}

	ctx.APIOutPut("添加协作者成功", "添加协作者成功")
	return
}

func ProjectDelUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &ProjectDelUserReq{}
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

func ProjectJoinList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	userList := GetUserList()

	pid := ctx.Query("pid")

	projectUserList, err := dao.NewProjectDao().GetUserList(pid, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	hasMap := make(map[string]struct{})
	for _, v := range projectUserList {
		hasMap[v] = struct{}{}
	}

	outList := make([]*UserInfo, 0)

	for _, u := range userList {
		if _, ok := hasMap[u.Account]; !ok {
			outList = append(outList, u)
		}
	}

	ctx.APIOutPut(outList, "")
	return
}

func ProjectHeaderAdd(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.ProjectGlobalHeader{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	err = dao.NewProjectDao().AddGlobalHeader(param.ProjectId, param.ReqHeader)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("保存成功", "保存成功")
	return
}

func ProjectHeaderGet(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")

	headerList, err := dao.NewProjectDao().GetGlobalHeader(pid)
	if err != nil {
		log.Error(err)
		headerList = make([]*entity.ReqHeaderItem, 0)
	}

	ctx.APIOutPut(headerList, "")
	return
}

func ProjectCodeAdd(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.ProjectGlobalCode{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	err = dao.NewProjectDao().AddGlobalCode(param.ProjectId, param.List)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("保存成功", "保存成功")
	return
}

func ProjectCodeGet(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")

	codeList, err := dao.NewProjectDao().GetGlobalCode(pid)
	if err != nil {
		log.Error(err)
		codeList = make([]*entity.GlobalCodeItem, 0)
	}

	ctx.APIOutPut(codeList, "")
	return
}
