package main

import "fmt"

// type 有点像的java中的class
// 有点像java的一个的对象的概念
// 如果要访问结构体成员，需要使用点号 . 操作符，格式为： 结构体.成员名

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量： 有点像是的传递java中的传递对象
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

// 小写表示的私有的类型 其他的不能访问和使用 如果是大写的，那就是外面的也是可以访问。
type person struct {
	name   string
	gender string
}

// 构造函数 返回的结构体还是结构体指针 当结构体比较大的时候 最好使用的返回的结构体指针：func newPerson(name, gender string) *person
func newPerson(name, gender string) person {
	return person{
		name:   name,
		gender: gender,
	}
}

// go语言中的函数永远都海拷贝 不会影响原理的值
func f(p person) {
	p.name = "xiaoling"
	fmt.Println(p.name)
}

// 如果是传入的内存地址 那就是的会修改的原来的值
func f2(p *person) {
	p.name = "xiaoling"
	fmt.Println(p.name)
}

type x struct {
	a int8
	b int8
	c string
}

func main() {
	fmt.Println("---------------------------结构体的实例化-----------------------------")
	var per person // 有点类似与java 中的Person per=new Person()的意思。
	fmt.Printf("per 的类型是：%T\n", per)

	var pers = new(person) // 实例化类的对象
	pers.name = "炸弹那三"     // (*pers).name="xiaoyan" 等价
	(*pers).name = "xiaoyan"
	fmt.Printf("%v", pers)

	// 实例化类的对象一种方法
	var p3 = &person{}
	p3.name = "zhangsan"
	p3.gender = "hahha"
	fmt.Printf((*p3).name) // p3.name

	p4 := person{
		name:   "hah",
		gender: "女",
	}
	fmt.Printf("%v", p4)

	p5 := &person{
		name:   "hah2",
		gender: "女",
	}
	fmt.Printf("%v", p5)

	p6 := &person{
		name: "hah3",
	}
	fmt.Printf("%v", p6)

	fmt.Println("---------------------------结构体的初始化-----------------------------")
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	fmt.Println("---------------------------访问结构体成员-----------------------------")
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
	fmt.Println("-------------------------结构体作为函数参数--------------------------")
	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
	fmt.Println("-------------------------匿名结构体-------------------------------")
	// 匿名结构体多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "xjl"
	s.age = 10
	fmt.Println(s.name, " ", s.age)
	fmt.Printf("%T\n", s)
	fmt.Println("-------------------------指针类型的结构体---------------------------")

	var p person
	p.name = "zhuangxiaoyan"
	f(p)
	fmt.Println(p.name)
	// 根据内存地址来实现对原理的进行拷贝 使用的指针来实现
	f2(&p)
	fmt.Println(p.name)
	fmt.Println("-----------------------------------------------------------------")
	var p2 = new(person)
	p2.name = "xiaoyan"       // 相当于是(*p2).name="xiaoyan" 不过是一种语法糖
	fmt.Printf("%T\n", p2)    //*main.person
	fmt.Printf("p2=%p\n", p2) //p2=&main.person{name:"", city:"", age:0}

	fmt.Println("----------------------内存空间-------------------------------------")
	m := x{
		a: int8(10),
		b: int8(20),
		c: "庄小焱",
	}

	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	// 内存对齐的原理

}
