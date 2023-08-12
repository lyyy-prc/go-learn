package main

import "fmt"

// 定义结构体

// type person struct{
// 	name string
// 	age int
// 	city string
// }

type person struct {
	name, city string //类型相同可以简写
	age        int
}

func main() {
	var p person
	p.name = "明镜"
	p.city = "江西"
	p.age = 18
	fmt.Printf("p=%#v", p)
	fmt.Println(p.name)
	fmt.Println(p.city)
	fmt.Println(p.age)

	// 匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "刘柳琉"
	user.married = false
	fmt.Printf("%#v\n", user)
}
