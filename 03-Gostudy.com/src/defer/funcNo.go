package main

import "fmt"

// 多次被调用的
var f = func(x, y int) {
	fmt.Println(x + y)
}

func f11(f func()) {
	fmt.Println("this is f11")
	f()
}

func f22(x, y int) {
	fmt.Println("this is f22")
	fmt.Println(x + y)
}

// 闭包函数
func f33(f44 func(int, int)) func() {
	tmp := func() {
		f44(10, 20)
		fmt.Println("********")
	}
	return tmp
}

func lixaing(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func main() {
	// 调用匿名函数
	f(1, 2)
	// 如果是的调用一次的函数 还可以简写立即调用的函数类型
	func(x int) int {
		fmt.Println(x)
		return x
	}(10) // 在函数的后面使用的（）表示立即执行

	fmt.Println("--------------------------")
	// f11(f22) 这样是不符f11的规则
	f11(f33(f22))
	fmt.Println("-----------lili---------------")
	lixaing(f22, 1, 3)

}
