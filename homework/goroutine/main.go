package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
开启一个goroutine循环生成int64类型的随机数，发送到jobs
开启24个goroutine从jobs中取出随机数计算各位数的和，将结果发送到results
主goroutine从results取出结果并打印到终端输出
*/

func worker(id int, jobs <-chan int64, results chan<- int64) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- numSum(j)
	}
}

func creatJob(jobs chan<- int64) {
	for {
		jobs <- rand.Int63()
	}
}

func numSum(j int64) (sum int64) {
	for j > 0 {
		a := j % 10
		j = j / 10
		sum += a
	}
	return
}

func main() {
	jobs := make(chan int64, 100)
	results := make(chan int64, 100)
	// 开启一个goroutine循环生成int64类型的随机数，发送到jobs
	wg.Add(1)
	go creatJob(jobs)

	//开启24个goroutine从jobs中取出随机数计算各位数的和，将结果发送到results
	for w := 0; w <= 24; w++ {
		wg.Add(1)
		go worker(w, jobs, results)
	}

	// 主goroutine从results取出结果并打印到终端输出
	for {
		res := <-results
		fmt.Println(res)
	}
	wg.Wait()
}
