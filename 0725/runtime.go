package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched() 让出CPU时间片，重新分配任务
// runtime.Goexit() 退出当前协程

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine 1 is running")
			runtime.Gosched()
		}
	}()



	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("=----------------------")
			runtime.Gosched()
		}

	}()


	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine 2 is running")
			runtime.Goexit()
		}
	}()

	time.Sleep(100 * time.Millisecond)
}









