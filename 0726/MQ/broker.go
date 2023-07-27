package mq

import (
	"errors"
	"time"

	"sync"
)

// 消息队列（MQ）：

// 定义接口先把需要实现的方法列出来

type Broker interface {
	//多个发布者将消息发送到Topic,系统将这些消息传递给多个订阅者。
	publish (topic string, msg interface{}) error
	subscribe(topic string) (sub <-chan interface{}, err error)
	unsubscribe (topic string, sub <- chan interface{}, err error)
	close()
	broadcast(msg interface{}, subscribes []chan interface{})
	setConditions(capacity int)
}

type BrokerImpl struct {
	exit chan bool //关闭通道
	capacity int  //设置容量
	// 这里使用一个map结构，key即是topic，其值则是一个切片，chan类型，
	// 这里这么做的原因是我们一个topic可以有多个订阅者，所以一个订阅者对应着一个通道
	// topic msg集合
	topics map[string] []chan interface{}
	sync.RWMutex
}


func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}


func (b *BrokerImpl) broadcast(msg interface{}, subscribers []chan interface{}) {
	count := len(subscribers)
	concurrency := 1

	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}

	pub := func(start int) {
		idleDuration := 5 * time.Millisecond
		idleTimeout := time.NewTimer(idleDuration)
		defer idleTimeout.Stop()
		for j := start; j < count; j += concurrency {
			if !idleTimeout.Stop() {
				select {
				case <-idleTimeout.C:
				default:
				}
			}
			idleTimeout.Reset(idleDuration)
			select {
			case subscribers[j] <- msg:
			case <-b.exit:
				return
			}
		}
	}
	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}

func (b *BrokerImpl) publish (topic string, pub interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker close")
	default:
	}
	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil

	}
	b.broadcast(pub, subscribers)

	return nil

}

func (b *BrokerImpl) close()  {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		b.topics = make(map[string][]chan interface{})
		b.Unlock()
	}
	return
}

func (b *BrokerImpl)setConditions(capacity int)  {
	b.capacity = capacity
}



func (b *BrokerImpl) unsubscribe (topic string, sub <-chan interface{}) error {

	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}
	b.RLock()
	subscribes, ok := b.topics[topic]
	b.RUnlock()
	if !ok{
		return nil
	}

	// delete subscribe
	var newSubs [] chan interface{}
	for _, subscriber := range subscribes {
		if subscriber ==sub{
			continue
		}
		newSubs = append(newSubs,subscriber)
	}

		b.Lock()
		b.topics[topic] = newSubs
		b.Unlock()
		return nil



}








