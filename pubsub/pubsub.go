package pubsub

import (
	"sync"
	"time"
)

//订阅者对象
type (
	subscriber chan interface{}         //订阅者为一个通道
	topicFunc  func(v interface{}) bool //topic 过滤器
)

//发布者对象
type Publisher struct {
	mu          sync.Mutex
	buffer      int64                    //队列的缓冲大小
	timeout     time.Duration            //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

//构建一个发布者对象，可以设置发布超时时间，缓存队列大小
func NewPublisher(timeout time.Duration, buffers int64) *Publisher {
	return &Publisher{
		buffer:      buffers,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscriberTopic(nil)
}

//添加一个新的订阅者
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{} {
	//创建订阅通道 带有缓冲的
	ch := make(chan interface{}, p.buffer)
	p.mu.Lock()
	defer p.mu.Unlock()
	//设置订阅者
	p.subscribers[ch] = topic
	return ch
}

//发布一个主题
func (p *Publisher) Publisher(v interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendToTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//关闭发布者对象
func (p *Publisher) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for sub, _ := range p.subscribers {
		//移除订阅者
		delete(p.subscribers, sub)
		//关闭通道channel
		close(sub)
	}
}

//发送主题，允许有一定的超时时间
func (p *Publisher) sendToTopic(
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	//topic 过滤校验
	if topic != nil && !topic(v) {
		return
	}
	//基于select实现的管道的超时判断
	select {
	case sub <- v:
	case <-time.After(p.timeout): //// 超时
	}

}
