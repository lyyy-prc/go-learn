package main

import "fmt"

// 结构体的初始化

type person struct {
	name, city string
	age        int8
}

func main() {
	// 1.键值对初始化
	p4 := person{
		name: "明镜",
		city: "江西",
		age:  18,
	}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "明镜",
		city: "江西",
		age:  18,
	}
	fmt.Printf("%#v\n", p5)

	// 2.值的列表进行初始化(1.必须初始化结构体的所有字段。2.初始值的填充顺序必须与字段在结构体中的声明顺序一致 3.该方式不能和键值初始化方式混用)
	p6 := person{
		"明镜",
		"江西",
		18,
	}
	fmt.Printf("%#v\n", p6)
}
