package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err :=os.OpenFile("xx.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("open err,err:%s",err)
	}
	//默认输出终端
	//log.SetOutput(os.Stdout)
	log.SetOutput(fileObj)
	for{
		time.Sleep(time.Second*3)
		log.Println("测试的语句")
	}
}
