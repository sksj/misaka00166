package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端

func processConn(conn net.Conn) {
	// 3. 与客户端通信
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please repeat:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:", err)
		return
	}
	for {
		// 2. 等待别人来跟我建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}
}
