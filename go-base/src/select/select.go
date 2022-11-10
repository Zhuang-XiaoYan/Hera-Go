package main

import "fmt"

// Select 的使用方式类似于之前学到的 switch 语句，它也有一系列 case 分支和一个默认的分支。
// 每个 case 分支会对应一个通道的通信（接收或发送）过程。
func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
