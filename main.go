package main

import (
	"fmt"
	"math"
)

//常量
//const pi = 3.14
//const e = 2.7

//const (
//	pi = 3.14
//  e = 2.7
//  )

//const (
//	n1 = 10
//	n2
//	n3
//)

const (
	n1 = iota
	n2 = iota
	n3 = 1000
 	n4 = iota

)

const n5 = iota

func main()  {
	fmt.Println(n1,n2,n3,n4,n5)

// uint8
var age uint8
age = 89
fmt.Println(age)


//浮点数
fmt.Println(math.MaxFloat32)
fmt.Println(math.MaxFloat64)

//布尔制
var a bool
fmt.Println(a)
a = true
fmt.Println(a)

//字符串
s1:="hello beijing"
s2:="你好 北京"
fmt.Println(s1)
fmt.Println(s2)

//打印win平台下的一个路径 c:\Users\TOE\Desktop\lyyy.exe
fmt.Println("c:\\Users\\TOE\\Desktop\\lyyy.exe")
fmt.Println("\t 制表符 \n换行符")

s3 :=`
多行字符串
两个反引号
之间
的内容
会 原样输出
\t
\n\
`
fmt.Println(s3)

}
