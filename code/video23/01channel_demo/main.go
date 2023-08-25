package main

import "fmt"

func main() {
	// var ch1 chan int// 引用类型，需要初始化之后才能使用
	// ch1 = make(chan int,8)
	// ch1 := make(chan int)    // 无缓冲区通道,又称为同步通道
	ch1 := make(chan int, 8) // 带缓冲区通道
	ch1 <- 10                //发送值
	x := <-ch1
	fmt.Println(x)
	// len(ch1)//取通道中的元素
	// cap(ch1)//拿到通道的容量
	close(ch1)
}
