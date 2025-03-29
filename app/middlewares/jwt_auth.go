package middlewares

import (
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// JwtAuth 基于JWT的认证中间件
func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// Token放在Header的Authorization中
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.MakeFail(c, http.StatusBadRequest, "Authroization is empty")
			c.Abort()
			return
		}

		glog.Infoln("token:", token)

		userClaims, err := service.ParseToken(token)
		if err != nil {
			response.MakeFail(c, http.StatusBadRequest, "invalid token")
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文gin.context中
		c.Set("userID", userClaims.UserID)
		c.Set("userType", userClaims.UserType)
		c.Next() // 后续的处理函数可以用过c.Get("userAccount")来获取当前请求的用户信息
	}
}
