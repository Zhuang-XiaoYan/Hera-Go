package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func hello2() {
	fmt.Println("hello")
	wg.Done() // 告知当前goroutine完成
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1) // 登记1个goroutine
		go hello2()
	}
	fmt.Println("你好")
	wg.Wait() // 阻塞等待登记的goroutine完成
}
