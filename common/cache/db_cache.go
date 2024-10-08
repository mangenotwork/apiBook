package cache

import (
	"fmt"
	"sync"
	"time"
)

var (
	onceCache sync.Once
	memCache  *MemCache
)

func GetCache() *MemCache {
	onceCache.Do(func() {
		memCache = NewCache(24*7*time.Hour, 24*1*time.Hour)
	})
	return memCache
}

var (
	dirCacheKey = "dirCache-%s"
	docCacheKey = "docCache-%s_%s"
)

func DirCacheKey(pid string) string {
	return fmt.Sprintf(dirCacheKey, pid)
}

func DocCacheKey(pid, docId string) string {
	return fmt.Sprintf(docCacheKey, pid, docId)
}
