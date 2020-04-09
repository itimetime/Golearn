package main

import "fmt"

//append()为切片追加元素
//会自动初始化切边

func main()  {
	//a1 := [3]string{"上海","深圳"}
	s1 := []string{"北京","上海","深圳"}
	s2 := append(s1, "广州")
	//append追加元素，原来的底层数组放不下的情况下，会继续申请一块新的地址
	//a2 := append(a1, "广州") 实验证明 数组不可追加
	fmt.Println(s1,len(s1),cap(s1))
	fmt.Println(s2,len(s2),cap(s2))


	a1 := [5]string{"北京","上海","深圳"}
	s3 := a1[0:3]
	s4 := append(s3, "广州")
	fmt.Println(s4,len(s4),cap(s4))
	fmt.Println(a1,len(a1),cap(a1))
	ss := []string{"武汉","西安","杭州"}
	s5 := append(s4,ss...) //表示分开
	fmt.Println(s5,len(s5),cap(s5))



}
