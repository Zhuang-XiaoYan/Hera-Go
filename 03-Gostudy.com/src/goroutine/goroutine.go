package main

import (
	"fmt"
	"time"
)

// Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。
// 区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。
// 例如Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。
// Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动也会创建一盒Goroutine去执行
func main() {
	for i := 0; i < 10; i++ {
		// 开启一个Goroutine去执行一个hello的函数任务。
		// go-base hello(i)
		// 如果没有传入i 该函数就是的用的外面的i 有可能导致的相同的输出
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("main")
	time.Sleep(time.Second)

	// main函数的结束后，有main函数启动的Goroutine也都结束了。
}
