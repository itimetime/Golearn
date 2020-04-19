package main

import (
	"fmt"
	"strconv"
)

//strconv
//类型进行转换

func main() {

	i := int32(2361)
	ret := string(i)
	fmt.Println(ret)
	ret2 := fmt.Sprintf("%d", i) //int 转换str
	fmt.Println(ret2)

	str := "10000"
	ret3, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("error!")
	}
	fmt.Println(ret3)

	ret4 := int32(ret3)
	fmt.Println(ret4)
	//str to int
	ret5, _ := strconv.Atoi(str)
	fmt.Println(ret5)

	//解析bool类型
	boolStr := "true"
	ret6, _ := strconv.ParseBool(boolStr)
	fmt.Println(ret6)

	//解析浮点字符串
	floatStr := "1.234"
	ret7, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println(ret7)
}
