package main

import "fmt"

func main() {

	var (
		a = 3
		b = 5
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a, b)
	a += 1
	b += 1
	fmt.Println(a, b)
	//逻辑运算符
	age := 14
	if age >= 10 && age < 15 {
		fmt.Println("10-15")
	}
	//not 取反
	isMarried := false
	fmt.Println(!isMarried)

	//位运算符 针对的二进制数
	//5的二进制101
	//2的二进制10
	//&： 按位与（均为1为1） | 按位或（有1位就为1）

	fmt.Println(5 & 2)
	fmt.Println(5 | 2)

	// ^ 按位异或，不同为1
	fmt.Println(5 ^ 2)

	// << 左移 右边加0  101 => 1010
	fmt.Println(5 << 1)

	// >> 右移 101 => 01
	fmt.Println(5 >> 2)

	var m = int8(1)
	fmt.Println(m<<10) //超出界限


	//赋值运算符
	x := 10

	x <<= 2 //x = x<<2
	x &= 2 //x = x&2
	x |= 2 //x = x|2
	x ^= 2 //x = x^2
	x >>= 2 //x = x>>2



}
