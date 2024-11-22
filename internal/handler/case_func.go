package handler

import (
	"apiBook/common/docIE"
	"apiBook/common/fenci"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func CaseFuncGo(c *gin.Context) {

	ctx := ginHelper.NewGinCtx(c)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error(err)
	}
	text := string(body)
	log.Info(text)

	for _, v := range strings.Split(text, "\n") {

		if isHaveStr(`(?is:type.*?struct)`, v) || v == "}" {
			continue
		}

		log.Info("lin: ", v)
		field := ""
		varType := ""
		description := ""

		section := make([]string, 0)

		for _, vItem := range strings.Split(v, " ") {
			if len([]byte(vItem)) > 0 {
				section = append(section, vItem)
			}
		}

		field = utils.CleaningStr(section[0])

		if len(section) > 1 {

			section1 := utils.CleaningStr(section[1])

			switch section1 {
			case "int64", "int", "int32", "uint64", "uint8", "uint16", "uint32", "float32", "float64":
				varType = "number"
			case "string":
				varType = "string"
			case "bool":
				varType = "boolean"
			}

			if varType == "" {

				if strings.Contains(section1, "*") {
					varType = "object"
				}
				if strings.Contains(section1, "[]") {
					varType = "array"
				}

			}
		}

		zsReg := regexp.MustCompile(`(?is:\/\*(.*?)\*\/)`)
		zsList := zsReg.FindAllStringSubmatch(v, -1)
		zsReg.FindStringSubmatch(v)

		if len(zsList) > 0 {
			if len(zsList[0]) > 0 {
				log.Info("zsList = ", zsList[0][len(zsList)])
				description = zsList[0][len(zsList)]
			}
		}

		zsSplit := strings.Split(v, "//")
		if len(zsSplit) > 1 {
			description = zsSplit[len(zsSplit)-1]
		}

		reg := regexp.MustCompile(`(?is:json:"(.*?)")`)
		list := reg.FindAllStringSubmatch(v, -1)
		reg.FindStringSubmatch(v)
		if len(list) > 0 {
			if len(list[0]) > 0 {
				field = list[0][len(list)]
			}
		}

		log.Info("field = ", field, " | varType = ", varType, " | description = ", description)
		// 判断是否含有  `json`
		// 否则取结构体名
		// 取类型
		// 取注释

	}

	ctx.APIOutPut("保存成功", "保存成功")
	return

}

// isHaveStr 是否含有正则匹配的字符
func isHaveStr(regStr, rest string) bool {
	isHave, err := regexp.MatchString(regStr, rest)
	if err != nil {
		return false
	}
	return isHave
}

func CaseFenCi(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	str := ctx.Query("str")

	data := fenci.TermExtract(str)

	ctx.APIOutPut(data, "ok")
}

func CaseSearch(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	str := ctx.Query("str")
	pid := ctx.Query("pid")

	log.Info("str = ", str)
	log.Info("pid = ", pid)

	list := make([]*entity.InvertIndex, 0)

	strList := fenci.TermExtract(str)
	for _, v := range strList {

		log.Info("v.Text = ", v.Text)

		item, err := dao.NewInvertIndexDao().Get(pid, v.Text)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Info("item = ", item)

		list = append(list, item...)
	}

	ctx.APIOutPut(list, "ok")
}

func CaseExportApiBook(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")

	obj, err := docIE.NewDocExport(define.SourceSwagger)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	resp := obj.Export(pid)

	//ctx.APIOutPut(resp, "ok")

	ctx.IndentedJSON(http.StatusOK, resp)

	return
}

func CasePanic(c *gin.Context) {
	panic(0)
}
