package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"io"
	"os"
	"time"

	"../initialize"
	"net/http"
)

func Server(){
	cfg, err := ini.Load("config/config.ini")

	if err != nil {
		fmt.Print(err)
	}

	//获取一个类型为字符串（string）的值
	fmt.Print("HTTP PORT:", cfg.Section("server").Key("HTTP_PORT").String())

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router :=initialize.Router()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.Section("server").Key("HTTP_PORT").String()),
		Handler:        router,
		ReadTimeout:    10*time.Second,
		WriteTimeout:   10*time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err =s.ListenAndServe()

	if err != nil {
		fmt.Print(err)
	}

}
