// fib.go
package f

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	if n < 2{
		return n
	}
	n1 := 0
	n2 := 1
	var tmp int
	for ; n>=2;n--{
		tmp = n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}
