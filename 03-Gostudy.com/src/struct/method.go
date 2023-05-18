package main

import "fmt"

// 标识符： 变量名函数名 类型名 方法名
// Go语言中如果是标识符号首写字母是大写的就表示外部可见（暴露的，共有的）

// Dog 这是一个注释
type Cat struct {
}

// 方法
type dog struct {
	name string
}

func NewDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是的作用与特定的类型的函数
// 接受者表示的调用该方法的具体的类型变量 多用于类型变量的小写字母表示  有点像类的方法名
// 表示的 struct 的属性方法  类似与类的方法

// 表示的值接收者 值传递
func (d dog) wang() {
	fmt.Printf("%s 旺旺……", d.name)
}

// 表示的指针的接收者 指针传递 也是表示引用传递改变副本也是改变原理的值
func (d *dog) wang2() {
	fmt.Printf("%s hah……", d.name)
}

// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {

	var p1 = Person{
		Name: "李四",
		Age:  20,
		Sex:  "男",
	}
	p2 := p1
	p2.Name = "王五"

	fmt.Printf("%v\n", p2)
	fmt.Printf("%v\n", p1)
	// 结构体是一种值类型的变量
	// 值类型：改变变量副本的时候 是不会概念变量本身的值（数组 基本数据类型 结构体）
	// 引用类型： 改变变量的副本的时候 会改变变量本身的值（切片 map）
	fmt.Printf("-----------------------------------------\n")

	d1 := NewDog("xiaoyan")
	d1.wang()
	d1.wang2()
}
