package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp/server/main.go

// TCP server端





// 处理函数

func process(coon net.Conn)  {
	defer coon.Close()
	for  {
		reader := bufio.NewReader(coon)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from failed, err:",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		coon.Write([]byte(recvStr)) // 发送数据

	}
}

func main() {
//	监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Println("listen failed err ", err)
		return
	}

//	建立连接
	for  {
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("addept failed err", err)
		}
		go process(coon)
	}

}

