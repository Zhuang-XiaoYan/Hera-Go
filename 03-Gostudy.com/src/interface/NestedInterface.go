package main

import "fmt"

type Ainterface interface {
	SetName(name string)
}

type Binterface interface {
	GetName() string
}

// 表示接口的嵌套
type Animaler3 interface {
	Ainterface
	Binterface
}

type Dog3 struct {
	Name string
}

func (d *Dog3) SetName(name string) {
	d.Name = name
}

func (d Dog3) GetName() string {
	return d.Name
}

func main() {
	d := &Dog3{
		Name: "小黑",
	}
	var d3 Animaler3 = d // 表示让d实现Animaler1 这个接口  相当于是implement 这个关键字等价形式
	fmt.Println(d3.GetName())
	d3.SetName("我是大傻逼")
	fmt.Println(d3.GetName())
}
