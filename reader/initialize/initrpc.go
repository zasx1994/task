package initialize

import (
	"../userservice"
	"fmt"
	"google.golang.org/grpc"
	"net"
	global "../global"
)

func GRPC(){
	global.GVA_LOG.Info("RPC服务注册")
	lis,err := net.Listen("tcp",":9000")

	if err != nil {
		fmt.Println("failed to listen:",err)
	}

	grpcService := grpc.NewServer()

	s := userservice.Server{}

	userservice.RegisterUserServiceServer(grpcService,&s)


	err = grpcService.Serve(lis)

	global.GVA_LOG.Info("RPC服务已注册")

	if err != nil {
		fmt.Println("failed to listen",err)
	}
}
