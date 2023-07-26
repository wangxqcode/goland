package main

import "fmt"
// select 语句特点：
//1.可以处理一个或者多个channel的发送或者接收
//2. 如果多个case同时满足，select会随机选择一个执行
//3. 对于没有case的select 会一直阻塞main，防止退出


func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case  ch <- i:
		}
	}
}
