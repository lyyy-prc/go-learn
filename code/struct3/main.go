package main

import "fmt"

//嵌套结构体的字段冲突

type Adderss struct {
	province    string
	city        string
	uppdateTime string
}

type Email struct {
	addr        string
	uppdateTime string
}

type Persond struct {
	name    string
	gender  string
	age     int8
	Adderss //嵌套另外一个结构体
	Email   //嵌套另一个结构体
}

func main() {
	P1 := Persond{
		name:   "明镜",
		gender: "男",
		age:    18,
		Adderss: Adderss{
			province:    "江西",
			city:        "赣州",
			uppdateTime: "2018-09-01",
		},
		Email: Email{
			addr:        "@qq.com",
			uppdateTime: "2016-09-01",
		},
	}
	fmt.Printf("%#v\n", P1)
	// fmt.Println(P1.uppdateTime)//嵌套结构体中包含多个同名的updateTime字段
	fmt.Println(P1.Email.uppdateTime)
	fmt.Println(P1.Adderss.uppdateTime)
}
