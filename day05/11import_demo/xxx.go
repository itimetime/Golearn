package main

import (
	"fmt"
	_ "golearnProject/day05/10calc" //匿名导入包 比如我想要包的初始化 只执行init方法
	calc "golearnProject/day05/10calc"
)

//全局声明--> init() --> main()
//包中的标识符（变量名、函数名、结构体、接口等）如果首字母小写的，表示私有（只能子啊当前这个包中使用）
//首字母大写的标识符可以被外部调用
//如果main函数导入a包，a包导入b包，b包导入c包，最后会先执行c包中的ini(),最后执行main包

func main()  {
    sum := calc.Add(2,4)
    fmt.Println(sum)
}