package main

import "fmt"

/*
// 无缓冲通道

func main() {

	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- "goland factory"
	}()

	go func() {
		defer wg.Done()
		a := <-c
		//time.Sleep(time.Second)
		fmt.Println(a)
	}()

	wg.Wait()

}


 */

/*
// 有缓冲通道

func main() {
	c := make(chan string,2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		c <- `Golang梦工厂`
		c <- `asong`

	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`公众号: `+ <-c)
		println(`作者: `+ <-c)
	}()
	fmt.Println(&c)
	wg.Wait()
}

 */









/*
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

 */


// 单向通道

/*
// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()

	res := Consumer(ch)
	fmt.Println(res) // 25

}

 */

// <- chan int // 只接收通道，只能接收不能发送
// chan <- int // 只发送通道，只能发送不能接收

// Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch2 := Producer2()

	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25
}


// channel操作
//发送 ch <- 10
//接收 <-ch 或 a := <- ch
//关闭 close(ch)



