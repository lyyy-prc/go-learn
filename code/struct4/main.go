package main

import "fmt"

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type Dog struct {
	feet    int8
	*Animal //匿名嵌套，而且嵌套的是一个结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会狗叫\n", d.name)
}

func main() {
	d1 := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "小白",
		},
	}
	d1.wang() //狗会叫
	d1.move() //狗会动
}
