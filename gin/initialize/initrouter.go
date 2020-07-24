package initialize

import (
	"../router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	var Router = gin.Default()
	router.InitUserRouter(Router)
	return Router
}