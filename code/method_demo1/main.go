package main

import "fmt"

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewpPerson 是一个Person类型的构造函数
func NewpPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Deram 是为Person类型定义方法
func (p Person) Deram() {
	fmt.Printf("%s的梦想是学好Go语言 \n", p.name)

}

// SetAge 是一个修改年龄的方法
// 指针接受者指的就是接受者的类型是指针
func (p *Person) SetAge(newAge int8) {
	p.age = newAge

}

// SetAge2 是一个使用值接受者来修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge

}

func main() {
	p1 := NewpPerson("明镜", int8(18))
	// (*p1).Deram()
	// p1.Deram()

	// 调用修改年龄的方法
	fmt.Println(p1.age)
	p1.SetAge(int8(88))
	fmt.Println(p1.age)
	p1.SetAge2(int8(99)) //88
	fmt.Println(p1.age)

}
