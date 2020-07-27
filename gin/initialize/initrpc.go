package initialize

import (
	"../global"
	"../rpcservice"
	"fmt"
	"google.golang.org/grpc"
)

func GRPC(){
	var conn *grpc.ClientConn
	conn,err := grpc.Dial(":9000",grpc.WithInsecure())


	if err != nil {
		fmt.Printf("could not connect: %s",err)
	}


	c := rpcservice.NewUserServiceClient(conn)


	global.GVA_RPC = c

	fmt.Print("成功链接GRPC服务器")
}
