package handler

import (
	"apiBook/common/log"
	"github.com/gin-gonic/gin"
)

func Simulator(c *gin.Context) {
	path := c.Param("path")
	log.Info(path)
	_, _ = c.Writer.Write([]byte(path))
}
