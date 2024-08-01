package routers

import (
	"apiBook/common/conf"
	"apiBook/common/ginHelper"
	"apiBook/common/utils"
	"apiBook/internal/define"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthPG 权限验证中间件
func AuthPG() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, _ := c.Cookie(define.UserToken)

		j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
		if err := j.ParseToken(token); err == nil {
			c.Next()
			return
		}

		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	}
}

// AuthAPI 权限验证中间件
func AuthAPI() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, _ := c.Cookie(define.UserToken)

		j := utils.NewJWT(conf.Conf.Default.Jwt.Secret, conf.Conf.Default.Jwt.Expire)
		if err := j.ParseToken(token); err == nil {
			c.Next()
			return
		}

		ginHelper.AuthErrorOut(c)
		c.Abort()
		return

	}
}
