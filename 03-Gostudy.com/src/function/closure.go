package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc4(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 使用闭包函数
// 使用闭包 使用的函数之外的变量
// 闭包是指有权房屋内另外一个函数在作用域中的变量的函数
// 创建闭包的常见方式就是的在函数内部创建另外一函数 通过另外一个函数来访问这个函数局部的变量
// 由于闭包的作用域的局部变量资源不会被立刻收回，所有有可能会占用更多的内存资源。使用过度有可能会导致性下降，建议在非常有必要的时候才使用闭包。
func adder5() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder1() func() int {
	var i = 10
	return func() int {
		// 这里没有修改i的值
		return i + 1
	}
}

func addder2() func(y int) int {
	var i = 10
	return func(y int) int {
		// 这里会修改内存i的值类型
		i += y
		return i
	}
}

func main() {
	var f1 = adder()
	fmt.Println(f1(10)) //10
	fmt.Println(f1(20)) //30
	fmt.Println(f1(30)) //60
	fmt.Println("----------------------------------------")
	f2 := adder()
	fmt.Println(f2(40)) //40
	fmt.Println(f2(50)) //90
	fmt.Println("----------------------------------------")
	var f3 = adder2(10)
	fmt.Println(f3(10)) //20
	fmt.Println(f3(20)) //40
	fmt.Println(f3(30)) //70
	fmt.Println("----------------------------------------")
	f4 := adder2(20)
	fmt.Println(f4(40)) //60
	fmt.Println(f4(50)) //110
	fmt.Println("----------------------------------------")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
	fmt.Println("----------------------------------------")
	f16, f17 := calc4(10)
	fmt.Println(f16(1), f17(2)) //11 9
	fmt.Println(f16(3), f17(4)) //12 8
	fmt.Println(f16(5), f17(6)) //13 7
}
