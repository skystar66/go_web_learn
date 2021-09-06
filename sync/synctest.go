package main



import (
	"fmt"
	"sync"
)



func main() {

	group:=sync.WaitGroup{}
	group.Add(2)
	go worker(&group)
	go worker(&group)

	group.Wait()

	fmt.Printf("value：%d \n",total.value)


}

var total struct {
	sync.Mutex
	value int
}


func worker(group *sync.WaitGroup) {

	defer group.Done()


	for i := 0; i < 100; i++ {
		total.Lock()
		total.value+=i
		total.Unlock()
	}
}
