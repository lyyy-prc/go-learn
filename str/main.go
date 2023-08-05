package main

import (
	"fmt"
	"strings"
)

//字符串常见操作
func main()  {
	//求字符串的长度
	s := "hello"
	fmt.Println(len(s))
	s2 :="明镜"
	fmt.Println(len(s2))

//拼接字符串
fmt.Println(s+s2)
s3 :=fmt.Sprintf("%s - %s",s,s2 )
fmt.Println(s3)

//字符串的分割
s4 :="床 前 明 月 光"
fmt.Println(strings.Split(s4," "))
fmt.Println("%T\n",strings.Split(s4," "))

//判断是否包含
fmt.Println(strings.Contains(s4,"明"))

//判断前缀
fmt.Println(strings.HasPrefix(s4, "床"))
//判断后缀
fmt.Println(strings.HasSuffix(s4, "光"))

//判断子串的位置
fmt.Println(strings.Index(s4, "明"))
// 判断子串最后出现的位置
fmt.Println(strings.LastIndex(s4, "明"))

//join
s5 := []string{"床","前","明","月","光"}
fmt.Println(s5)
fmt.Print(strings.Join(s5, "+"))

}
