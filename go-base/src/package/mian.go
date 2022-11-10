package main

import "fmt"

const pi = 3.14

var x = 100

// 全局变量 -- init --main 函数
func init() {
	fmt.Println(pi, x)
}

func main() {

}
