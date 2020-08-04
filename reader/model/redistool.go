package model

import (
	"context"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)


type Redistool struct {
	REDIS_CIL *redis.Client
	DB_CIL  *xorm.Engine
}

func NewRedistool(redis *redis.Client,DB *xorm.Engine)*Redistool{
	var result = new(Redistool)
	result.DB_CIL=DB
	result.REDIS_CIL = redis
	return result
}

func (c* Redistool)Flush()  {
	var users []User
	c.DB_CIL.Table("user").Find(&users)

	pip:= c.REDIS_CIL.Pipeline()
	for _,user := range users{
		pip.HMSet(context.Background(),user.Id,"Id",user.Id,"NickName",user.NickName,"ProfilePic",user.ProfilePic)
	}
	pip.Exec(context.Background())
	pip.Close()
}

func (c* Redistool)Get(id string)(map[string]string){
	res,err:=c.REDIS_CIL.HGetAll(context.Background(),id).Result()
	if err == nil {
		return res
	}
	return nil
}

func (c* Redistool)Delete(id string){
	c.REDIS_CIL.Del(context.Background(),id).Result()
}
