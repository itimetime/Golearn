package main

import (
	"fmt"
	"strconv"
	"sync"
)

//
//import (
//	"fmt"
//	"strconv"
//	"sync"
//)
//
//var m = make(map[string]int)
//var lock sync.Mutex
//
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string, value int) {
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 200; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			lock.Lock()
//			set(key, n)
//			lock.Unlock()
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
