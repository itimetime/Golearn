package main

import . "fmt"

func main() {
	//对应Python中的字典
	var m1 map[string]int
	Println(m1)
	Println(m1 == nil)
	//初始化的时候要计算和map容量，避免动态扩容
	m1 = make(map[string]int, 10)
	m1["ck"] = 12
	m1["wy"] = 11
	Println(m1)
	Println(m1["ck"])
	v, ok := m1["dd"]
	if !ok {
		Println("key error")
	} else {
		Println(v)
	}

	for key, value := range m1 {
		Println(key, value)
	}

	for key := range m1{
		Println(key)
	}

	for _, value := range m1{
		Println(value)
	}

	//删除key
	delete(m1,"ck")
	Println(m1)
	//删除不存在的key,什么都不干
	delete(m1,"df")
	Println(m1)

}
