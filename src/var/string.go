package main

import (
	"fmt"
	"strings"
)

func main() {
	// go 语言中字符串必须是双引号“”   如果是单引号‘’ 表示的字符
	// 不是字符串这个与java中的定义是相同 但是与python不同 python的字符是不区分单引号和双引号了
	// 涉及到的转义符号的使用
	path := "D:\\gitee\\Hera-Go\\src\\var"
	path1 := "\"D:\\gitee\\Hera-Go\\src\\var\""
	path2 := "'D:\\gitee\\Hera-Go\\src\\var'"
	fmt.Println(path)
	fmt.Println(path1)
	fmt.Println(path2)

	// 定义的多行的字符串
	s1 := `第一行
第二行
			第三行
		`
	fmt.Println(s1)
	// 打印字符串的长度
	fmt.Println(len(s1))
	// 打印字符串的拼接
	s2 := "xjl\\dhakdhka\\jdkajda"
	ss := s2 + s1
	fmt.Println(ss)
	fmt.Println(s1 + s2)
	fmt.Sprintf("%s%s", s1, s2)
	fmt.Println("----------------------------------------------------")
	// 字符串的分割
	arry := strings.Split(s2, "\\") // 分割成为的字符串的数组的类型的与java中的是一样的类型
	fmt.Printf("%T\n%v", arry, arry)

	// 字符串是否包含 利用的strings的函数来实现
	fmt.Println(strings.Contains(s2, "xjl"))
	// 后缀
	fmt.Println(strings.HasSuffix(s2, "xjl"))
	// 前缀
	fmt.Println(strings.HasPrefix(s2, "xjl"))

	// 子串出现的位置
	fmt.Println(strings.LastIndex(s2, "x"))
	// 拼接
	fmt.Println(strings.Join(arry, "+"))

}
