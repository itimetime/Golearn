package main

import "fmt"

//学生管理系统 结构体版
type student struct {
	name string
	id   int64
}

//造一个学生的管理者

type studentMgr struct {
	allStudents map[int64]student
}

func (s studentMgr) showStudents() {
	for k, v := range s.allStudents {
		fmt.Printf("学号：%d，姓名：%s", k, v.name)
	}
}

func (s studentMgr) addStudents() {
	var (
		name string
		id   int64
	)
	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&name)
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&id)
	s.allStudents[id] = student{name, id}
}

func (s studentMgr) editStudents() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入你要修改的学号：")
	fmt.Scanln(&id)
	//可以不用for循环
	/*第一种
	value := s.allStudents[id]
	if value == nil{
	}else{
	}
	*第二种
	value , ok:= s.allStudents[id]
	if ok{
	}else{
	}
	 */

	for k, _ := range s.allStudents {
		if k == id {
			fmt.Println("请输入你修改后的姓名：")
			fmt.Scanln(&name)
			s.allStudents[id] = student{name, id}
			//value:= s.allStudents[id] //另一种方法
			//value.name = name
			fmt.Printf("已成功修改学生：%d\n", id)
		}
	}
	fmt.Println("输入无效！")
}

func (s studentMgr) removeStudents() {
	var (
		id   int64
	)
	fmt.Println("请输入你要修改的学号：")
	fmt.Scanln(&id)
	for k, _ := range s.allStudents {
		if k == id {
			fmt.Printf("已成功删除学生：%d\n", id)
			delete(s.allStudents, id)
		}
	}
	fmt.Println("输入无效！")
}
