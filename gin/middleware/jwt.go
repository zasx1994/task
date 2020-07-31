package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)



func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token =="" {

			c.Abort()
		}
		fmt.Println(token)

		// 处理请求
		c.Next()
	}
}