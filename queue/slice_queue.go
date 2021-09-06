package main

import "fmt"

func main() {

	//创建并获得一个队列
	queue := NewQueue()
	//向队列中添加元素
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}
	fmt.Println("size:",queue.Size())
	fmt.Println("获取并移除队列中最前面的元素",queue.Poll())
	fmt.Println("size:",queue.Size())
}

//队列中的元素
type Elemet interface{}
type Queue interface {
	Offer(e Elemet) //向队列中添加元素
	Poll() Elemet   //获取并移除队列中最前面的元素
	Clear()         //清空队列
	Size() int      //队列大小
	IsEmpty() bool  //队列是否为空
}
type sliceEntrty struct {
	elemens []Elemet //切片实现的队列
}

//创建获取一个队列
func NewQueue() *sliceEntrty {
	return &sliceEntrty{}
}

//向队列中添加元素
func (receiver *sliceEntrty) Offer(e Elemet) {
	receiver.elemens = append(receiver.elemens, e)
}

//获取并移除队列中最前面的元素
func (receiver *sliceEntrty) Poll() Elemet {
	if receiver.IsEmpty() {
		return nil
	}
	firstElement := receiver.elemens[0]
	receiver.elemens = receiver.elemens[1:]
	return firstElement
}

//校验队列是否为空
func (receiver *sliceEntrty) IsEmpty() bool {
	return len(receiver.elemens) <= 0
}

//获取队列大小
func (receiver *sliceEntrty) Size() int {
	if receiver.IsEmpty() {
		return 0
	}
	return len(receiver.elemens)
}

//清空队列
func (receiver *sliceEntrty) Clear() {
	for i := 0; i < len(receiver.elemens); i++ {
		receiver.elemens[i] = nil
	}
	receiver.elemens = nil
}
