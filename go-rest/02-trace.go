package main

import (
	"context"
	"fmt"
	"time"

	"io/ioutil"
	"log"
	"os"
	"runtime/trace"
	"sync"

	utiltrace "k8s.io/utils/trace"

)

/*
func main()  {
	client := resty.New()


	// Trace

	resp, err :=client.R().EnableTrace().Get("https://activity.lenovo.com.cn/")
	if err != nil {
		fmt.Println(err)
	}

	ti := resp.Request.TraceInfo()
	fmt.Println("Request Trace Info:")
	fmt.Println("DNSLookup:", ti.DNSLookup)
	fmt.Println("ConnTime:", ti.ConnTime)
	fmt.Println("TCPConnTime:", ti.TCPConnTime)
	fmt.Println("TLSHandshake:", ti.TLSHandshake)
	fmt.Println("ServerTime:", ti.ServerTime)
	fmt.Println("ResponseTime:", ti.ResponseTime)
	fmt.Println("TotalTime:", ti.TotalTime)
	fmt.Println("IsConnReused:", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime:", ti.ConnIdleTime)
	fmt.Println("RequestAttempt:", ti.RequestAttempt)
	fmt.Println("RemoteAddr:", ti.RemoteAddr.String())



}

 */


func main() {

	// 创建trace持久化的文件句柄

	f, err := os.Create("trace.out")
	if err != nil {
			log.Fatalf("failed to create trace output file: %v", err)
	}

	// trace 绑定文件句柄
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	ctx, task := trace.NewTask(context.Background(),"main start")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		r := trace.StartRegion(ctx, "reading file")
		defer r.End()
		ioutil.ReadFile(`n1.txt`)
	}()

	go func() {
		defer wg.Done()
		r := trace.StartRegion(ctx, "writing file")
		defer r.End()

		ioutil.WriteFile(`n2.txt`, []byte(`42`), 0644)
	}()

	wg.Wait()
	defer task.End()
	Trace()





}

const (
	LogIfLongTime = time.Second * 1
)




func Trace( ) {
	trace := utiltrace.New("check if model has changed")
	trace.Step("check if model has changed")
	fmt.Println("trace-----------------",trace)
	defer trace.LogIfLong(LogIfLongTime)
}


