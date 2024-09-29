package dao

import (
	"apiBook/common/db"
	"apiBook/common/fenci"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/entity"
	"errors"
	"time"
)

type InvertIndexDao struct {
}

func NewInvertIndexDao() *InvertIndexDao {
	return &InvertIndexDao{}
}

func (dao *InvertIndexDao) Add(pid, word string, item *entity.InvertIndex) error {

	err := db.DB.Set(db.GetDocInvertIndexListTable(item.DocId), item.Word, 1)
	if err != nil {
		log.Error(err)
		return err
	}

	list := make([]*entity.InvertIndex, 0)
	err = db.DB.Get(db.GetInvertIndexTable(pid), word, &list)
	if err != nil {
		if errors.Is(err, db.ISNULL) || errors.Is(err, db.TableNotFound) {
			list = append(list, item)
			return db.DB.Set(db.GetInvertIndexTable(pid), word, list)
		} else {
			log.Error(err)
			return err
		}
	}

	has := false

	for _, v := range list {
		if v.DocId == item.DocId {
			has = true
			break
		}
	}

	if !has {
		list = append(list, item)
		return db.DB.Set(db.GetInvertIndexTable(pid), word, list)
	}

	return nil
}

func (dao *InvertIndexDao) Get(pid, word string) ([]*entity.InvertIndex, error) {
	list := make([]*entity.InvertIndex, 0)
	err := db.DB.Get(db.GetInvertIndexTable(pid), word, &list)
	if err != nil {
		return list, err
	}
	return list, err
}

func (dao *InvertIndexDao) Del(pid, docId, word string) error {
	list, err := dao.Get(pid, word)
	if err != nil {
		log.Error(err)
		return err
	}

	has := false

	for i, v := range list {
		if v.DocId == docId {
			has = true
			list = append(list[:i], list[i+1:]...)
		}
	}

	if has {
		return db.DB.Set(db.GetInvertIndexTable(pid), word, list)
	}

	return nil
}

func (dao *InvertIndexDao) DocDelAllWord(pid, docId string) error {
	wordKey, err := db.DB.AllKey(db.GetDocInvertIndexListTable(docId))
	if err != nil {
		log.Error(err)
		return err
	}

	for _, v := range wordKey {
		err = dao.Del(pid, docId, v)
		if err != nil {
			log.Error(err)
		}

		err = db.DB.Delete(db.GetDocInvertIndexListTable(docId), v)
	}

	return err
}

func (dao *InvertIndexDao) DocFenCi(pid, docId, content, modType string) {
	fcList := fenci.TermExtract(content)

	fcList = utils.SliceDeduplicate[*fenci.Term](fcList)

	now := time.Now().Unix()
	for _, v := range fcList {

		item := &entity.InvertIndex{
			DocId:      docId,
			Word:       v.Text,
			Sentence:   content,
			ModType:    modType,
			CreateTime: now,
			Term:       v,
		}

		err := dao.Add(pid, v.Text, item)
		if err != nil {
			log.Error(err)
		}

	}
}

func (dao *InvertIndexDao) DocInvertIndex(doc *entity.DocumentContent) {
	// title:标题  description:文档说明  header:请求header  req:请求参数  resp:响应参数  url:请求url

	dao.DocFenCi(doc.ProjectId, doc.DocId, doc.Url, "url")

	dao.DocFenCi(doc.ProjectId, doc.DocId, doc.Name, "title")

	dao.DocFenCi(doc.ProjectId, doc.DocId, doc.Description, "description")

	for _, v := range doc.ReqHeader {
		dao.DocFenCi(doc.ProjectId, doc.DocId, v.Field, "header")
		dao.DocFenCi(doc.ProjectId, doc.DocId, v.Description, "header")
	}

	for _, v := range doc.ReqBodyInfo {
		dao.DocFenCi(doc.ProjectId, doc.DocId, v.Field, "req")
		dao.DocFenCi(doc.ProjectId, doc.DocId, v.Description, "req")
	}

	if len(doc.Resp) > 0 {
		for _, v := range doc.Resp[0].RespBodyInfo {
			dao.DocFenCi(doc.ProjectId, doc.DocId, v.Field, "resp")
			dao.DocFenCi(doc.ProjectId, doc.DocId, v.Description, "resp")
		}
	}

}
