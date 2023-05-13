package main

import "fmt"

// go-base 语言的变量必须先声明才能使用

// 声明为全局变量 可以不引用不会报错
var name string
var age int
var isOk bool

// 批量声明变量
var (
	name2 string
	ages2 int
	isOk2 bool
)

func main() {
	// 函数内部的变量必须使用 否则编译不通过
	var name1 string
	var age1 int
	var isOk1 bool
	// 打印变量
	fmt.Print(name1)
	// 格式化打变量的类型
	fmt.Printf("%T \n", name)
	// 打印变量的值
	fmt.Printf("%v \n", age1)
	// 打印变量的十进制类型
	fmt.Printf("%d \n", isOk1)
	// 打印变量并空行 等于 fmt.Printf("%v \n",name)
	fmt.Println(name1, age1, isOk1)

	//变量的声明与赋值  不推荐这样写
	var s1 = "xjl"
	fmt.Println(s1)

	// 变量是有类型的推道的功能
	var s2 = 20
	fmt.Println(s2)

	// 简单变量声明 只能在函数中使用 也是成为局部变量  :=不能使用在函数外
	s3 := "zhuangxiaoyan"
	fmt.Println(s3)

	// 匿名变量 _
	// 函数外的每个语句都必须以关键字开始（var、const、func等）,_多用于占位，表示忽略值。
	x, _ := 22, 10
	fmt.Println(x)
	// 同一个作用域中不能重复声明同名变量
}
