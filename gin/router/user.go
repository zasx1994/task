package router

import (
	api "../api"
	"github.com/gin-gonic/gin"
	"../middleware"
)

func InitUserRouter(Router *gin.Engine){
	UserRouter := Router.Group("user")

	UserRouter.POST("register",api.Register)
	UserRouter.POST("edit",middleware.JWT,api.Edit)
	UserRouter.POST("login",api.Login)
	UserRouter.POST("info",middleware.JWT,api.GetInfo)

}
