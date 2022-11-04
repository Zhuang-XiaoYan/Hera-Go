package main

import "fmt"

// 函数就是一段代码的封装 这样可以重用
func sum(x int, y int) (ret int) {
	return x + y
}

func sum1(x int, y int) {
	fmt.Println(x + y + 10)
	fmt.Println("没有返回值")
}

// 没有参数的 但是有返回值  返回值的可与没有名
func sum2() int {
	return 0
}

func f3() int {
	res := 2
	return res
}

func f4(x int, y int) (res int) {
	res = x * y
	return // 使用的命名的返回值可以省略
}

func f5() (int, int) {
	return 10, 8
}

func f6(x, y, z int) (a, b int) {
	return x + y, y * z
}

// 当参数中的连续多个参数的类型的一致性的 我可以将非最后一个的参数类型的省略
func f7(x, y, z int, s1, s2 string) (a, b int) {
	return x + y, y * z
}

// 可边长的参数必须放在函数的最后面
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y类型的切片
}

func main() {
	// 函数的定义
	fmt.Println(sum(1, 2))
	sum1(1, 4)
	fmt.Println(sum2())
	_, x := f5()
	fmt.Println(x)
	x1, x2 := f6(3, 5, 7)
	fmt.Println(x1, x2)

	f8("xjl", 6, 7)
	f8("xjl")
}
