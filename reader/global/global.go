package global

import (
	log "github.com/cihub/seelog"
	"xorm.io/xorm"
)

var (

	GVA_DB  *xorm.Engine
	GVA_LOG log.LoggerInterface
)
