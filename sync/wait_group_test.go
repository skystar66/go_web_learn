package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- `Golang梦工厂`
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println(`Message: `+ <-c)
	}()
	wg.Wait()
}






