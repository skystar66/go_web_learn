package main

import (
	"fmt"
	"my-hello/pubsub"
	"strings"
	"time"
)

func main() {

	//创建发布者
	publisher := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer publisher.Close()
	//订阅所有主题
	all:=publisher.Subscribe()
	//创建订阅者
	golang:=publisher.SubscriberTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			//过滤器 过滤 包含 golang的消息
			return strings.Contains(s, "golang")
		}
		return false
	})
	publisher.Publisher("hello,xl")
	publisher.Publisher("hello,golang")
	go func() {
		for  msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
