package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	mu          sync.Mutex
	initialized int64
)

func main() {

	//newInstance:=Instance()

}

func Instance() *singleton {

	if atomic.LoadInt64(&initialized) == 1 {
		return instance
	}

	mu.Lock()

	defer mu.Unlock()

	if instance ==nil {
		defer atomic.StoreInt64(&initialized,1)
		instance = &singleton{}
	}
	return instance
}
