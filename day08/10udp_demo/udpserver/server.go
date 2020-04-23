package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{net.IPv4(127, 0, 0, 1), 40000, ""})
	if err != nil {
		fmt.Println("upd server start failed")
		return
	}
	defer conn.Close()
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed,err", err)
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr)
	}
}
