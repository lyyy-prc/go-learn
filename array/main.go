package main

import "fmt"

// 数组
func main()  {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	//数组的初始化
	// 1.定时使用初始化列表的方式初始化
	var cityAeeay = [4]string{"江西","广东","湖南","福建"}
	fmt.Println(cityAeeay)
	// 2.编译器推导数量的长度
	var bollArray = [...]bool{true,false,false,true,false}
	fmt.Println(bollArray)

	// 数组的遍历
	// 1.for循环遍历
	for i := 0; i <len(cityAeeay) ; i++ {
fmt.Println(cityAeeay[i])
	}
// 2.for range 遍历
for index,value := range cityAeeay{
	fmt.Println(index,value)
}





}
