package main

import (
	"../gin/initialize"
	"./core"
)

func main(){
	initialize.GRPC()
	core.Server()
}