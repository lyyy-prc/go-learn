package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 1; i < 100; i++ {
		fmt.Println("A", i)
	}
	wg.Done()
}

func b() {
	for i := 0; i < 100; i++ {
		fmt.Println("B", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(8) //只占用一个CPU核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

// windows平台上效果不显示
