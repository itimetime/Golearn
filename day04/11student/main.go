package main

import (
	"fmt"
	"os"
)

/*
学生管理系统 函数版
写一个系统，能够查看、新增学生、删除学生
*/

type student struct {
	id   int64
	name string
}

var allStudents map[int64]*student //变量声明

func showAllStudents() {
	for _, v := range allStudents {
		fmt.Printf("学号：%d，姓名: %s\n", v.id, v.name)
	}
}

func addStudents() {
	var(
		name string
		id int64
	)
	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&name)
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&id)
	allStudents[id] = &student{id,name}
	fmt.Printf("已成功添加学生：%s ,%d\n",name,id)

}

func deleteStudent() {
	var id int64
	fmt.Println("请输入你要输出的学号：")
	fmt.Scanln(&id)
	for k, _ := range allStudents{
		if k == id{
			delete(allStudents,id)
			fmt.Printf("已成功删除学生：%d\n",id)
			return
		}
	}
	fmt.Println("输入无效！")


}

func main() {
	allStudents = make(map[int64]*student, 60) //开辟内存空间
	for {
		//1.打印菜单
		fmt.Println("欢迎进入学生管理系统")
		fmt.Println(`
	    1.查看所有学生
        2.新增学生
        3.删除学生
        0.推出
                   `)
		//2.用户选择
		fmt.Println("请选择你要选择的功能：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//3.执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudents()
		case 3:
			deleteStudent()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}

	}

}
