package main

import "fmt"

type Animaler interface {
	SetName(name string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c Cat) GetName() string {
	return c.Name
}

func main() {
	d := &Dog{
		Name: "xiaoyan",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("zhaungxiaoyan")
	fmt.Println(d1.GetName())

	c := &Cat{
		Name: "yatou",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())
	c1.SetName("shabi")
	fmt.Println(c1.GetName())

}
