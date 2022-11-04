package main

import "fmt"

func main() {
	// map 相当于java中的hashmap
	var hash map[string]int
	hash = make(map[string]int) // 初始化
	hash["xjl"] = 20
	fmt.Println(hash) // 还么有在在内存开辟空间 导致没有初始化
	// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
}
