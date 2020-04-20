package main

import (
	"fmt"
	"sync"
	"time"
)

//func (rw *RWMutex) RLock() 读锁，当有写锁时，无法加载读锁，
//当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景

var (
	x    int
	lock sync.RWMutex
	wg   sync.WaitGroup
)

func read() {
	lock.RLock()
	fmt.Println(x)
	time.Sleep(10 * time.Millisecond)
	lock.RUnlock()
	wg.Done()

}
func write() {
	lock.Lock()
	x = x + 1
	time.Sleep(20 * time.Millisecond)
	lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start))

}
