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

func calc(base int) (func(int) int, func(int) int) {
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
	f16, f17 := calc(10)
	fmt.Println(f16(1), f17(2)) //11 9
	fmt.Println(f16(3), f17(4)) //12 8
	fmt.Println(f16(5), f17(6)) //13 7
}
