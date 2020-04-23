package main

import (
	"fmt"
	"net"
)

func main()  {
	socket, err := net.DialUDP("udp",nil, &net.UDPAddr{net.IPv4(127, 0, 0, 1), 40000, ""})
	if err != nil {
		fmt.Println("upd client start failed")
		return
	}
	defer socket.Close()
	var reply [1024]byte
	var msg string
	for{
		fmt.Scanln(&msg)
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed, err:",err)
		}
		fmt.Println("收到回复信息：",string(reply[:n]))
	}
}
