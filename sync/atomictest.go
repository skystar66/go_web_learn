package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


var value int64


func main() {


	group:=sync.WaitGroup{}
	group.Add(2)
	go atomicWorker(&group)
	go atomicWorker(&group)
	group.Wait()
	fmt.Printf("atomic value：%d \n",value)



}



func atomicWorker(group *sync.WaitGroup) {
	defer group.Done()
	var i int64
	for i = 0; i < 100; i++ {
		atomic.AddInt64(&value,i)
	}
}


