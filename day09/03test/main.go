package main

import (
	"fmt"
	"golearnProject/day09/03test/split"
	"strings"
)

func main()  {
	str := "fdfsdf"
	//n := strings.Index(str,"fs5")
	//fmt.Println(n)
	//for i,r := range []rune(str){
	//	fmt.Println(i,r)
	//}
	newStr:=split.Split(str,"")
	fmt.Println(newStr)
	fmt.Printf("%#v",newStr)
	test := strings.Split(str,"")
	fmt.Printf("%#v",test)

}


