package main

import (
	"fmt"
	"os"
)
var smr studentMgr

func showMeau() {
	fmt.Println("Welcome to sms!")
	fmt.Println(`1.查看所有学生
2.添加学生
3.修改学生
4.删除学生
0.退出`)
}

func main()  {
	smr = studentMgr{allStudents: make(map[int64]student, 100)} //修改全局的对象
	for {
		showMeau()
		fmt.Print("请输入要执行的序号")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是",choice)
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.removeStudents()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}

	}
}