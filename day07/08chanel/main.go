package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//chanel练习
/*
1.启动一个goroutine，生成100个数发送到ch1
2.启动一个goroutine，从ch1中取值，计算其平方到ch2中
3.在main中 从ch2 取值打印出来
*/

var ch1, ch2 chan int
var wg sync.WaitGroup

func main() {
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- rand.Intn(10)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			num := <-ch1 //可返回两个值，一个ok
			ch2 <- num * num
		}
		wg.Done()
	}()
	for i := 0; i < 100; i++ {
		res := <-ch2
		fmt.Println(i, res)
	}
	wg.Wait()
	close(ch1)
	close(ch2)
}
