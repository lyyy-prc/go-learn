package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与JSON序列号

//如果一个Go语言包中定义的标识符是首字母大写的,那么就是对外可见的。
// 如果结构体中的一个字段名首字母是大写的，那么该字符就是对外可见的

type student struct {
	ID   int
	Name string
}

//student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string    `json:"title"` //(:)与("")之间不能有空格//多个键值对之间用空格分隔
	Students []student `json:"student_list" db:"student"`
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "土匪班",
		Students: make([]student, 0, 20),
	}
	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		//造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
	// JAON序列化：Go语言中的数据 -> JSON格式的字符串
	dota, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed", err)
		return
	}
	// fmt.Printf("%T\n", dota)
	fmt.Printf("%s\n", dota)

	// JSON反序列化:JSON格式的字符串 -> JAON序列化：Go语言中的数据
	// jsonStr := `{"Title":"土匪班","Students":[{"ID":0,"Name":"土匪头子"},{"ID":1,"Name":"二当家"}]}`
	// var c2 class
	// err = json.Unmarshal([]byte(jsonStr), &c2)
	// if err != nil {
	// 	fmt.Println("json unmarsha failed err", err)
	// 	return
	// }
	// fmt.Printf("%#v\n", c2)

}
