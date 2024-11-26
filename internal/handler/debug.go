package handler

import (
	"apiBook/common/conf"
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

	return
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
