package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func a() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A %d \n", i)
	}
}

func b() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B %d \n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(6)         // 表示的6核处理器 默认是的物理的CPU的核心数
	fmt.Println(runtime.NumCPU()) // 12 个线程数
	wg1.Add(2)
	go a()
	go b()
	wg1.Wait()
}
