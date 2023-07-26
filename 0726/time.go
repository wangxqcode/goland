
package main

import (
	"fmt"
	"time"
)

//func main() {
//	//
//	t := time.NewTimer(time.Second * 5)
//	//defer t.Stop()在这里并不会停止定时器。这是因为Stop会停止Timer，停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误。
//	defer t.Stop()
//	for {
//		// 只接收通道
//		<-t.C
//		fmt.Println("timer running...")
//		// 需要重置Reset 使 t 重新开始计时
//		t.Reset(time.Second * 2)
//	}
//}
/*
func main() {

	t := time.NewTimer(time.Second * 2)

	ch := make(chan bool)
	go func() {
		defer t.Stop()
		for  {
			select {
			case <-t.C:
				fmt.Println("time is runninger")
				t.Reset(time.Second * 2)
			case stop := <- ch:
				if stop{
					fmt.Println("time is stop")
					return
				}
			}
		}
	}()
	time.Sleep(time.Second * 10)
	ch <- true
	close(ch)

}

 */

/*
// 定时器

func main() {
	t := time.NewTicker(time.Second*2)
	defer t.Stop()
	for {
		<- t.C
		fmt.Println("Ticker running...")
	}
}

 */

// time.after
/*
func main() {
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	fmt.Println("t=", <-t)

}

 */


func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	for {
		select {
		case e1 := <-ch1:
			//如果ch1通道成功读取数据，则执行该case处理语句
			fmt.Printf("1th case is selected. e1=%v\n", e1)
		case <-time.After(time.Second*2):
			fmt.Println("Timed out")
		}
	}

}






