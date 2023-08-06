package main

import (
	"fmt"
	"sort"
)

// 切片声明和定义时不用定义长度，[]里面是空的（切片与数组的区别）

// // 切片（slice）
// func main()  {
// // 	// 基于数组得到切片
// // 	//a := [5]int{12,34,56,78,90}
// // 	b :=a[1:4]
// // 	fmt.Println(b)
// // 	fmt.Printf("%T\n",  b)
// // 	// 切片再次切片
// // 	c :=b[0:len(b)]
// // 	fmt.Println(c)
// // 	fmt.Printf("%T\n", c)

// 	// 切片的扩容
//   var a []int //此时它并没有申请内存
// 	for i := 0; i < 10; i++ {

// 	a =append(a,i)
// 	fmt.Println("%v ,a)
// 	}

// }
func main() {
	// 切片的扩容
	// var a []int //此时它并没有申请内存
	// // for i := 0; i < 18; i++ {
	// // 	a = append(a, i)
	// // 	fmt.Printf("%v len:%d ptr:%p\n ", a, len(a), cap(a), a)

	// // }
	// a = append(a,1,2,3,4,5)
	// fmt.Println(a)
	// b := []int{12,34,56,67}
	// a = append(a, b...)
	// fmt.Println(a)

	//  切片的copy
	// a := []int{1,2,3,4,5}
	// b := make([]int,5,5)
	// c := b
	// copy(b, a)
	// b[1]=10086
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// //切片删除元素
	// a :=[]string{"南昌","赣州","宜春","萍乡"}
	// a =append(a[0:2], a[3:]...)
	// fmt.Println(a)

	// // append(a[:index]，a[index+1:]...)

	// var a = make([]string, 5, 10) //[     ]
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a) //[     123456789]

	var a = [...]int{3, 7, 8, 9, 1}
	// a[:]得到的是一个切片,指向了底层的数组a
	sort.Ints(a[:])
	fmt.Println(a)

}
