package main

import (
	"fmt"
	"net"
)

//tcp client

func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil{
		fmt.Println("建立连接失败，err：",err)
		return
	}
	//2.发送数据

	//通过命令行获取
	//if len(os.Args) < 2{
	//	msg = "Hello socket pogroming"
	//}else {
	//	msg = os.Args[1]
	//}
	var msg string
	for{
		fmt.Scanln(&msg)
		//另一种写法间注释
		conn.Write([]byte(msg))
	}
	conn.Close()

}
//
//reader := bufio.NewReader(os.Stdin)
//text,_ := reader.ReadString('\n')
//text = strings.Trimspace(text)