package middleware

import (
	"../model/response"
	"../util"
	"github.com/gin-gonic/gin"
)


//自定义中间件第1种定义方式
func JWT(c *gin.Context)  {
	token := c.Request.Header.Get("token")
	if token =="" {
		response.FailWithMessage(c,"未登陆")
		//c.Abort()
		c.Abort()
		return
	}
	claim := util.Decrypt("321423u9y8d2fwfl",token)
	c.Set("claim",claim)
	// 处理请求
	c.Next()
}