package main

import (
	"fmt"
	"net"
)

//tcp.server

func processConn(conn net.Conn) {
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("Read failed!:", err)
			return
		}

		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	//1.本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Start tcp server failed!:", err)
		return
	}
	//2.等待别人跟我建立连接

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed!:", err)
			return
		}
		//3.与客户端通信
		go processConn(conn)

	}

}
