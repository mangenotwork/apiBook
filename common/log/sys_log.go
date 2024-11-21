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
		Info(fName)
		file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer func() {
			_ = file.Close()
		}()
		_, err = file.Write(logContent("sys", content))
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
		_, err = file.Write(logContent("operation", fmt.Sprintf("[userAccount:%s]%s", userAccount, content)))
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
		_, err = file.Write(logContent("event", content))
		if err != nil {
			return
		}
	}()
}
