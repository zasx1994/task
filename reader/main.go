package main

import (
	"./initialize"
	"fmt"
)


func main(){
	fmt.Println("starting...")
	initialize.Logger()
	initialize.Mysql()
	initialize.Redis()
	initialize.GRPC()
}
