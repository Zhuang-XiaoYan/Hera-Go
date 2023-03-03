package main

import "fmt"

func f10() {
	fmt.Println("hello")
}

func f20() int {
	fmt.Println("hahah")
	return 10
}

func f30(x func() int) {
	// 表示的x 是传进来是一个函数的，同时该函数是一个有返回值的函数
	ret := x()
	fmt.Println(ret)
}

// 函数也是可以成为返回值  表示传递的是一个参数是具有返回值函数类型 同时要求这个函数返回一个具有两个入参同时具有返回值int类型
func f50(x func() int) func(int, int) int {
	fmt.Println("lililili")
	return func(i int, i2 int) int { // 匿名函数
		return x()
	}
}

func main() {
	a := f10
	fmt.Printf("a==%T\n", a)
	b := f20
	fmt.Printf("b==%T\n", b)
	fmt.Printf("%v", b) // 函数的作为的函数就是的地址值

	fmt.Println("----------------------------------------")
	f30(f20)
}
