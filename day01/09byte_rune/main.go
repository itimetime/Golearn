package main

import "fmt"

func main() {
	//byte类型代表ASCII字符和rune类型，代表UTF-8字符
	s := "We are family!"
	n := len(s)

	for i := 0; i < n; i++ {
		fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])
	}

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s4 := []byte(s2)
	s3[0] = '红'
	s4[0] = 'p'
	fmt.Println(string(s3),string(s4))

	c1 := '红' //rune(int32)
	c2 := "红"
	fmt.Printf("c1: %T, c2: %T\n",c1, c2)
	fmt.Printf("s3[0]: %T, s4[0]: %T\n",s3[0], s4[0])


	//类型转化
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
}
