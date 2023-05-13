package main

import "fmt"

// 接口还是可以嵌套的
type animal1 interface {
	mover
	eater
}

// 同一个结构体可以实现多个接口
type mover interface {
	move()
}

type eater interface {
	eat(food string)
}

// 该结构体实现了两个接口
type cat2 struct {
	name string
	feet int8
}

func (c *cat2) move() {
	fmt.Println("is moving …………")
}

func (c *cat2) eat(food string) {
	fmt.Println("cat eat", food)
}

func main() {
	// 空接口没有必要起名字
	// interface{} 空接口
	// 所有的类型都实现了空接口，也就是任意类型的都能够保存在空接口中

}
