package userservice

import (
	global "../global"
	model "../model"
	"encoding/json"
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedUserServiceServer
}

func(s *Server)Login(ctx context.Context,message *Message)(*Message,error){
	input := new(model.User)
	user := new(model.User)

	body := string(message.Body)
	json.Unmarshal([]byte(body),&input)
	global.GVA_LOG.Info(body)
	has,err := global.GVA_DB.Table("user").Where("id = ? and password = ?",input.Id,input.PassWord).Get(user)
	if err == nil {
		if has {
			res,_:=json.Marshal(user)
			return &Message{Code: 200,Body:string(res)},nil
		}else {
			return &Message{Code: 404},nil
		}
	}else{
		global.GVA_LOG.Error("DB error:",err)
		return &Message{Code: 500},nil
	}
}
func(s *Server)Register(ctx context.Context,message *Message)(*Message,error){
	input := new(model.User)

	body := string(message.Body)
	json.Unmarshal([]byte(body),&input)
	global.GVA_LOG.Info(body)
	_,err := global.GVA_DB.Table("user").Insert(input)
	if err == nil {
		return &Message{Code: 200},nil
	}else{
		global.GVA_LOG.Error("DB error:",err)
		return &Message{Code: 500},nil
	}
}
func(s *Server)Edit(ctx context.Context,message *Message)(*Message,error){
	input := new(model.User)
	body := string(message.Body)
	json.Unmarshal([]byte(body),&input)
	global.GVA_LOG.Info(body)
	_,err:=global.GVA_DB.Table("user").Where("id = ?",input.Id).Cols("password","nickname","profilepic").Update(input)
	if err == nil {
			return &Message{Code: 200},nil
	}else{
		global.GVA_LOG.Error("DB error:",err)
		return &Message{Code: 500,},nil
	}
}

func(s *Server)Info(ctx context.Context,message *Message)(*Message,error){
	input := new(model.User)
	user := new(model.User)
	body := string(message.Body)
	json.Unmarshal([]byte(body),&input)
	global.GVA_LOG.Info(body)
	res := global.GVA_REDIS.Get(input.Id)
	if res != nil {
		jsonStr,_ := json.Marshal(res)
		global.GVA_LOG.Info(jsonStr)
		return &Message{Code: 200,Body: string(jsonStr)},nil
	}else {
		_,err:=global.GVA_DB.Table("user").Where("id = ?",input.Id).Cols("id","nickname","profilepic").Get(user)
		if err != nil {
			global.GVA_LOG.Error("DB error:",err)
			return &Message{Code: 500,},nil
		}else {
			res,_:=json.Marshal(user)
			return &Message{Code: 200,Body: string(res)},nil
		}
	}
	//_,err:=global.GVA_DB.Table("user").Where("id = ?",input.Id).Cols("id","nickname","profilepic").Get(user)
	//if err == nil {
	//	res,_:=json.Marshal(user)
	//	return &Message{Code: 200,Body: string(res)},nil
	//}else{
	//	global.GVA_LOG.Error("DB error:",err)
	//	return &Message{Code: 500,},nil
	//}
}