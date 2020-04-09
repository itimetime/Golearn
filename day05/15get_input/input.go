package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入的时候如果有空格

func useScan(){
	var s string
	fmt.Print("请输入内容")
	fmt.Scanln(&s)
	fmt.Println(s)
}

func useBufio()  {
	//可以获取空格
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}

func main() {
	//useScan()
	//useBufio()
	/*
	logger.Trace()
	logger.Debug()
	logger.Warning()
	logger.Info()
	logger.Error()
	 */
	//写日志
	fmt.Fprint(os.Stdout, "这是一条技术日志")
	fileObj, _ := os.OpenFile("./xxx.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprint(fileObj, "这是一条技术日志")
	fileObj.Close()
}

