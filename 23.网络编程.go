package main

import (
	"fmt"
	"net"
)

// 这是一个服务端
func main() {
	// 创建tcp的监听地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	// tcp监听
	listen, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		//等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		//获取到客户端的地址
		fmt.Println(conn.RemoteAddr().String() + " 进来了")
		// 读取客户端传来的数据
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(conn.RemoteAddr().String() + " 出去了")
				break
			}
			fmt.Println(string(buf[0:n]))

		}
	}

}
