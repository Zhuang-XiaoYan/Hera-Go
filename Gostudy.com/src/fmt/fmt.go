package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// %b	表示为二进制
// %c	该值对应的unicode码值
// %d	表示为十进制
// %o	表示为八进制
// %x	表示为十六进制，使用a-f
// %X	表示为十六进制，使用A-F
// %U	表示为Unicode格式：U+1234，等价于”U+%04X”
// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
// %b	无小数部分、二进制指数的科学计数法，如-123456p-78
// %e	科学计数法，如-1234.456e+78
// %E	科学计数法，如-1234.456E+78
// %f	有小数部分但无指数部分，如123.456
// %F	等价于%f
// %g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
// %G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
// %s	直接输出字符串或者[]byte
// %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
// %x	每个字节用两字符十六进制数表示（使用a-f
// %X	每个字节用两字符十六进制数表示（使用A-F）
// %p	表示为十六进制，并加上前导的0x
// %f	默认宽度，默认精度
// %9f	宽度9，默认精度
// %.2f	默认宽度，精度2
// %9.2f	宽度9，精度2
// %9.f	宽度9，精度0

// 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

// 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
// func Fscan(r io.Reader, a ...interface{}) (n int, err error)
// func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
// func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

// 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
// func Sscan(str string, a ...interface{}) (n int, err error)
// func Sscanln(str string, a ...interface{}) (n int, err error)
// func Sscanf(str string, format string, a ...interface{}) (n int, err error)

func main() {

	var (
		name    string
		age     int
		married bool
	)
	// fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
	// fmt.Scan(&name, &age, &married)
	// Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
	// fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
