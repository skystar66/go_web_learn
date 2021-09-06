package main

import "sync"
var(
	onceinstance *singleton
	once sync.Once

)



func ObceInstance() *singleton {
	once.Do(func() {
		onceinstance = &singleton{}
	})
	return onceinstance
}

func main() {









}
