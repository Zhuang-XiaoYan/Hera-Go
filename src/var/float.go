package main

import "fmt"

func main() {
	// go 语言的小数默认都是float64类型
	f1 := 1.2345
	fmt.Println(f1)
	// 显示的声明float32类型
	f2 := float32(1.2345)
	fmt.Println(f2)
	// f1=f2   float32类型的值是不能直接赋值给float64类型的变量的
}
