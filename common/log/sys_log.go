package log

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func InitSysLog() {
	workPath, _ := os.Getwd()
	logDirName := filepath.Join(workPath, "/logs/")
	_, err := os.Stat(logDirName)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logDirName, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// 日志类型
// s sys 系统
// o operation 操作
// e event 事件

func logContent(module, content string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05.000 "))
	buffer.WriteString(" \t| ")
	buffer.WriteString(module)
	buffer.WriteString(" \t| ")
	content = strings.Replace(content, "\n", " ", -1)
	buffer.WriteString(content)
	buffer.WriteString("\n")
	return buffer.Bytes()
}

func SendSysLog(content string) {
	go func() {
		workPath, _ := os.Getwd()
		logDirName := filepath.Join(workPath, "/logs/")
		now := time.Now()
		year, month, _ := now.Date()
		u := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Unix()
		fName := fmt.Sprintf("%s/access_%d.log", logDirName, u)
		file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer func() {
			_ = file.Close()
		}()
		_, err = file.Write(logContent("s", content))
		if err != nil {
			return
		}
	}()
}

func SendOperationLog(userAccount, content string) {
	go func() {
		workPath, _ := os.Getwd()
		logDirName := filepath.Join(workPath, "/logs/")
		now := time.Now()
		year, month, _ := now.Date()
		u := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Unix()
		fName := fmt.Sprintf("%s/access_%d.log", logDirName, u)
		Info(fName)
		file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer func() {
			_ = file.Close()
		}()
		_, err = file.Write(logContent("o", fmt.Sprintf("[userAccount:%s]%s", userAccount, content)))
		if err != nil {
			return
		}
	}()
}

func SendEventLog(content string) {
	go func() {
		workPath, _ := os.Getwd()
		logDirName := filepath.Join(workPath, "/logs/")
		now := time.Now()
		year, month, _ := now.Date()
		u := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Unix()
		fName := fmt.Sprintf("%s/access_%d.log", logDirName, u)
		Info(fName)
		file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer func() {
			_ = file.Close()
		}()
		_, err = file.Write(logContent("e", content))
		if err != nil {
			return
		}
	}()
}

func SendErrorLog(errContent string, stack string) {
	workPath, _ := os.Getwd()
	logDirName := filepath.Join(workPath, "/logs/")
	fName := fmt.Sprintf("%s/error.log", logDirName)
	file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()

	var buffer bytes.Buffer
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05.000 "))
	buffer.WriteString(" \t| [Error]")
	buffer.WriteString(errContent)
	buffer.WriteString(" \t| [Stack]")
	buffer.WriteString(stack)
	buffer.WriteString("\n")

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return
	}
}
