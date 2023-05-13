package main

import "fmt"

type myIntFunction int

// 给自定的类型添加方法 不能给其他的包中添加方法 只能在自己包中的添加自定义的方法
func (my myIntFunction) hello() {
	fmt.Println("我是一个myIntdefin")
}

func main() {

}
