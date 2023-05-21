package tools

import "fmt"

const pi = 3.14

var x = 100

// 全局变量 -- init --main 函数
// init 函数是优先于的main 执行的 运行时被最后导入的包会最先初始化并调用init的函数
func init() {
	fmt.Println(pi, x)
}

// 表示公有方法
func Add(x, y int) int {
	return x + y
}

// 表示私有方法
func add(x, y int) int {
	return x + y
}
