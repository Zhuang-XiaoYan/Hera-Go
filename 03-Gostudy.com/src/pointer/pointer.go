package main

import "fmt"

// 大写表示public 方法 小写是private 方法
func PointDemo() {
	// &取地址
	// *：表示的是根据地址取值
	n := 18
	m := &n
	fmt.Println(&n)
	fmt.Println(*(&n))
	fmt.Println(m)

	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
func main() {

	// 执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，
	// 还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，
	// 是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
	//var a *int // 表示空指针
	//*a = 100
	//fmt.Println(*a)
	//
	//var b map[string]int
	//b["沙河娜扎"] = 100
	//fmt.Println(b)

	a1 := new(int)
	b1 := new(bool)
	fmt.Printf("%T\n", a1) // *int
	fmt.Printf("%T\n", b1) // *bool
	fmt.Println(*a1)       // 0
	fmt.Println(*b1)       // false
	fmt.Println("--------------------------------------------")
	var s2 *int // 表示是的一空指针  nil
	fmt.Println(s2)
	var s3 = new(int)
	fmt.Println(s3)
	fmt.Println("--------------------------------------------")
	// make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	// new和make都是用于申请内存的
	// new 很少用 一般是用于基本类型数据的申请内存 string int 返回值都是对应的类型的指针(*string *int)
	// make 是用于slice map chan申请内存的 make函数返回的都是对应的这三个类型的本身
	var b2 map[string]int
	b2 = make(map[string]int, 10)
	b2["沙河娜扎"] = 100
	fmt.Println(b2)
}
