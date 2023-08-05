 package main


 import "fmt"

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
func main()  {
	// 切片的扩容
	var a []int//此时它并没有申请内存
  for i := 0; i < 18; i++ {
		a = append(a,i)
		fmt.Printf("%v len:%d ptr:%p\n ", a,len(a),cap(a),a)

	}
}
