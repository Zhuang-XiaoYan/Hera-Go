package main

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello world")
	}
}

func test() {
	// 这里可以使用defer+recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test 发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "jiale" // error
}

func main() {
	go SayHello()
	go test()
	time.Sleep(time.Second)
}
