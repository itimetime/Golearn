package main

import (
	"fmt"
	"strings"
)

func main(){
	//视频中runes切片，进行的判断
	s := "山西运煤车煤运西山"
	res := strings.Split(s,"")
	l := len(res)
	for i := 0; i <l-i-1 ;i++{
		if res[i] != res[l-1-i]{
			fmt.Println("不是回文数")
			return
		}
	}
	fmt.Println("是回文数")
}
