package main

import "fmt"

// if 条件的判断与的java是否相同
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
func ifDemo2() {
	// score的作用域只在if 中 而是在整个函数中。
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// 与java是相同的
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

// 相等于的java中for(;i<10;)省略了
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
func forDEmo4() {
	// 这是一个死循环的结构
	for {
		fmt.Println(2)
	}
}
func main() {
	// if 条件的判断
	age := 19
	if age < 18 {
		fmt.Println("hahahh")
	} else {
		fmt.Println("51415645")
	}
	fmt.Println("-------------------------------------------")
	s := "hello xjl"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
