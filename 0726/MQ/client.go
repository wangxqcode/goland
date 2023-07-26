package mq

// client 要实现消息队列的容量 关闭  消息 锁
type Client struct {
	bro *BrokerImpl
}

func NewClient() *Client {

	return &Client{
		bro: NewBroker(),
	}

}

func (c *Client) SetConditions(capacity int) {
	c.bro.setConditions(capacity)

}

func (c *Client) Publish(topic string, msg interface{}) (<-chan interface{}, error)  {
	return c.bro.publish(topic,msg)


}

func (c *Client) Subscribe(topic string ) (sub <-chan interface{}) {
	return c.bro.subscribe(topic)

}

func (c *Client) Unsubscribe(topic string, sub <-chan interface{}) error {
	return c.bro.unsubscribe(topic, sub)

}

func (c *Client) Close() {
	return c.bro.close


}



//func (c *Client)SetConditions(capacity int)  {
//	c.bro.setConditions(capacity)
//}

//func (c *Client)Publish(topic string, msg interface{}) error{
//	return c.bro.publish(topic,msg)
//}

//func (c *Client)Subscribe(topic string) (<-chan interface{}, error){
//	return c.bro.subscribe(topic)
//}

//func (c *Client)Unsubscribe(topic string, sub <-chan interface{}) error {
//	return c.bro.unsubscribe(topic,sub)
//}

//func (c *Client)Close()  {
//	c.bro.close()
//}

func (c *Client)GetPayLoad(sub <-chan interface{})  interface{}{
	for val:= range sub{
		if val != nil{
			return val
		}
	}
	return nil
}
