package main

import (
	"fmt"
	"sync"
)

func main() {
	ch :=make(chan int, 10)

	// 2个生产者往channel写入元素
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <10; i++ {
			ch <-i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <10; i++ {
			ch <-i
		}
	}()

	// 1个消费者
	//wg2 := sync.WaitGroup{}
	mc := make(chan struct{},0)
	go func() {
		sum := 0

		//wg2.Add(1)

		for  {
			//a := <- ch
			//sum += a
			a, ok := <-ch
			fmt.Println(ok)
			if !ok{		// channel 为空并且 channel 被关闭时候，OK才为false
				break
			}else {
				sum += a
			}
		}
		fmt.Println(sum)
		// 帮助main解除阻塞
		mc <- struct {}{}  // 使用空结构体进行阻塞，不占用内存，可读性也高
		//wg2.Done()

	}()

	wg.Wait()

	close(ch) 	// 关闭管道之后只是不能写入，还是可以读出的
	// main程序等待协程完成，设置阻塞
	<- mc
	//time.Sleep(time.Second* 10)

	//wg2.Wait()

}
