package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "D:\\go\\src\\golearnProject\\day01"
	fmt.Printf(path)

	//多行的字符串
	s := `
少年不识愁滋味，
爱上层楼
`
    fmt.Printf(s)
    s1 := "This is "
    s2 := "Thu"
    ss := s1 + s2
    fmt.Println(s1+s2)
    fmt.Println(ss)
    s3 := fmt.Sprintf("%s%s",s1,s2)
    fmt.Println(s3)

    //分割
    ret := strings.Split(path, "\\")
    fmt.Println(ret)

    //包含
    fmt.Println(strings.Contains(ss,"Thu"))
    fmt.Println(strings.Contains(ss,"Thud"))

    //前后缀
    fmt.Println(strings.HasPrefix(s3,"Th"))
    fmt.Println(strings.HasSuffix(s3,"nu"))

    s4 := "abcdeb"
    fmt.Println(strings.Index(s4,"bc"))
    fmt.Println(strings.LastIndex(s4,"b"))

    //拼接
    fmt.Println(strings.Join(ret, "+"))



}
