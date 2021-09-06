package service

import "net/rpc"

//rpc 服务名称
const HelloServiceRpcName = "grpc/v2/server/service/HelloServiceRpc"

type HelloServiceRpc = interface {
	Hello(request string, replay *string) error
}

//注册rpc服务
func RegisterHelloService(helloRpcService HelloServiceRpc) error {
	return rpc.RegisterName(HelloServiceRpcName, helloRpcService)
}



