package main

import (
	"fmt"
	"time"
)

var done = make(chan bool)
var msg string

var limit = make(chan int, 3)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second) },
	func() { println("2"); time.Sleep(1 * time.Second) },
	func() { println("3"); time.Sleep(1 * time.Second) },
	func() { println("4"); time.Sleep(1 * time.Second) },
	func() { println("5"); time.Sleep(1 * time.Second) },
}



func aGoroutine() {
	msg = "你好, 世界"
	time.Sleep(5*time.Second)
	done <- true

	//close(done) 关闭通道

}


//channel 是异步阻塞式通道
func main() {

	done := make(chan int,1)

	go func(){
		fmt.Println("你好, 世界")
		time.Sleep(5*time.Second)
		<-done
	}()

	done <- 1
	fmt.Println("success！！！")



	//for _, w := range work {
	//	go func(w func()) {
	//		limit <- 1
	//		w()
	//		<-limit
	//	}(w)
	//}
	//select{}
	//println(time.Now().String())
	//go aGoroutine()
	//<-done
	//println(msg)
	//println(time.Now().String())
}
