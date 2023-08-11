package main

import "fmt"

// 结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	// (*p2).name = "明镜"
	// (*p2).city = "江西"
	// (*p2).age = 18
	p2.name = "明镜"
	p2.city = "江西"
	p2.age = 18
	fmt.Printf("%#v\n", p2)

	// 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
	p3.name = "明镜"
	p3.city = "江西"
	p3.age = 18
	fmt.Printf("%#v\n", p3)

}
