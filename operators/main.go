package main

import "fmt"

//GO 语言中的运算符
func main() {
	// 1.算数运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(87 % 4) //求余
	a++                 //自增  a--(自减)
	fmt.Println(a)

	// 2. 关系运算符
	fmt.Println(10 > 288)
	fmt.Println(10 != 4)

	//3.逻辑运算符
	fmt.Println(10 > 8 && 4 == 4)
	fmt.Println(!(10 < 9))
	fmt.Println(1 > 5 || 9 > 2)

	// 4.位运算符
	a = 1 //001
	b = 5 //101
	fmt.Println(a & b)

}
