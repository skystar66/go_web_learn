package main

import (
	"fmt"
	"log"
	"my-hello/app/grpc/v2/server/service"
	"net/rpc"
)

func main() {
	//客户端 调用
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call(service.HelloServiceRpcName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}




