package main

import (
	"log"
	"my-hello/app/grpc/server/server"
	"net"
	"net/rpc"
)

func main() {

	rpc.RegisterName("HelloService",new(server.HelloService))

	listerer,err:=net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn,err:=listerer.Accept()

	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)


	//其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在“HelloService”服务空间之下。
	//然后我们建立一个唯一的TCP链接，并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。

}


