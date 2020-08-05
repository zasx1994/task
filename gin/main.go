package main

import (
	"../gin/initialize"
)

func main(){
	initialize.GRPC()
	initialize.Server()
}