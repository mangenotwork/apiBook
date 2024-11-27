package handler

import (
	"apiBook/common/conf"
	"apiBook/common/db"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

func SysInfo(c *gin.Context) {
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

	sysInfo := make(map[string]interface{})

	// 服务器信息
	sysInfo["hostName"] = utils.GetHostName()
	sysInfo["hostType"] = utils.GetSysType()
	sysInfo["hostArch"] = utils.GetSysArch()
	sysInfo["hostCpuCoreNumber"] = utils.GetCpuCoreNumber()
	sysInfo["hostInterfaceInfo"] = utils.GetInterfaceInfo()
	sysInfo["hostIP"] = utils.GetIP()
	sysInfo["webServer"] = utils.GetIP() + ":" + conf.Conf.Default.HttpServer.Prod

	// 配置信息
	sysInfo["confPath"] = conf.Conf.YamlPath
	sysInfo["confInfo"] = conf.Conf.YamlData

	// 项目总数量
	sysInfo["projectNum"] = dao.NewProjectDao().GetAllProjectNum()
	sysInfo["projectIdList"] = dao.NewProjectDao().GetAllProjectIdList()

	// 用户总数量，
	sysInfo["userNum"] = dao.NewUserDao().GetAllUserNum()

	// db文件大小,
	var size int64 = 0
	if dbPath, ok := conf.GetString("dbPath"); ok {

		file, err := os.Stat(dbPath)
		if err != nil {
			log.Error(err)
		}

		size += file.Size()
	}

	if invertIndexDBPath, ok := conf.GetString("invertIndexDBPath"); ok {

		file, err := os.Stat(invertIndexDBPath)
		if err != nil {
			log.Error(err)
		}

		size += file.Size()
	}

	sysInfo["dbSize"] = utils.SizeFormat(size)

	// 图片存储大小及数量，
	var mediaSize, mediaFileNum int64 = 0, 0
	if mediaPath, ok := conf.GetString("mediaPath"); ok {

		err := filepath.Walk(mediaPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				mediaFileNum++
				mediaSize += info.Size()
			}
			return nil
		})

		if err != nil {
			log.Error(err)
		}
	}

	sysInfo["mediaSize"] = utils.SizeFormat(mediaSize)
	sysInfo["mediaFileNum"] = mediaFileNum

	// 运行时间
	seconds := time.Now().Unix() - define.TimeStamp
	sysInfo["runTime"] = utils.ResolveTimeStr(seconds)

	// 日志信息
	workPath, _ := os.Getwd()
	logDirName := filepath.Join(workPath, "/logs/")
	if runtime.GOOS == "windows" {
		logDirName = strings.Replace(logDirName, "\\", "/", -1)
	}
	sysInfo["logPath"] = logDirName

	ctx.APIOutPut(sysInfo, "")
	return
}

func ProjectInfo(c *gin.Context) {
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

	selectType := ctx.GetQuery("type")
	pid := ctx.GetQuery("pid")

	switch selectType {
	case "list":

		list, err := db.DB.AllKey(db.GetProjectTable())
		if err != nil {
			log.Error(err)
			ctx.APIOutPutError(err, err.Error())
			return
		}

		ctx.APIOutPut(list, "")
		return

	default:

		if pid == "" {
			ctx.APIOutPutError(nil, "pid不能为空")
			return
		}

		data := make(map[string]interface{})
		err := db.DB.Get(db.GetProjectTable(), pid, &data)
		if err != nil {
			log.Error(err)
			ctx.APIOutPutError(err, err.Error())
			return
		}

		docList, err := db.DB.AllKey(db.GetDocumentTable(pid))
		if err != nil {
			log.Error(err)
			ctx.APIOutPutError(err, err.Error())
			return
		}

		data["docList"] = docList

		userList, err := db.DB.AllKey(db.GetProjectPrivateUserTable(pid))
		if err != nil {
			log.Error(err)
			ctx.APIOutPutError(err, err.Error())
			return
		}

		data["userList"] = userList

		ctx.APIOutPut(data, "")
		return

	}

}

func SysLog(c *gin.Context) {
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

	workPath, _ := os.Getwd()
	logDir := filepath.Join(workPath, "/logs/")

	fileList, err := utils.MatchSearchFileFromDir(logDir, "access")
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	if len(fileList) == 0 {
		ctx.APIOutPutError(fmt.Errorf("没找到日志文件"), "没找到日志文件")
		return
	}

	sort.Slice(fileList, func(i, j int) bool {
		if fileList[i] > fileList[j] {
			return true
		}
		return false
	})

	data, err := utils.ReadLastNLines(fileList[0], 1000)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut(data, "")
	return
}

func DB(c *gin.Context) {

	//dbStats := db.DB.GetDBStats()
	//log.Info(dbStats)

	bucketList, err := db.DB.GetAllBucket()
	if err != nil {
		log.Error(err)
	}
	log.Info(bucketList)
	return
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
	return
}

func DBSearchBucket(c *gin.Context) {
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

	dbName := ctx.GetQuery("dbName")
	search := ctx.GetQuery("search")

	if search == "" {
		ctx.APIOutPutError(nil, "search为空")
		return
	}

	var (
		list []string
		err  error
	)

	switch dbName {
	case "invertIndexDB":
		list, err = db.InvertIndexDB.SearchAllBucket(search)
	default:
		list, err = db.DB.SearchAllBucket(search)
	}

	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut(list, "")
	return
}

func DBSelect(c *gin.Context) {
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

	dbName := ctx.GetQuery("dbName")
	bucket := ctx.GetQuery("bucket")
	selectType := ctx.GetQuery("selectType")
	key := ctx.GetQuery("key")
	search := ctx.GetQuery("search")

	// bucket=dir:12653102d4686a68&selectType=allKey
	// bucket=dir:12653102d4686a68&selectType=searchKey&search=default
	// bucket=dir:12653102d4686a68&key=default_12653102d4686a68

	var (
		data = make(map[string]any)
		list = make([]string, 0)
		flag = 1 // 1:data  2:list
		err  error
	)

	switch selectType {
	case "allKey": // 获取所有的key
		flag = 2
		switch dbName {
		case "invertIndexDB":
			list, err = db.InvertIndexDB.AllKey(bucket)
		default:
			list, err = db.DB.AllKey(bucket)
		}
		break

	case "searchKey": // 搜索key
		flag = 2
		switch dbName {
		case "invertIndexDB":
			list, err = db.InvertIndexDB.SearchKey(bucket, search)
		default:
			list, err = db.DB.SearchKey(bucket, search)
		}
		break

	default:
		flag = 1
		switch dbName {
		case "invertIndexDB":
			err = db.DB.Get(bucket, key, &data)
		default:
			err = db.DB.Get(bucket, key, &data)
		}
		break
	}

	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	switch flag {
	case 1:
		ctx.APIOutPut(data, "")
	case 2:
		ctx.APIOutPut(list, "")
	}

	return
}
