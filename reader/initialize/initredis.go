package initialize

import (
	"../global"
	"../model"
	"github.com/go-redis/redis"
)




func Redis(){
	var rdb *redis.Client
	rdb = redis.NewClient(&redis.Options{
		Addr:":6379",
		Password:"",
		DB:0,
	})
	rt := model.NewRedistool(rdb,global.GVA_DB)
	rt.Flush()
	global.GVA_REDIS = rt
}