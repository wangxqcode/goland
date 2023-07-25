package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//func main() {
//	ch := make(chan int)
//	go recv(ch) // 创建一个 goroutine 从通道接收值
//	ch <- 10
//	fmt.Println("发送成功")
//}

func recevice(c chan int)  {

	ret := <-c
	fmt.Println("accept", ret)

}

func main() {

	ch := make(chan int)
	go recevice(ch)
	ch <- 10
	fmt.Println("send",ch)
}