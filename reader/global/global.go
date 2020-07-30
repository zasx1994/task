package global

import (
	"../model"
	log "github.com/cihub/seelog"
	"xorm.io/xorm"
)

var (
	GVA_DB  *xorm.Engine
	GVA_LOG log.LoggerInterface
	GVA_REDIS *model.Redistool
)
