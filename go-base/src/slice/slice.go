package main

import "fmt"

func main() {
	// 切片的定义  与的java的Arraylist的一样的作用  在声明的时候不需要指定数组的长度 同时slice具有自动扩容的功能能。
	var s1 []int    //定义一个存放int类型的元素的切片
	var s2 []string // 定义存放一个string的元素的切片

	fmt.Println(s1, s2)

	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(c)              //[false true]
	fmt.Println(c == nil)       //false
	fmt.Println(d)
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	// 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
	// 切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]  与python中的元组和数组的概念类似 具有相同的方法
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// 切片的只想底层数组
	// 切片的长度就是元素的个数
	// 切片的容量就是底层数组从切片的第一个元素到最后一个元素的数量
	// 切片是支持扩容的
	// 切片是引用类型 只想一个底层的数组

	// make使用make()函数构造切片 :([]T, size, cap)
	// T:切片的元素类型
	// size:切片中元素的数量
	// cap:切片的容量
	a1 := make([]int, 2, 10)
	fmt.Println(a)       //[0 0]
	fmt.Println(len(a1)) //2
	fmt.Println(cap(a1)) //10
	// 切片的本质就是对底层数组的封装， 只想了一段连续的内存，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。 真正的数据都是保存在内存中。
	fmt.Println("-------------------------------------------------------------------")
	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 \切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：
	var s3 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s4 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s5 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(s3, s4, s5)

	// 切片的赋值
	s6 := []int{1, 3, 5}
	s7 := s6 // 两个都是的指向同样的内存空间 切片是一种引用
	fmt.Println(s6, s7)
	s6[0] = 100
	fmt.Println(s6, s7)
	fmt.Println("-------------------------------------------------------------------")
	// 切片的遍历
	s8 := []int{1, 3, 5}
	for i := 0; i < len(s8); i++ {
		fmt.Println(i, s8[i])
	}
	for index, value := range s8 {
		fmt.Println(index, value)
	}
}
