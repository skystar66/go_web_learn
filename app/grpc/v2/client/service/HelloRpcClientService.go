package service

import (
	"my-hello/app/grpc/v2/server/service"
	"net/rpc"
)

type HelloRpcClientService struct {
	rpc *rpc.Client
}

//连接服务端，获取服务端rpc
func DailServerClient(network string, addr string) (*HelloRpcClientService, error) {

	client, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &HelloRpcClientService{
		rpc: client,
	}, nil
}


//调用服务端hello方法
func (s *HelloRpcClientService) Hello(request string, reply *string) error  {
	err:=s.rpc.Call(service.HelloServiceRpcName+".Hello",request,&reply)

	return err
}




