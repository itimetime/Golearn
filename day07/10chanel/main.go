package main

import (
	"fmt"
	"sync"
	"time"
)

var ch2 chan string
var wg sync.WaitGroup

func f1(ch2 chan string) {
	fmt.Println("运行")
	for {
		i := <-ch2
		fmt.Println("结束", i)
	}
	fmt.Println("结束>????????")

}

func main() {
	//ch1 = make(chan string, 20)
	ch2 = make(chan string, 20)
	ch2 <- "556"
	ch2 <- "556"
	ch2 <- "556"
	wg.Add(1)
	go f1(ch2)
	time.Sleep(1000000)
	wg.Done()
	ch2 <- "556"
	wg.Wait()

}
