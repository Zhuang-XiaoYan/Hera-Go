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

type person struct {
	name   string
	gender string
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

func main() {
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
	fmt.Println("-------------------------结构体作为函数参数-------------------------------")
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
	fmt.Println("-------------------------指针类型的结构体-------------------------------")

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

}
