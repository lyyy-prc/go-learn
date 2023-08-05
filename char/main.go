package main

import "fmt"

//字符
func main()  {
	// byte uint8的别名 ASCII码
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1,c2)
	fmt.Println("c1:%T c2%T",c1,c2)


	s :="明镜"
	for i := 0; i < len(s); i++ {
		fmt.Println("%c\n",s[1])
		for _,r :=range s{
		fmt.Println("%c\n",r)
	}
} 


}
