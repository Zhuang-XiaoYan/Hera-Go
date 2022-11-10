package main

import "fmt"

// 定义一个car的接口类型
type car interface {
	run()
}

type baoma struct {
	brand string
}

func (b baoma) run() {
	fmt.Println("我是宝马 run……")
}

type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Println("我是法拉利 run……")
}

type biyadi struct {
	brand string
}

func (bi biyadi) run() {
	fmt.Println("我是比亚迪 run……")
}

func driver(c car) {
	c.run()
}

func main() {
	var b1 = baoma{
		brand: "宝马",
	}

	var f1 = falali{
		brand: "法拉利",
	}

	var bi = biyadi{
		brand: "比亚迪",
	}
	driver(b1)
	driver(f1)
	driver(bi)
	// 如果一个变量实现了解耦中的规定的所有的方法 那么这个变量就实现了这个接口 可以理解为这个接口类型的变量
}
