package handler

import (
	"apiBook/common/cache"
	"apiBook/common/log"
	"apiBook/common/utils"
	"encoding/json"
)

func setDocCache(pId, docId string, data *DocumentItemResp) {
	value, err := utils.AnyToJsonB(data)
	if err != nil {
		log.Error(err)
	}

	_ = cache.GetCache().Set(cache.DocCacheKey(pId, docId), value)
}

func getDocCache(pId, docId string) (*DocumentItemResp, bool) {
	data := &DocumentItemResp{}
	value, has := cache.GetCache().Get(cache.DocCacheKey(pId, docId))

	v, ok := value.([]byte)
	if !ok {
		has = false
	}

	if has {
		err := json.Unmarshal(v, &data)
		if err != nil {
			log.Error(err)
		}
	}

	return data, has
}

func delDocCache(pId, docId string) {
	cache.GetCache().Delete(cache.DocCacheKey(pId, docId))
}

func setDirCache(pid string, list []*DocumentDirListItem) {
	value, err := utils.AnyToJsonB(list)
	if err != nil {
		log.Error(err)
	}

	_ = cache.GetCache().Set(cache.DirCacheKey(pid), value)
}

func getDirCache(pid string) []*DocumentDirListItem {
	list := make([]*DocumentDirListItem, 0)
	data, has := cache.GetCache().Get(cache.DirCacheKey(pid))

	if v, ok := data.([]byte); has && ok {
		err := json.Unmarshal(v, &list)
		if err != nil {
			log.Error(err)
		}
	}

	return list
}

func delDirCache(pid string) {
	cache.GetCache().Delete(cache.DirCacheKey(pid))
}
