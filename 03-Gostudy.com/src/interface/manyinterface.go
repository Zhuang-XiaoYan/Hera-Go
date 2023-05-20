package main

import "fmt"

type Animaler1 interface {
	SetName(name string)
}

type Animaler2 interface {
	GetName() string
}

type Dog1 struct {
	Name string
}

func (d *Dog1) SetName(name string) {
	d.Name = name
}

func (d Dog1) GetName() string {
	return d.Name
}

func main() {
	d := &Dog1{
		Name: "xiaohei",
	}
	var d1 Animaler1 = d // 表示让d实现Animaler1 这个接口  相当于是implement 这个关键字等价形式
	var d2 Animaler2 = d // 表示让d实现Animaler2 这个接口
	d1.SetName("大傻逼")
	fmt.Println(d2.GetName())
}
