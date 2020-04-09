package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 := "how do you do"
	ret := strings.Split(s1, " ")
	m1 := make(map[string]int)

	for _ , v := range ret{
		m1[v] += 1
		//视频给的答案，没我的精简
		//if _,ok := m1[v]; !ok{
		//	m1[v] = 1
		//}else{
		//	m1[v]++
		//}
	}

	//教程给的答案

	fmt.Println(m1)
}
