package main

import "fmt"

// 函数
// 定义一个不需要参数也没有返回值的函数 sayHello
func sayHello() {
	fmt.Println("Hello 明镜!")
}

// 定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello", name)
}

// 定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
} //另一种方式

//函数接收可变参数,在参数后面加... 表示可变参数
// 可变参数在函数体中是切片类型
func intSum3(a ...int) int {
	// fmt.Println(a)
	// fmt.Printf("%T\n",a)
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

// // Go 语言中参数类型简写
// func intSum2(a, b int) (ret int) {
// 	ret = a + b
//  	return
//  }

// 定义具有多个返回值的函数(多返回值也支持类型简写)
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

//  固定参数和可变参数同时出现时，可变参数要放在最后
//go 语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}

// 函数调用
func main() {
	// sayHello()
	// name := "明镜"
	// sayHello2(name)
	// r :=intSum(123, 654)
	// fmt.Println(r)
	//intSum(123, 654)
	// 	r1 :=intSum3()
	// 	r2 :=intSum3(123)
	// 	r3 :=intSum3(123,654)
	// 	r4 :=intSum3(123,654,333)
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println(r3)
	// fmt.Println(r4)
	// 	r1 :=intSum4(0)//a=0 ,b=[]
	//   r2 :=intSum4(123)//a=123, b=[]
	// 	r3 :=intSum4(123,654)//a=123,b[654]
	// 	r4 :=intSum4(123,654,333)//a=123,b[654,333]
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println(r3)
	// fmt.Println(r4)
	x, y := calc(100, 200)
	fmt.Println(x, y)

}
