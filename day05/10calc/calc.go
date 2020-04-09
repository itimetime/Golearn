package calc

import "fmt"

func Add(x,y int) int{
	return x+y
}

func init()  {
	fmt.Println("导入我的时候自动执行")
}