package initialize

import (
	"../global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"xorm.io/core"
	"xorm.io/xorm"
)


func Mysql() {
	cfg, err := ini.Load("config/config.ini")

	if err != nil {
		global.GVA_LOG.Error(err)
	}

	//获取一个类型为字符串（string）的值
	global.GVA_LOG.Info("MySQL IP:", cfg.Section("mysql").Key("ip").String())
	//获取一个类型为字符串（string）的值
	global.GVA_LOG.Info("MySQL port:", cfg.Section("mysql").Key("port").String())

	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	database := cfg.Section("mysql").Key("database").String()
	user :=cfg.Section("mysql").Key("user").String()
	password :=cfg.Section("mysql").Key("password").String()
	charset := cfg.Section("mysql").Key("charset").String()

	mysqllink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",user,password,ip,port,database,charset)

	engine,err := xorm.NewEngine("mysql",mysqllink)

	if err != nil {
		global.GVA_LOG.Error(err)
	}
	global.GVA_LOG.Info("数据库已经链接")
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	global.GVA_DB = engine
}