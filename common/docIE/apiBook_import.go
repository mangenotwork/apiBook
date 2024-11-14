package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"time"
)

type ApiBookImport struct {
	Project map[string]interface{} `json:"project"`
	Dir     []interface{}          `json:"dir"`
	Doc     []interface{}          `json:"doc"`
}

func NewApiBookImport() *ApiBookImport {
	return &ApiBookImport{}
}

func (obj *ApiBookImport) Whole(text, userAcc string, private define.ProjectPrivateCode) error {

	err := obj.analysis(text)
	if err != nil {
		log.Error(err)
		return err
	}

	project := obj.analysisProject(userAcc, private)

	if dao.NewProjectDao().HasName(project.Name) {
		project.Name += utils.NowDateNotLine()
	}

	err = dao.NewProjectDao().Create(project, userAcc)
	if err != nil {
		log.Error("创建项目失败: ", err)
		return err
	}

	err = dao.NewDirDao().CreateInit(project.ProjectId)
	if err != nil {
		log.Error("创建项目失败: ", err)
		return err
	}

	obj.analysisDoc(project, "user",
		func(project *entity.Project, dir *entity.DocumentDir) {
			err = dao.NewDirDao().Create(project.ProjectId, dir)
			if err != nil {
				log.Error("创建目录失败:", err)
			}

		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {
			documentData := &entity.Document{
				DocId:     doc.DocId,
				DirId:     dirId,
				ProjectId: doc.ProjectId,
				Name:      doc.Name,
				Url:       doc.Url,
				Method:    doc.Method,
			}

			err = dao.NewDocDao().Create(documentData, doc)
			if err != nil {
				log.Error("接口文档创建失败， err: ", err)
				return
			}

			dirItem := &entity.DocumentDirItem{
				DocId: doc.DocId,
				Sort:  0,
			}

			err = dao.NewDirDao().AddDoc(dirId, dirItem)
			if err != nil {
				log.Error("接口文档加入目录失败， err: ", err)
				return
			}
			log.Info("导入成功")

		},
	)

	return nil
}

func (obj *ApiBookImport) Increment(text, pid, userAcc, dirId string) error {

	project, err := dao.NewProjectDao().Get(pid, userAcc, false)
	if err != nil {
		log.Error("获取项目失败, err = ", err)
		return err
	}

	err = obj.analysis(text)
	if err != nil {
		return err
	}

	obj.analysisDoc(project, "user",
		func(project *entity.Project, dir *entity.DocumentDir) {
			err = dao.NewDirDao().Create(project.ProjectId, dir)
			if err != nil {
				log.Error("创建目录失败")
			}

		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {
			documentData := &entity.Document{
				DocId:     doc.DocId,
				DirId:     dirId,
				ProjectId: doc.ProjectId,
				Name:      doc.Name,
				Url:       doc.Url,
				Method:    doc.Method,
			}

			err = dao.NewDocDao().Create(documentData, doc)
			if err != nil {
				log.Error("接口文档创建失败， err: ", err)
				return
			}

			dirItem := &entity.DocumentDirItem{
				DocId: doc.DocId,
				Sort:  0,
			}

			err = dao.NewDirDao().AddDoc(dirId, dirItem)
			if err != nil {
				log.Error("接口文档加入目录失败， err: ", err)
				return
			}
			log.Info("导入成功")

		},
	)

	return nil
}

func (obj *ApiBookImport) analysis(text string) error {

	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (obj *ApiBookImport) analysisProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          utils.AnyToString(obj.Project["name"]),
		Description:   utils.AnyToString(obj.Project["description"]),
		CreateUserAcc: userAcc,
		CreateDate:    utils.AnyToString(obj.Project["createDate"]),
		Private:       private,
	}
}

func (obj *ApiBookImport) analysisDoc(project *entity.Project, userAcc string,
	createDir func(project *entity.Project, dir *entity.DocumentDir),
	f func(project *entity.Project, doc *entity.DocumentContent, dirId string)) {

	now := time.Now().Unix()

	for i, dirItem := range obj.Dir {
		dirItemMap := utils.AnyToMap(dirItem)
		createDir(project, &entity.DocumentDir{
			DirId:   utils.AnyToString(dirItemMap["dirId"]),
			DirName: utils.AnyToString(dirItemMap["dirName"]),
			Sort:    i,
		})
	}

	for _, docItem := range obj.Doc {
		docItemMap := utils.AnyToMap(docItem)

		doc := &entity.DocumentContent{
			DocId:           utils.AnyToString(docItemMap["docId"]),
			ProjectId:       project.ProjectId,
			Name:            utils.AnyToString(docItemMap["name"]),
			Url:             utils.AnyToString(docItemMap["url"]),
			Method:          utils.AnyToString(docItemMap["method"]),
			Description:     utils.AnyToString(docItemMap["description"]),
			DescriptionHtml: utils.AnyToString(docItemMap["description"]),
			ReqHeader:       make([]*entity.ReqHeaderItem, 0),
			ReqType:         define.ReqTypeCode(utils.AnyToString(docItemMap["json"])),
			ReqBodyJson:     utils.AnyToString(docItemMap["reqBodyJson"]),
			ReqBodyInfo:     make([]*entity.BodyInfoItem, 0),
			Resp:            make([]*entity.RespItem, 0),
			CreateTime:      now,
			UserAcc:         userAcc,
		}

		for _, headerItem := range utils.AnyToArr(docItemMap["reqHeader"]) {
			headerItemMap := utils.AnyToMap(headerItem)
			doc.ReqHeader = append(doc.ReqHeader, &entity.ReqHeaderItem{
				Field:       utils.AnyToString(headerItemMap["field"]),
				VarType:     utils.AnyToString(headerItemMap["varType"]),
				Description: utils.AnyToString(headerItemMap["description"]),
				Example:     utils.AnyToString(headerItemMap["example"]),
				IsRequired:  utils.AnyToInt(headerItemMap["isRequired"]),
				Sort:        utils.AnyToInt(headerItemMap["sort"]),
				IsOpen:      utils.AnyToInt(headerItemMap["isOpen"]),
			})
		}

		for _, reqBodyItem := range utils.AnyToArr(docItemMap["reqBodyInfo"]) {
			reqBodyItemMap := utils.AnyToMap(reqBodyItem)
			doc.ReqBodyInfo = append(doc.ReqBodyInfo, &entity.BodyInfoItem{
				Field:       utils.AnyToString(reqBodyItemMap["field"]),
				VarType:     utils.AnyToString(reqBodyItemMap["varType"]),
				Description: utils.AnyToString(reqBodyItemMap["description"]),
				Example:     utils.AnyToString(reqBodyItemMap["example"]),
				IsRequired:  utils.AnyToInt(reqBodyItemMap["isRequired"]),
				Sort:        utils.AnyToInt(reqBodyItemMap["sort"]),
				IsOpen:      utils.AnyToInt(reqBodyItemMap["isOpen"]),
			})
		}

		resp := &entity.RespItem{
			RespBodyInfo: make([]*entity.BodyInfoItem, 0),
		}

		for _, respItem := range utils.AnyToArr(docItemMap["resp"]) {
			respItemMap := utils.AnyToMap(respItem)

			resp.Tag = utils.AnyToString(respItemMap["tag"])
			resp.RespType = define.ReqTypeCode(utils.AnyToString(respItemMap["respType"]))
			resp.RespBody = utils.AnyToString(respItemMap["respBody"])

			for _, respBodyItem := range utils.AnyToArr(respItemMap["respBodyInfo"]) {
				respBodyItemMap := utils.AnyToMap(respBodyItem)
				resp.RespBodyInfo = append(resp.RespBodyInfo, &entity.BodyInfoItem{
					Field:       utils.AnyToString(respBodyItemMap["field"]),
					VarType:     utils.AnyToString(respBodyItemMap["varType"]),
					Description: utils.AnyToString(respBodyItemMap["description"]),
					Example:     utils.AnyToString(respBodyItemMap["example"]),
					IsRequired:  utils.AnyToInt(respBodyItemMap["isRequired"]),
					Sort:        utils.AnyToInt(respBodyItemMap["sort"]),
					IsOpen:      utils.AnyToInt(respBodyItemMap["isOpen"]),
				})
			}

		}

		doc.Resp = append(doc.Resp, resp)

		dirId := utils.AnyToString(docItemMap["dirId"])

		f(project, doc, dirId)

	}

}
