package api

import (
	"../global"
	request "../model/request"
	response "../model/response"
	rpcservice "../rpcservice"
	"../util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func Register(c *gin.Context){
	var R request.RegisterStruct
	_= c.BindJSON(&R)
	b,_ :=json.Marshal(R)
	m := rpcservice.Message{Body: string(b)}
	res,err:=global.GVA_RPC.Register(context.Background(),&m)

	if err!= nil{
		fmt.Println("error when calling register",err)
		response.Fail(c)
		return
	}
	code := res.Code
	fmt.Printf("response from server : %s",code)
	if code == 200 {
		response.Ok(c)
		return
	}else{
		response.Fail(c)
	}
}

func Login(c *gin.Context){
	var R request.LoginStruct
	_= c.BindJSON(&R)
	b,_ :=json.Marshal(R)
	m := rpcservice.Message{Body: string(b)}
	res,err:=global.GVA_RPC.Login(context.Background(),&m)

	if err!= nil{
		fmt.Println("error when calling register",err)
		response.Fail(c)
		return
	}
	code := res.Code
	fmt.Printf("response from server : %s",code)
	if code == 200 {
		//签发token
		token := util.Encrypt("321423u9y8d2fwfl",R.Id)
		response.OkWithData(c,token)
		return
	}else{
		response.Fail(c)
		return
	}
}

func Edit(c *gin.Context){
	var R request.EditStruct
	_= c.BindJSON(&R)
	b,_ :=json.Marshal(R)
	m := rpcservice.Message{Body: string(b)}
	res,err:=global.GVA_RPC.Edit(context.Background(),&m)

	if err!= nil{
		fmt.Println("error when calling register",err)
		response.Fail(c)
		return
	}
	code := res.Code
	fmt.Printf("response from server : %s",code)
	if code == 200 {
		response.Ok(c)
		return
	}else{
		response.Fail(c)
		return
	}
}

func GetInfo(c *gin.Context){
	var R request.InfoStruct
	_= c.BindJSON(&R)
	w,_ := c.Get("claim")
	//权限检查
	if  R.Id!= w {
		response.FailWithMessage(c,"权限错误")
		return
	}

	b,_ :=json.Marshal(R)
	m := rpcservice.Message{Body: string(b)}
	res,err:=global.GVA_RPC.Info(context.Background(),&m)

	if err!= nil{
		fmt.Println("error when calling register",err)
		response.Fail(c)
		return
	}
	code := res.Code
	fmt.Printf("response from server : %s",code)
	if code == 200 {
		response.OkWithData(c,res.Body)
		return
	}else{
		response.Fail(c)
		return
	}
}