package main

import (
	"fmt"
	"strings"
)

// map映射
func main() {
	// 光声明map类型，但是没有初始化，a就是初始值nil
	// var a map[string]int
	// fmt.Println(a == nil)
	// // map的初始化
	// a = make(map[string]int, 8)
	// fmt.Println(a == nil)
	// // map中添加键值对
	// a["鸿蒙"] = 9000
	// a["明镜"] = 100
	// a["原神启动"] = 200

	// fmt.Println(a)
	// fmt.Printf("type:%T\n", a)

	// // 声明map的同时完成初始化
	// b := map[int]bool{
	// 	1: true,
	// 	2: false,
	// }
	// fmt.Printf("b:%#v\n", b)
	// fmt.Printf("type:%T\n", b)

	// var c map{int}int
	// c[100] = 200//c 这个map没有初始化不能直接操作
	// fmt.Println(c)

	// // 判断某个键存不存在
	// var scoreMap = make(map[string]int, 8)
	// scoreMap["明镜"] = 100
	// scoreMap["修卡拉咪"] = 200
	// // 判断 修卡拉咪 在不在 scoreMap中
	// value, ok := scoreMap["修卡拉咪"]
	// if ok {
	// 	fmt.Println("修卡拉咪在scoreMap中", value)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	// // for range
	// // map是无序的，键值对和添加的顺序无关
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	// // 只遍历map中的key
	// for k := range scoreMap {
	// 	fmt.Println(k)
	// }
	// // 只遍历map中的value
	// for _, v := range scoreMap {
	// 	fmt.Println(v)
	// }
	// // 删除修卡拉咪 这个键值对
	// delete(scoreMap, "修卡拉咪")
	// fmt.Println(scoreMap)

	// 按照某个固定顺序遍历map
	// var scoreMap = make(map[string]int, 100)

	// // 添加66个键值对
	// for i := 0; i < 66; i++ {
	// 	key := fmt.Sprintf("stu%02d", i)
	// 	value := rand.Intn(100) //0~99的随机整数
	// 	scoreMap[key] = value
	// }
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	// // 按照key从大到小的顺序去遍历scoreMap
	// // 1.先取出所有的key存放到切片中
	// keys := make([]string,0,100)
	// for k :=range scoreMap{
	// 	keys = append(keys, k)
	// }
	// // 2.对key做排序
	// sort.Strings(keys)//keys目前是一个有序的切片
	// // 3.按照排序后的key对
	// for _,key :=range keys{
	// 	fmt.Println(key,scoreMap[key])
	// }

	//元素类型为map的切片
	// 	var mapSlice = make([]map [string]int,8,8)//只是完成了切片的初始化
	// 	// [nil nil nil nil nil nil nil nil]
	// 	fmt.Println(mapSlice[0] ==nil)
	// 	// 还需要完成内部map的初始化
	// mapSlice[0] = make (map[string]int,8)//完成了map的初始化
	// mapSlice[0]["明镜"] = 100
	// fmt.Println(mapSlice)

	// // 值为切片的map
	// var sliceMap = make (map [string][]int, 8)//只是完成了map的初始化
	// v, ok  :=sliceMap["中国"]
	// if ok{
	// 	fmt.Println(v)
	// }else {
	// 	// scoreMap中没有中国这个键
	// 	sliceMap["中国"]  = make([]int,8,8)//完成了对切片的初始化
	// 	sliceMap["中国"][0]=100
	// 	sliceMap["中国"][2]=200
	// 	sliceMap["中国"][3]=300
	// }
	// // 遍历sliceMap
	// for k,v :=range sliceMap{
	// 	fmt.Println(k,v)
	// }

	// 统计一个字符串中一个单词出现的次数
	// "how do you do"中每个单词出现的次数
	// 1.定义一个map[String]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 2.字符串中都有哪些单词
	wodes := strings.Split(s, " ")
	// 3.遍历单词做统计
	for _, wode := range wodes {
		v, ok := wordCount[wode]
		if ok {
			// map 中有这个单词的统计结记录
			wordCount[wode] = v + 1
		} else {
			// map 中没有这个单词的统计记录
			wordCount[wode] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}
