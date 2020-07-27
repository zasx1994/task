package api

import (
	"../global"
	request "../model/request"
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
	response,err:=global.GVA_RPC.Register(context.Background(),&m)

	if err!= nil{
		fmt.Println("error when calling register",err)
	}
	code := response.Code

	
	fmt.Printf("response from server : %s",response)

	return
}

func Login(c *gin.Context){
	return
}

func Edit(c *gin.Context){
	return
}

func GetInfo(c *gin.Context){
	return
}