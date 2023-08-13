package main

import "fmt"

// 嵌套结构体

type Adderss struct {
	province string
	city     string
}

type Persond struct {
	name    string
	gender  string
	age     int8
	Adderss //嵌套另外一个结构体

}

func main() {
	p1 := Persond{
		name:   "明镜",
		gender: "男",
		age:    18,
		Adderss: Adderss{
			province: "江西",
			city:     "赣州",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.name, p1.gender, p1.age, p1.Adderss)
	fmt.Println(p1.Adderss.province) //通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(p1.province)
}
