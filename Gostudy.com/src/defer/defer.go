package main

import "fmt"

func deferdemo() {
	fmt.Println("start print")
	fmt.Println("hahaha")
	fmt.Println("end print")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("start print")
	defer fmt.Println("hahaha") // defer 把这个的语句延迟到函数的即将返回的时候在执行
	fmt.Println("end print")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("start print")
	defer fmt.Println("001") // Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，
	defer fmt.Println("002") // 将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	defer fmt.Println("003") // defer 多用于函数之前释放资源
	fmt.Println("end print")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 要求返回的x的值
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 全局作用域
var a = 10

func main() {
	// defer
	// deferdemo()
	// go语言中的的函数return 不是原子的操作 而是在底层分两步来执行的操作的
	// 第一步;返回函数值
	// 第二步：真正的ret返回
	// 如果函数存在defer 那么执行defer的实际就是在第一步和第二步之间
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	fmt.Println("--------------------------------")
	// 函数的作用域的是的现在内部变量中查找
	// 找不到的就往函数的外部查找 一直到全局查找
	fmt.Println(a)

	// 语句快的作用域
	if i := 10; i < 18 {
		fmt.Println(i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

}
