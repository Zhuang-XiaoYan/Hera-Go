package main

import "fmt"

type cat struct {
}

type dog struct {
}

type person struct {
}

// 定义一个speaker的类型
type speaker interface {
	// 抽象的方法 只要是实现了speak()方法的变量都是speaker类型
	speak()
}

func (c cat) speak() {
	fmt.Println("喵喵……")
}

func (d dog) speak() {
	fmt.Println("旺旺……")
}
func (p person) speak() {
	fmt.Println("啊啊啊……")
}

// 传入的参数的是一个interface类型
func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

}
