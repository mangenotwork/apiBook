package routers

import (
	"apiBook/common/conf"
	"apiBook/common/ginHelper"
	"apiBook/common/utils"
	"apiBook/internal/dao"
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
			acc := j.GetString("userAcc")
			c.Set("userAcc", acc)

			data, _ := dao.NewUserDao().Get(acc)
			if data.Account == "" {
				c.SetCookie(define.UserToken, "", define.TokenExpires, "/", "", false, true)
				c.Redirect(http.StatusFound, "/")
				c.Abort()
				return
			}

			c.Set("userName", data.Name)

			c.Next()
			return
		}

		c.SetCookie(define.UserToken, "", define.TokenExpires, "/", "", false, true)
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
			c.Set("userAcc", j.GetString("userAcc"))
			c.Next()
			return
		}

		ginHelper.AuthErrorOut(c)
		c.Abort()
		return

	}
}
