package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//goroutine
func hello() {
	time.Sleep(3)
	fmt.Println("Hello!")
	//defer wg.Done()
}

//程序启动之后创建 主 goroutine
func main() {
	//wg.Add(1)
	go hello()
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		fmt.Println("xiix")
	}()
	fmt.Println("Main function")

	//main 函数结束了 由main函数启动的goroutine 也都结束了

}
