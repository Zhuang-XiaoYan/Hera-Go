package main

import "fmt"

func fn1() {
	fmt.Println("fn1")
}

func fn2() {
	// panic 相当于是python 中的关键字raise。 用于终止程序
	panic("抛出一个异常错误")
}

func fn3() {
	defer func() {
		// 用于捕获panic的错误请求
		err := recover()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}()
	panic("抛出一个异常错误")
}

func main() {
	fn1()
	fn3()
	fmt.Println("结束main函数")
}
