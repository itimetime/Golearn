package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作

var x int64
var wg sync.WaitGroup

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	//比较并交换
	ok := atomic.CompareAndSwapInt64(&x, 100000, 666)
	fmt.Println(x, ok)

}
