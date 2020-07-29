package router

import (
	api "../api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.Engine){
	UserRouter := Router.Group("user")

	UserRouter.POST("register",api.Register)
	UserRouter.POST("edit",api.Edit)
	UserRouter.POST("login",api.Login)
	UserRouter.POST("info",api.GetInfo)

}
