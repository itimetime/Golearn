package main

import "fmt"

func main()  {
	fmt.Println("ck")
	// 格式化字符串fmt.Printf()
    //%T查看类型
    //%d 十进制
    //%b 二进制
    //%o 八进制
    //%x 十六进制 使用a-f
    //%X 十六进制 使用 A-F
    //%c 字符
    //%s 字符串
    //%p 指针
    //%v 值
    //%f 浮点型
    //%#v
    //%% %号
    //%t bool值
    //%q
    //%U uncode模式



    var m1 = make(map[string]int,1)
    m1["ck"] = 100
    fmt.Printf("%v",m1)
    fmt.Printf("%#v",m1)
	fmt.Printf("%q",65)


    //%b 无小数部分，二进制指数的科学计数法
    //%e 科学计数法
    //%E 科学计数法
    //%f 有小数部分的但无指数部分
    //g 根据实际情况采用%e 或者 %f 格式
    //G 根据实际情况采用%E 或者 %E 格式


    //%s 直接输出字符串或者[]byte
    //%q 输出带双引号的字符串
    //%x 每个字节用十六进制表示
    //%X 每个字节用十六进制表示

    // %9f 宽度为9 默认精度
    // %.2f 默认宽度 精度为2
    // 9.2f 宽度9 精s度2
    //9.f 宽度9 精度0
	s:= "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)

	var ck string
	fmt.Scan(&ck)
	fmt.Println(ck)

	var(
		name string
		age int
		class string
	)
	fmt.Scanf("%s %d %s",&name,&age,&class)

	fmt.Println(name,age,class)

}
