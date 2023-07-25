package main

import (
	"fmt"
)

/*
进程：程序在操作系统执行中，系统进行资源分配和调度的一个独立单位
线程：线程是进程中执行的一个实体，是CPU调度和执行的基本单位，他是比进程更小的单位，一个线程上可以跑多个协程
协程：独立的栈空间，共享堆空间，调度由用户自己控制

并发：多线程程序在一个核上运行就是并发
并行：多线程程序在多个核上运行 就是并行
注：并发和并行的区别在于是否同时运行

 */

/*
func main() {
	go hello()
	fmt.Println("main goroutine over")
	// main 函数默认是一条单独的goroutine，如果main结束，其他自动结束，所以有sleep等一下
	time.Sleep(time.Second)
}

func hello() {
	fmt.Println("hello goroutine")
}

 */


/*
// 启动多个goroutine 即并发
var wg sync.WaitGroup

func hello(i int)  {

	defer wg.Done()  //goroutine结束就登记-1
	fmt.Println("hello Goroutine!", i)

}

func main() {

	for i := 0; i <8; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}

	wg.Wait()  //等所有的goroutine都结束
}


 */


// 如果主协程退出了，其他协程也不会再运行
func main() {
	go func() {
		i :=0
		for  {
			i ++
			fmt.Printf("new goroutine: i = %d\n",i)
		}
	}()
	i :=0
	for  {

		i ++
		fmt.Printf("new goroutine: i = %d\n",i)
		if i >2 {
			break
		}

	}
}




