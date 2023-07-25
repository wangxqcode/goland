package main

// background: go语言中，每一个请求都是一个goroutine,例如用户登录，需要身份信息认证，验证token，二维码验证，
//当任何一个请求超时，都需要取消goroutine,释放资源
//
/*
var wg sync.WaitGroup

func worker() {

	for {
		fmt.Println("worker")
		time.Sleep(time.Second)

		wg.Done()
	}
}

func main() {
	wg.Add(3)
	go worker()
	wg.Wait()
	fmt.Println("over")
}


 */

/*
// 使用通道方式

var wg sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func worker(exitChan chan struct{})  {
Loop:
	for  {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <- exitChan:
			break Loop
		default:

		}
		wg.Done()
	}

}

func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exitChan <- struct{}{}      // 给子goroutine发送退出信号
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}

 */

/*
// 官方给的方案

var wg sync.WaitGroup


func sleep(ctx context.Context)  {
LOOP:
	for  {
		fmt.Println("sleep")
		select {
		case <-ctx.Done(): // 接受上级指令
			break LOOP
		default:

		}
		wg.Done()

	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go sleep(ctx)
	time.Sleep(6)
	cancel()
	wg.Wait()

}

 */

/*
func main() {
	d := time.Now().Add(100 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

 */




/*
// context.WithTimeout

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

 */




