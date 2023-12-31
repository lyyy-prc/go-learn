package main

import "fmt"

// 指针

/*
func main() {
	a := 10//int类型
	b := &a  //*int类型
	fmt.Printf("%v %p\n", a, &a)//10 0xc0000140c0
	fmt.Printf("%v %p\n", b, b)//0xc0000140c0 0xc0000140c0
	fmt.Println(&b)
	//变量b是一个int类型的指针(*int)，它存储的是变量a的内存地址
	c := *b//根据内存地址去取值
	fmt.Println(c)//10
}
*/

func modifyl(x int) {
	x = 100
}

func modifyl2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modifyl(a)
	fmt.Println(a) //1
	modifyl2(&a)
	fmt.Println(a)
}
