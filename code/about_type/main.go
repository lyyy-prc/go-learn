package main

import "fmt"

// 自定义类型和类型别名示意

// 1.自定义类型

// Myint 基于int类型的自定义
type Myint int

//2.类型别名

// Newint int类型别名
type Newint = int

func main() {
	var i Myint
	fmt.Printf("type:%T value%v\n", i, i)

	var x Newint
	fmt.Printf("type:%T value%v\n", x, x)
}
