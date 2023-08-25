package main

import (
	"fmt"
	"sync"

)

// goroutine  demo
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 坤坤",i)
	wg.Done()//通知wg把计数器-1
}

func main() { //开启一个主goroutine去执行main函数

	wg.Add(888)//计数牌+1
	for i :=0;i<888;i++{
		// wg.Add(1)//也可
	go hello(i) //开起来一个goroutine去执行这个函数
	}

	fmt.Println("hello main")
	// time.Sleep(time.Second)
	wg.Wait()//阻塞，等所有的小弟都干完活之后才收兵
}
