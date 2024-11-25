package handler

import (
	"apiBook/common/conf"
	"apiBook/common/ginHelper"
	"apiBook/internal/dao"
	"github.com/gin-gonic/gin"
)

func SysInfo(c *gin.Context) {

}

func ProjectInfo(c *gin.Context) {

}

func SysLog(c *gin.Context) {

}

func DB(c *gin.Context) {

}

func Conf(c *gin.Context) {
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

	data := make(map[string]interface{})
	data["path"] = conf.Conf.YamlPath
	data["conf"] = conf.Conf.YamlData

	ctx.APIOutPut(data, "")
}
