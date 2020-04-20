package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) //添加随机数种子
	for i := 0; i < 5; i++ {
		r1 := rand.Intn(10)
		r2 := rand.Int()
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

func main() {
	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
