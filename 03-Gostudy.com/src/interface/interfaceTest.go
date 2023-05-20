package main

import "fmt"

// 定义一个Usber的接口
type Usber interface {
	Start(string2, string3 string)
	Stop()
}

// 如果接口里面有方法的话，必须通过结构体或者是自定义类型的实现这个接口

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

// 值接收者类型的实现
// 结构体值接收者后的结构体和指针类型的都是可以赋值给接口变量
func (p Phone) Start2(str1, str2 string) {
	fmt.Printf("Phone Start function start………… %v--%v \n", str1, str2)
}

// 指针接收者类型的实现
func (p *Phone) Start(str1, str2 string) {
	fmt.Printf("Phone Start function start………… %v--%v \n", str1, str2)
}

func (p *Phone) Stop() {
	fmt.Println("Phone Stop function start…………")
}

func (p *Camera) Start(str1, str2 string) {
	fmt.Printf("Camera Start function start………… %v--%v \n", str1, str2)
}

func (p *Camera) Stop() {
	fmt.Println("Camera Stop function start…………")
}

func (p *Camera) Run() {
	fmt.Println("Camera Run function start…………")
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.Start("Computer", "work")
	usb.Stop()
}

func main() {
	p := Phone{
		Name: "苹果",
	}
	// p.Start("hhaa", "heheh")
	// p.Stop()

	var p1 Usber // 接口就是一种数据类型
	p1 = &p      // 表示手机实现Usber接口
	p1.Start("lllll", "xxxxxxxx")
	p1.Stop()

	c := Camera{Name: "佳能"}
	var c1 Usber = &c // 表示实现了Usber 接口 就是等于implement 关键字  接口中所有的方法必须实现 否者就会报错
	c1.Stop()

	// c1.Run() c1 不能调用Camera本身调用方法 只能调用的接口本身的方法
	c.Run()
	c.Stop()
	c.Start("test1", "test2")
	fmt.Println("----------------------------------------------------------")
	var computer = Computer{}
	//  computer.work(p) 这个work 需要提供的是Usber 这个的一个接口类型。
	computer.work(p1)
	computer.work(c1)
}
