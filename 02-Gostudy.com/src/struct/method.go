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

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是的作用与特定的类型的函数
// 接受者表示的调用该方法的具体的类型变量 多用于类型变量的小写字母表示  有点像类的方法名

// 表示的值接收者 值传递
func (d dog) wang() {
	fmt.Printf("%s 旺旺……", d.name)
}

// 表示的指针的接收者 指针传递
func (d *dog) wang2() {
	fmt.Printf("%s hah……", d.name)
}

// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

func main() {
	d1 := newDog("xiaoyan")
	d1.wang()
}
