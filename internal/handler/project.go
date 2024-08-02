package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func ProjectList(c *gin.Context) {
	ginHelper.NewGinCtx(c)
}

func ProjectItem(c *gin.Context) {

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
	log.Println(userAcc)

	param.ProjectId = utils.IDMd5()
	param.CreateDate = utils.NowDate()
	err = dao.NewProjectDao().Create(param, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, "创建失败")
		return
	}

	ctx.APIOutPut("创建成功", "创建成功")
	return
}

func ProjectModify(c *gin.Context) {

}

func ProjectDelete(c *gin.Context) {

}

func ProjectUsers(c *gin.Context) {

}

func ProjectAddUser(c *gin.Context) {

}
