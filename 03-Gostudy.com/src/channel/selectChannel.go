package main

import (
	"fmt"
	"time"
)

func main() {
	stringChan := make(chan string, 5)
	intChan := make(chan int, 5)

	for i := 0; i < 5; i++ {
		intChan <- i
	}
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("hhaha") + string(i)
	}

	// 使用多路复用的读取多个channel的数据  使用select 不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("inchan 的数据是%v\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("stringchan 的数据是%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数组完成所有的数据")
			return
		}
	}
}
