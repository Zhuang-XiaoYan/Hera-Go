package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat1 struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat1) move() {
	fmt.Println("cat1 is moving……")
}

func (c cat1) eat(food string) {
	fmt.Println("cat is eat ", food)
}

func (c chicken) move() {
	fmt.Println("chickne is moving……")
}

func (c chicken) eat(food string) {
	fmt.Println("chickne is eat ", food)
}

func main() {

	var a1 animal // 定义一个接口类型
	bc := cat1{   // 定一个cat类型的变量bc
		name: "tom",
		feet: 4,
	}
	a1 = bc
	a1.eat("大鱼")
	fmt.Println(a1)

	ch := chicken{ // 定一个cat类型的变量bc
		feet: 4,
	}
	ch.eat("米…………")
	fmt.Println(ch)
	a1 = ch
	fmt.Printf("%T", a1)
	fmt.Println("------------------------------------")

	c1 := cat1{"jerry", 5} // cat
	c2 := &cat1{"tom", 10} // *cat1

	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

	// 使用值接收者和使用指针接收者实现接口的区别？
	// 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能够存储
	// 指针接收者只能实现接口只能存结构体指针
}
