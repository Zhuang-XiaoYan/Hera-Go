package main

import "fmt"

// 递归函数
func f1() {
	// 递归函数的类型表达式
	f1()
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	fmt.Println("----------------------------------------------")
	var i1 int
	for i1 = 0; i1 < 10; i1++ {
		fmt.Printf("%d  ", fibonacci(i1))
	}
}
