package main

import "fmt"

// 结构体的匿名字段

type Person struct {
	string
	int8
}

func main() {
	p1 := Person{
		"明镜",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int8)
}
