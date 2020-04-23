package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println(os.Args)
	name := flag.String("name", "ck","pleas input name")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(flag.Args()) //所有值
	fmt.Println(flag.NArg()) //number
	fmt.Println(flag.NFlag()) // 除name外的值的数量

}
