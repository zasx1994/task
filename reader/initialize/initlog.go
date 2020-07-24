package initialize

import (
	global "../global"
	"fmt"
	log "github.com/cihub/seelog"
)


func Logger()  {
	defer log.Flush()
	//加载配置文件
	logger, err := log.LoggerFromConfigAsFile("config/logconfig.xml")

	if err!=nil{
		fmt.Println("../config/logconfig.xml")
	}

	global.GVA_LOG = logger
	global.GVA_LOG.Info("日志已启动")

}
