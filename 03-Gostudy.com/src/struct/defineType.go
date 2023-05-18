package main

import "fmt"

// 结构体 自定义类型个类型别名  注意非本地类型是不能定义方法的 也就是说我们不能给其他包定义方法

// type后面是的类型
type myInt int      // 自定义类型
type youInt = myInt // 类型别名 就是仍外去了一个其他的名字

// 给自定义类型定义方法
func (m myInt) Hell() {
	fmt.Println("hahahhhhhhhhhhhhhhhhhhhhhhhh")
}

type myIntFunction int

// 给自定的类型添加方法 不能给其他的包中添加方法 只能在自己包中的添加自定义的方法
func (my myIntFunction) hello() {
	fmt.Println("我是一个myIntdefin")
}

type Person struct {
	Name  string
	Age   int
	Hobby []string
	Map1  map[string]string
}

func main() {

	var p Person
	p.Name = "xiaoyan"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "java"
	p.Hobby[1] = "goland"
	p.Hobby[2] = "python"

	p.Map1 = make(map[string]string)
	p.Map1["test"] = "背景"
	p.Map1["address"] = "上海"
	p.Map1["phone"] = "13878786390"

	fmt.Printf("%#v", p)

	var n myInt
	n = 100
	var m youInt
	m.Hell()
	m = 100
	fmt.Println(n)
	fmt.Println(m)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)
}
