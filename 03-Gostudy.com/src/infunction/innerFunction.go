package main

import "fmt"

// close	        主要用来关闭channel
// len	            用来求长度，比如string、array、slice、map、channel
// new	            用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make	            用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append	        用来追加元素到数组、slice中
// panic和recover	用来做错误处理

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// 这个是内部匿名函数
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}() // 立即执行的函数
	panic("panic in B") // 程序崩溃了
}

func funcC() {
	fmt.Println("func C")
}

// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。

func main() {
	funcA()
	funcB()
	funcC()
}
