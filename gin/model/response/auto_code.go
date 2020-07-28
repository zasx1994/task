package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Response struct {
	Code int `json:"code"`
	Data interface{} `json:data`
	Msg string `json:"msg"`
}

const (
	ERROR = 500
	SUCCES = 200
	FAILED = 400
)

func Result(code int,data interface{},msg string,c *gin.Context){
	c.JSON(http.StatusOK,Response{
		Code:code,
		Data: data,
		Msg: msg,
	})
}

func Ok(c *gin.Context){
	Result(SUCCES,map[string]interface{}{},"操作成功",c)
}

func OkWithMessage(c *gin.Context,msg string){
	Result(SUCCES,map[string]interface{}{},msg,c)
}

func Fail(c *gin.Context){
	Result(FAILED,map[string]interface{}{},"操作失败",c)
}

func FailWithMessage(c *gin.Context,msg string){
	Result(FAILED,map[string]interface{}{},msg,c)
}

