package initialize

import (
	"fmt"
	"google.golang.org/grpc"
	"../rpcservice"
	"../global"
)

func GRPC(){
	var conn *grpc.ClientConn
	conn,err := grpc.Dial(":8000",grpc.WithInsecure())

	if err != nil {
		fmt.Printf("could not connect: %s",err)
	}

	defer conn.Close()

	c := rpcservice.NewUserServiceClient(conn)

	global.GVA_RPC = c
}
