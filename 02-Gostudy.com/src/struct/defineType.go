package main

import "fmt"

// 结构体 自定义类型个类型别名

// type后面是的类型
type myInt int      // 自定义类型
type youInt = myInt // 类型别名 就是仍外去了一个其他的名字

func main() {
	var n myInt
	n = 100
	var m youInt
	m = 100
	fmt.Println(n)
	fmt.Println(m)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)
}
