package main

import (
	"fmt"
	"reflect"
)

// reflect demo

func reflectType(x interface{}) {
	// 我是不知道别人会调用我这个函数的时候会传进来什么类型的变量
	// 1.方式1:通过类型断言
	// 2.方法2:借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind()) //Go语言中数组，切片，Map，指针等类型的变量，它们的。Name()返回都是空
	fmt.Printf("%T\n", obj)                  //*reflect.rtype
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Kind() //拿到值对应的类型种类
	fmt.Println(k)
	// 如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成一个fmoat32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针去对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(4.321)

	}
}

//Cat 是一个结构体
type Cat struct{}

//Dog 是一个结构体
type Dog struct{}

func main() {
	// var a float32 = 1.234
	// reflectType(a)
	// var b int8 = 10
	// reflectType(b)
	// // 结构体
	// var c Cat
	// reflectType(c)
	// var d Dog
	// reflectType(d)
	// // slice
	// var e []int
	// reflectType(e)
	// var f []string
	// reflectType(f)

	// ValueOf
	// var aa int32 = 88
	// reflectValue(aa)
	// var bb float32 = 1.234
	// reflectValue(bb)

	// Set Value
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}

// type Kind uint
// const (
//     Invalid Kind = iota  // 非法类型
//     Bool                 // 布尔型
//     Int                  // 有符号整型
//     Int8                 // 有符号8位整型
//     Int16                // 有符号16位整型
//     Int32                // 有符号32位整型
//     Int64                // 有符号64位整型
//     Uint                 // 无符号整型
//     Uint8                // 无符号8位整型
//     Uint16               // 无符号16位整型
//     Uint32               // 无符号32位整型
//     Uint64               // 无符号64位整型
//     Uintptr              // 指针
//     Float32              // 单精度浮点数
//     Float64              // 双精度浮点数
//     Complex64            // 64位复数类型
//     Complex128           // 128位复数类型
//     Array                // 数组
//     Chan                 // 通道
//     Func                 // 函数
//     Interface            // 接口
//     Map                  // 映射
//     Ptr                  // 指针
//     Slice                // 切片
//     String               // 字符串
//     Struct               // 结构体
//     UnsafePointer        // 底层指针
// )
