package main

import "fmt"

//空接口
//interface{} 空接口类型

func show(i interface{})  {
	fmt.Printf("%T\n",i)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周琳"
	m1["age"] = 9000
	m1["married"] = true
	m1["hobby"] = [...]string{"唱","跳","Rap","篮球"}
	fmt.Println(m1)

	show(false)
	show(m1)
}
