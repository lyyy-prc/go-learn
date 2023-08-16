package main

import "fmt"

type student struct {
	id    int //学号是唯一的
	name  string
	class string
}

//newStudent 是 student 类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}

}

// 学员管理类型
type structMgr struct {
	allStudents []*student
}

// newStudentMgr 是 studentMgr的构造函数
func newStudentMgr() *structMgr {
	return &structMgr{
		allStudents: make([]*student, 0, 88),
	}

}

// 1.添加学员生
func (s *structMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2.编辑学生
func (s *structMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { //当学号相同时，就表示找到了要修改的学生
			s.allStudents[i] = newStu //根据切片的索引直接把新学生赋值进来
			return
		}
	}
	// 如果走到这里说明输入的学生没有找到
	fmt.Printf("输入的学生信息有误，系统中没有学号是%d的学生\n", newStu.id)
}

// 3.展示学生
func (s *structMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
