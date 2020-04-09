package main

import "fmt"

//类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("assert error!")
	} else {
		fmt.Printf("传进来的字符串%s\n", str)
	}
}

//升级进化版
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Printf("是一个字符串:%s\n", t)
	case int:
		fmt.Printf("是一个int:%d\n", t)
	case int64:
		fmt.Printf("是一个int64:%d\n", t)
	case bool:
		fmt.Printf("是一个bool\n", t)

	}

}

func main() {
	assign("fff")
	i := 8
	i2 := int64(8)
	assign(i)
	assign2(i)
	assign2(i2)
}
