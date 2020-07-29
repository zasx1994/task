package api

import (
	"../global"
	request "../model/request"
	response "../model/response"
	rpcservice "../rpcservice"
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
		response.OkWithMessage(c,"登陆成功")
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
		response.OkWithData(c,res)
		return
	}else{
		response.Fail(c)
		return
	}
}