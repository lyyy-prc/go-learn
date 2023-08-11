package main

import "fmt"

// // 函数进阶之变量作用域

// // 定义全局变量num
// var num int = 10

// //定义函数
// func testGlobal() {
// 	num := 100

// 	//name :="明镜"

// 	// 可以在函数中s使用变量
// 	//1.先在自己函数中查找，找到了就用自己函数中的
// 	//2.函数中找不到变量就往外层找全局变量
// 	fmt.Println("变量num", num)
// }

// func main() {
// 	// testGlobal()
// 	// // 外层不能访问到函数内部变量（局部变量）
// 	// // fmt.Println(name)

// 	// // 变量i此时只在for循环中的语句块中生效
// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println(i)
// 	// }
// 	// //fmt.Println(i)//外层访问不到内部for语句块中的变量

// 	// 函数是可以作为变量
// 	abc := testGlobal
// 	fmt.Printf("%T\n", abc)
// 	abc()
// }

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
